package repository

import (
	"music-streaming-microservices/email-service/global"
)

type IRepository interface {
	GetOTP(key string) (string, error)
}

type repository struct{}

func (r *repository) GetOTP(key string) (string, error) {
	return global.Redisdb.Get(ctx, key).Result()
}

func NewRepository() IRepository {
	return &repository{}
}
