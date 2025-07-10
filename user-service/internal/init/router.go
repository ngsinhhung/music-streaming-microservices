package init

import (
	"github.com/gin-gonic/gin"
	"music-streaming-microservices/user-service/global"
	"music-streaming-microservices/user-service/internal/routers"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine

	mode := global.Configs.Server.Mode
	switch mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	case "release":
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	userRouter := routers.UserRouterInstance

	v1 := r.Group("/v1/api")

	{
		userRouter.InitUserRouter(v1)
	}

	return r
}
