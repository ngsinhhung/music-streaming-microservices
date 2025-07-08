package controllers

import (
	"github.com/gin-gonic/gin"
	"music-streaming-microservices/user-service/internal/services"
)

type UserController struct {
	userServices services.IUserService
}

func NewUserController(services services.IUserService) *UserController {
	return &UserController{
		userServices: services,
	}
}

func (uc *UserController) Register(c *gin.Context) {
}
