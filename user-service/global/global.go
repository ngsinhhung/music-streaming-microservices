package global

import (
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"music-streaming-microservices/user-service/pkg/types"
)

var (
	Configs      types.Configs
	PostgresConn *pgx.Conn
	Redisdb      *redis.Client
)
