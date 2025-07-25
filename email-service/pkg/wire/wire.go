//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"music-streaming-microservices/email-service/internal/handler"
	"music-streaming-microservices/email-service/internal/repository"
)

func InitEmailHandler() handler.IHandler {
	wire.Build(repository.NewRepository, handler.NewHandler)
	return nil
}
