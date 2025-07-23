package global

import (
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"music-streaming-microservices/email-service/pkg/types"
)

var (
	Configs       types.Configs
	Redisdb       *redis.Client
	NatsConn      *nats.Conn
	NatsJetStream nats.JetStreamContext
)
