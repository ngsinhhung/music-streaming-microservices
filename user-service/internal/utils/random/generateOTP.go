package random

import (
	"math/rand"
	"time"
)

func GenerateOTP() (otp int) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp = rng.Intn(999999-100000) + 100000 // Generates a random number between 100000 and 999999
	return
}
