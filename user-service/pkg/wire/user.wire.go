//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"music-streaming-microservices/user-service/internal/controllers"
	"music-streaming-microservices/user-service/internal/repositories"
	"music-streaming-microservices/user-service/internal/services"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		repositories.NewUserRepository,
		services.NewUserService,
		controllers.NewUserController,
	)
	return new(controllers.UserController), nil
}
