package controller

import "github.com/kataras/iris"

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/4
 * @description：
 */

func Jsonp(ctx iris.Context) {
	callback := ctx.URLParam("callback")
	ctx.WriteString(callback + "('jsonp测试')")
}
