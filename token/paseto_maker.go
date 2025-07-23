package token

import (
	"aidanwoods.dev/go-paseto"
	"time"
)

type PasetoMaker struct {
	privateKey paseto.V4AsymmetricSecretKey
	publicKey  paseto.V4AsymmetricPublicKey
}

func NewPasetoMaker(symmetricKey string) *PasetoMaker {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	return &PasetoMaker{
		privateKey,
		publicKey,
	}
}

func (maker *PasetoMaker) CreateToken(userID string, duration time.Duration) (string, *PasetoPayload, error) {
	now := time.Now()
	payload := &PasetoPayload{
		UserID:    userID,
		IssuedAt:  now,
		ExpiredAt: now.Add(duration),
	}

	token := paseto.NewToken()
	token.SetIssuedAt(payload.IssuedAt)
	token.SetExpiration(payload.ExpiredAt)
	token.SetString("user_id", userID)

	signed := token.V4Sign(maker.privateKey, nil)

	return signed, payload, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*PasetoPayload, error) {
	//token, err := v4

	return nil, nil
}
