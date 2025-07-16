package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetHashString(key string) string {
	hash := sha256.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
