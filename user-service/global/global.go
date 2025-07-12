package global

import (
	"github.com/jackc/pgx/v5"
	"music-streaming-microservices/user-service/pkg/types"
)

var (
	Configs      types.Configs
	PostgresConn *pgx.Conn
)
