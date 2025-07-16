package utils

import "fmt"

func GetKeyOTP(hashedKey string) string {
	return fmt.Sprintf("usr%s:otp", hashedKey)
}
