package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
	mockdb "udemy-course/db/mock"
	"udemy-course/db/sqlc"
	"udemy-course/utils"
)

type eqCreateUserParamsMatcher struct {
	args     sqlc.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {

	args, ok := x.(sqlc.CreateUserParams)
	if !ok {
		return false
	}

	err := utils.CheckPassword(e.password, args.HashedPassword)
	if err != nil {
		return false
	}

	e.args.HashedPassword = args.HashedPassword

	return reflect.DeepEqual(e.args, args)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches args %v and password %v", e.args, e.password)
}

func EqCreateUserParams(args sqlc.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{args, password}
}

func TestUserApi(t *testing.T) {
	user, password := randomUser()
	testCases := []struct {
		name          string
		body          gin.H
		buildStub     func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "ok",
			body: gin.H{
				"username":  user.Username,
				"full_name": user.FullName,
				"email":     user.Email,
				"password":  password,
			},
			buildStub: func(store *mockdb.MockStore) {
				args := sqlc.CreateUserParams{
					Username: user.Username,
					FullName: user.FullName,
					Email:    user.Email,
				}
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(args, password)).Times(1).Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		//{
		//	name: "internalServerError",
		//	body: gin.H{
		//		"username":  user.Username,
		//		"full_name": user.FullName,
		//		"email":     user.Email,
		//		"password":  password,
		//	},
		//	buildStub: func(store *mockdb.MockStore) {
		//
		//		store.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Times(1).Return(user, nil)
		//	},
		//	checkResponse: func(recorder *httptest.ResponseRecorder) {
		//		require.Equal(t, http.StatusOK, recorder.Code)
		//	},
		//},
	}

	for i, _ := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStub(store)

			recorder := httptest.NewRecorder()
			server := newTestServer(t, store)

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/user"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}

func randomUser() (user sqlc.Users, password string) {
	password = utils.RandomString(9)
	hashedPassword, _ := utils.HashPassword(password)

	user = sqlc.Users{
		Username:       utils.RandOwner(),
		FullName:       utils.RandomString(12),
		HashedPassword: hashedPassword,
		Email:          utils.RandomEmail(),
		CreatedAt:      time.Now(),
	}
	return
}
