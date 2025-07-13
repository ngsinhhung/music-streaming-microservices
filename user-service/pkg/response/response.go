package response

import (
	"github.com/gin-gonic/gin"
	r "music-streaming-microservices/common-lib/response"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponseData struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Detail     interface{} `json:"detail"`
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	if message == "" {
		message = r.ReasonPhrases[code]
	}

	c.JSON(code, Response{
		StatusCode: code,
		Message:    message,
		Data:       data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, detail interface{}) {
	if message == "" {
		message = r.ReasonPhrases[code]
	}
	c.JSON(code, ErrorResponseData{
		StatusCode: code,
		Message:    message,
		Detail:     detail,
	})
}
