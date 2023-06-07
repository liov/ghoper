package router

import (
	"github.com/kataras/iris"
	"hoper/controller"
)

func Category(app *iris.Application) {

	//获取标签
	app.Get("/api/tag", controller.GetTags)

	categoryRouter := app.Party("/api/category")
	{
		categoryRouter.Get("", controller.GetCategory)
	}
}
