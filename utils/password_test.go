package utils

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	password := "abcdeddd"

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(hashedPassword, "abcdeddd")
	require.NoError(t, err)

	wrongPass := RandString(8)
	err = CheckPassword(hashedPassword, wrongPass)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
