package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
)

func Category(app *iris.Application) {
	categoryRouter := app.Party("/api/category")
	{
		categoryRouter.Get("", controller.GetCategory)
	}
}
