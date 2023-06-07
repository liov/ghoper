package router

import (
	"github.com/kataras/iris"
	"hoper/controller"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/4
 * @description：
 */

func Jsonp(app *iris.Application) {
	app.Get("/api/jsonp", controller.Jsonp)

}
