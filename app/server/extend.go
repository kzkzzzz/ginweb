package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Response struct {
		Code    int
		Message string
		Data    any
	}

	MyContext struct {
		*gin.Context
	}

	handler func(ctx *MyContext)
)

func WrapHandler(h handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h(&MyContext{Context: ctx})
	}
}

func (c *MyContext) Success(data any) {
	c.Context.JSON(http.StatusOK, &Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func (c *MyContext) Fail(message string, code ...int) {
	r := &Response{
		Message: message,
	}
	if len(code) > 0 {
		r.Code = code[0]
	}
	c.Context.JSON(http.StatusOK, r)
}
