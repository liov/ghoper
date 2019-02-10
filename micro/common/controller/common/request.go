package common

import (
	"github.com/kataras/iris"
	"micro/common/controller/binding"
)

type HopeCtx iris.Context

func BindWithJson(ctx iris.Context, obj interface{}) error {
	return binding.JSON.Bind(ctx, obj)
}
