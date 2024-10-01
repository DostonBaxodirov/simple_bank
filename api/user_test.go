package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	mockdb "simpleBank/db/mock"
	"simpleBank/tutorial"
	"simpleBank/utils"
	"testing"
)

type eqCreateUserMatcher struct {
	arg      tutorial.CreateUserParams
	password string
}

func (e eqCreateUserMatcher) Matches(x interface{}) bool {
	arg, ok := x.(tutorial.CreateUserParams)
	if !ok {
		return false
	}

	log.Println("arg", arg)

	err := utils.CheckPassword(arg.HashedPassword, e.password)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword

	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg tutorial.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserMatcher{arg, password}
}

func TestCreateUserApi(t *testing.T) {
	user, password := CreateRandomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := tutorial.CreateUserParams{
					Username: user.Username,
					FullName: user.FullName,
					Email:    user.Email,
				}
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).Times(1).Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireMatchBodyUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalServerError",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(tutorial.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		//{
		//	name: "DuplicateUsername",
		//	body: gin.H{
		//		"username":  user.Username,
		//		"password":  password,
		//		"full_name": user.FullName,
		//		"email":     user.Email,
		//	},
		//	buildStubs: func(store *mockdb.MockStore) {
		//		store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(tutorial.User{}, tutorial.ErrUniqueViolation)
		//	},
		//	checkResponse: func(recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusForbidden, recorder.Code)
		//	},
		//},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			// Marshal body data to json
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func CreateRandomUser(t *testing.T) (user tutorial.User, password string) {
	password = utils.RandString(9)
	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	user = tutorial.User{
		Username:       utils.RandOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandString(12),
		Email:          utils.RandomEmail(),
	}

	return
}

func requireMatchBodyUser(t *testing.T, body *bytes.Buffer, user tutorial.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser tutorial.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
