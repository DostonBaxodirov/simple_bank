package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

var ErrExpiredToken = errors.New("Token has expired")
var ErrInvalidToken = errors.New("Token is invalid")

type Payload struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	IssuedDate  time.Time `json:"issued_date"`
	ExpiredDate time.Time `json:"expired_date"`
	// pastdagi methodlar yozmaslik uchun buni ishlatsa bo'ladi va faqat valid() method qo'shilsa yetarli bo'ladi
	//jwt.RegisteredClaims
}

func (payload *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.ExpiredDate), nil
}

func (payload *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedDate), nil
}

func (payload *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return jwt.NewNumericDate(payload.IssuedDate), nil
}

func (payload *Payload) GetIssuer() (string, error) {
	return "", nil
}

func (payload *Payload) GetSubject() (string, error) {
	return payload.Username, nil
}

func (payload *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return nil, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredDate) {
		return ErrExpiredToken
	}
	return nil
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:          tokenID,
		Username:    username,
		IssuedDate:  time.Now(),
		ExpiredDate: time.Now().Add(duration),
	}

	return payload, nil
}

type PasetoPayload struct {
	UserID    string    `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
