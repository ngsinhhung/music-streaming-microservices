package init

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"music-streaming-microservices/email-service/global"
)

var ctx = context.Background()

func InitRedis() {
	redisConfig := global.Configs.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Database, // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Unable to connect to Redis: %v\n", err)
		return
	}

	global.Redisdb = rdb
	log.Println("Connected to Redis successfully")
}
