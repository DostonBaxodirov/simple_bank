package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
	"simpleBank/utils"
	"testing"
	"time"
)

func TestNewPasetoMaker(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandString(32))
	require.NoError(t, err)

	username := utils.RandOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestPasetoTokenExpired(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(utils.RandOwner(), -time.Minute)
	require.NoError(t, err)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)

}

func TestInvalidPasetoTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(utils.RandOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(utils.RandString(33))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
