package api

import (
	"github.com/kataras/iris"
)

func ApiMiddle(ctx iris.Context) {
	currentRouteName:= ctx.GetCurrentRoute().Name()
	params:=ctx.Params().Store
	for i:= range params{
		key:=params[i].Key
		val:=params[i].ValueRaw
	}
}
