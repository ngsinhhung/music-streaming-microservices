package repository

import (
	"fmt"
	"music-streaming-microservices/email-service/global"
)

type IRepository interface {
	GetOTP(key string) string
}

type repository struct{}

func (r *repository) GetOTP(key string) string {
	otp, err := global.Redisdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Error getting OTP from Redis: %v\n", err)
	}
	return otp
}

func NewRepository() IRepository {
	return &repository{}
}
