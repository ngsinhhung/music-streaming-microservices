package response

import (
	"github.com/gin-gonic/gin"
	r "music-streaming-microservices/common-lib/response"
)

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}

func SuccessResponse(c *gin.Context, key string, data interface{}) {
	code := r.StatusCode[key]
	message := r.ReasonPhrases[key]
	c.JSON(code, Response{
		StatusCode: code,
		Message:    message,
		Data:       data,
	})
}

func ErrorResponse(c *gin.Context, key string, data interface{}) {
	code := r.StatusCode[key]
	message := r.ReasonPhrases[key]
	c.JSON(code, Response{
		StatusCode: code,
		Message:    message,
		Data:       data,
	})
}
