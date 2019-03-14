package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/middleware"
)

func Like(app *iris.Application) {
	app.Post("/api/like", controller.AddLike)
	app.Delete("/api/delete", controller.DelLike)
	app.Get("/api/favorite", middleware.JWT, controller.GetCollection)
	app.Put("/api/favorite", middleware.JWT, controller.AddCollection)
	app.Post("/api/favorite", controller.AddCollection)
	app.Delete("/api/favorite", controller.DelCollection)
}
