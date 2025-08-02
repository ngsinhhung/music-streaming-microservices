package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

func GenerateKey() *rsa.PrivateKey {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic("Failed to generate private key: " + err.Error())
	}
	return privateKey
}

func GetPublicKeyString(privateKey *rsa.PrivateKey) (string, error) {
	publicKey := &privateKey.PublicKey
	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)

	return base64.StdEncoding.EncodeToString(publicKeyBytes), nil
}
