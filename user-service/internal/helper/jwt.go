package helper

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"music-streaming-microservices/user-service/internal/database"
	"music-streaming-microservices/user-service/pkg/types"
	"time"
)

type TokenPair struct {
	Token   string `json:"token"`
	RfToken string `json:"rf_token"`
}

func GenerateJTI() uuid.UUID {
	return uuid.New()
}

func CreateTokenPair(sessionId string, payload database.User, privateKey *rsa.PrivateKey) (TokenPair, error) {
	payloadAccessToken := types.PayloadAccessToken{
		UserID: int64(payload.ID),
		Email:  payload.Email,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"payload": payloadAccessToken,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	token, err := t.SignedString(privateKey)
	if err != nil {
		return TokenPair{}, err
	}

	payloadRefreshToken := types.PayloadRefreshToken{
		UserID: int64(payload.ID),
		JTI:    sessionId,
	}

	rfToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"payload": payloadRefreshToken,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	})
	rfTokenString, err := rfToken.SignedString(privateKey)
	if err != nil {
		return TokenPair{}, err
	}

	return TokenPair{
		Token:   token,
		RfToken: rfTokenString,
	}, nil

}

func VerifyToken() {
	panic("implement me")
}
