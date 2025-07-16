package controllers

import (
	"github.com/gin-gonic/gin"
	r "music-streaming-microservices/common-lib/response"
	"music-streaming-microservices/user-service/internal/services"
	"music-streaming-microservices/user-service/pkg/response"
	"music-streaming-microservices/user-service/validation"
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
	var userRegisterRequest validation.UserRegisterSchema
	if err := c.ShouldBindJSON(&userRegisterRequest); err != nil {
		response.ErrorResponse(c, r.BAD_REQUEST, "", err)
		return
	}
	code, msg, data := uc.userServices.Register(userRegisterRequest)
	if code == r.CREATED {
		response.SuccessResponse(c, code, msg, data)

	} else {
		response.ErrorResponse(c, code, msg, nil)
	}

}
