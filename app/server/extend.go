package server

import (
	"fmt"
	"ginweb/common/errm"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Response struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
	}

	handler func(ctx *gin.Context) (result any, err error)
)

func WrapHandler(h handler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := h(ctx)
		if err != nil {
			fail(ctx, err)
			return
		}
		success(ctx, data)
	}
}

func success(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, &Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}

func fail(ctx *gin.Context, err error) {
	r := &Response{}
	switch v := err.(type) {
	case *errm.Error:
		r.Message = v.Error()
		r.Code = int(v.Code)
	default:
		r.Code = int(errm.CommonErrorCode)
		r.Message = fmt.Sprintf("Error: %s", err.Error())
	}
	ctx.JSON(http.StatusOK, r)
}
