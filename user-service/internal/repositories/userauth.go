package repositories

import (
	"music-streaming-microservices/user-service/global"
	"time"
)

type IUserAuthRepository interface {
	AddOTP(key string, otp int, expirationTime int64) error
}

type userAuthRepository struct {
}

func (u *userAuthRepository) AddOTP(key string, otp int, expirationTime int64) error {
	return global.Redisdb.Set(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
