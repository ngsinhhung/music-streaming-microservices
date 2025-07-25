package repositories

import (
	"music-streaming-microservices/user-service/global"
	_const "music-streaming-microservices/user-service/internal/const"
	"time"
)

type IUserAuthRepository interface {
	AddData(key string, data interface{}) error
	GetData(key string) (string, error)
}

type userAuthRepository struct {
}

func (u *userAuthRepository) AddData(key string, data interface{}) error {
	return global.Redisdb.SetEx(ctx, key, data, time.Duration(_const.TIME_OTP_REGISTER)*time.Minute).Err()
}

func (u *userAuthRepository) GetData(key string) (string, error) {
	return global.Redisdb.Get(ctx, key).Result()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
