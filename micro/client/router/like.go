package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/middleware"
)

func Like(app *iris.Application) {
	app.Post("/api/like", middleware.JWT, controller.AddLike)
	app.Delete("/api/delete", middleware.JWT, controller.DelLike)
	app.Get("/api/favorites", middleware.JWT, controller.GetFavorite)
	app.Put("/api/favorites", middleware.JWT, controller.AddCollection)
	app.Post("/api/favorites", middleware.JWT, controller.AddFavorite)
	app.Delete("/api/favorites", middleware.JWT, controller.DelCollection)
	//app.Get("/api/collection", middleware.JWT, controller.GetCollection)
	app.Put("/api/collection", middleware.JWT, controller.AddCollection)
	app.Post("/api/collection", middleware.JWT, controller.AddCollection)
	app.Delete("/api/collection", middleware.JWT, controller.DelCollection)
}
