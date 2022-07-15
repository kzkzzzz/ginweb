package server

import (
	"fmt"
	"ginweb/common/logz"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.JSON(http.StatusOK, &Response{
					Message: fmt.Sprintf("系统异常: %v", err),
					Code:    999,
				})
				logz.Errorf("panic: %v\n%s", err, color.RedString("%s", debug.Stack()))
			}
		}()
		ctx.Next()
	}
}
