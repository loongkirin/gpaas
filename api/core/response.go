package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	ERROR        = 500
	UNAUTHORIZED = 401
	SUCCESS      = 200
)

func Result(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context, msg string, data interface{}) {
	Result(c, SUCCESS, msg, data)
}

func Fail(c *gin.Context, msg string, data interface{}) {
	Result(c, ERROR, msg, data)
}

func Unauthorized(c *gin.Context, msg string, data interface{}) {
	Result(c, UNAUTHORIZED, msg, data)
}
