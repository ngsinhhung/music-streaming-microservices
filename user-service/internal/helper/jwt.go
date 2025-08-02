package helper

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenPair struct {
	Token   string `json:"token"`
	RfToken string `json:"rf_token"`
}

func CreateTokenPair(payload interface{}, privateKey *rsa.PrivateKey) (TokenPair, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"payload": payload,
		"iat":     time.Now().Unix(),
	})
	token, err := t.SignedString(privateKey)
	if err != nil {
		return TokenPair{}, err
	}

	rfToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"payload": payload,
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
