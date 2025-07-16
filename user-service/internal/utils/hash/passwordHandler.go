package hash

import "golang.org/x/crypto/bcrypt"

const (
	DefaultCode int = 10
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCode)
	return string(hash), err
}

func MatchingWithHashPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
