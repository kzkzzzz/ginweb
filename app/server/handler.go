package server

import (
	"ginweb/app/model"
	"ginweb/common/validate"
)

func index(ctx *MyContext) {
	arg := &model.UserSearch{}
	err := ctx.ShouldBindQuery(&arg)
	if err != nil {
		ctx.Fail(validate.TransGin(err))
		return
	}

	res, err := svc.Index(arg)
	if err != nil {
		ctx.Fail(err.Error())
		return
	}
	ctx.Success(res)
}
