package routers

import (
	"github.com/gin-gonic/gin"
	"music-streaming-microservices/user-service/wire"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {

	userController, _ := wire.InitUserRouterHandler()

	publicUserRouter := router.Group("/user")
	{
		publicUserRouter.GET("/signup", userController.Register)
	}

	privateUserRouter := router.Group("/user")
	//privateUserRouter.Use(Limitter())
	//privateUserRouter.Use(Authentication())
	//privateUserRouter.Use(Permissions())
	{
		privateUserRouter.POST("/login")
		privateUserRouter.GET("/info")
	}
}
