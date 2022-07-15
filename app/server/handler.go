package server

import (
	"ginweb/app/model"
	"ginweb/common/errm"
	"ginweb/common/validate"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) (any, error) {
	arg := &model.UserSearch{}
	err := ctx.ShouldBindQuery(&arg)
	if err != nil {
		return nil, errm.ParamsError(validate.TransGin(err))
	}
	res, err := svc.Index(arg)
	if err != nil {
		return nil, err
	}
	return res, nil
}
