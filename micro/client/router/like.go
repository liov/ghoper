package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
)

func Like(app *iris.Application) {
	app.Post("/api/like", controller.AddLike)
	app.Delete("/api/delete", controller.DelLike)
	app.Post("/api/collection", controller.AddCollection)
	app.Delete("/api/collection", controller.DelCollection)
}
