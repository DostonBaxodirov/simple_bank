package tutorial

import (
	"context"
	"github.com/stretchr/testify/require"
	"simpleBank/utils"
	"testing"
)

func CreateRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandString(9))

	if err != nil {
		panic("Cannot hash password")
	}
	arg := CreateUserParams{
		Username:       utils.RandOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandOwner(),
		Email:          utils.RandOwner(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)

	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)
	require.NotZero(t, user1.CreatedAt)

	//require.WithinDurationf(t, user1.PasswordChangedAt,user2.PasswordChangedAt,time.Second)
}
