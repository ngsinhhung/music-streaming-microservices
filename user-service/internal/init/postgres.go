package init

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"music-streaming-microservices/user-service/global"
	"os"
)

func InitPostgres() {
	pgConfig := global.Configs.Postgres
	stringConnPattern := "postgres://%s:%s@%s:%v/%s"

	var s = fmt.Sprintf(stringConnPattern, pgConfig.User, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.Database)

	conn, err := pgx.Connect(context.Background(), s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Println("Connected to database")

}
