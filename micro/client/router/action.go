package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/middleware"
)

func Like(app *iris.Application) {
	app.Post("/api/like", middleware.Login, controller.AddLike)
	app.Post("/api/approve", middleware.Login, controller.Approve)

	app.Get("/api/favorites", middleware.Login, controller.GetFavorite)
	app.Put("/api/favorites", middleware.Login, controller.AddCollection)
	app.Post("/api/favorites", middleware.Login, controller.AddFavorite)
	app.Delete("/api/favorites", middleware.Login, controller.DelCollection)
	//app.Get("/api/collection", middleware.GetUser, controller.GetCollection)
	app.Put("/api/collection", middleware.Login, controller.AddCollection)
	app.Post("/api/collection", middleware.Login, controller.AddCollection)
	app.Delete("/api/collection", middleware.Login, controller.DelCollection)

	app.Post("/api/comment/{kind}/{id}", middleware.Login, controller.AddComment)
	app.Get("/api/comment/{kind}/{id}", middleware.Login, controller.GetComment)
}
