package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/middleware"
)

func Like(app *iris.Application) {
	app.Post("/api/like", middleware.GetUser(false), controller.AddLike)
	app.Post("/api/approve", middleware.GetUser(false), controller.Approve)

	app.Get("/api/favorites", middleware.GetUser(false), controller.GetFavorite)
	app.Put("/api/favorites", middleware.GetUser(false), controller.AddCollection)
	app.Post("/api/favorites", middleware.GetUser(false), controller.AddFavorite)
	app.Delete("/api/favorites", middleware.GetUser(false), controller.DelCollection)
	//app.Get("/api/collection", middleware.GetUser, controller.GetCollection)
	app.Put("/api/collection", middleware.GetUser(false), controller.AddCollection)
	app.Post("/api/collection", middleware.GetUser(false), controller.AddCollection)
	app.Delete("/api/collection", middleware.GetUser(false), controller.DelCollection)

	app.Post("/api/comment/{kind}/{refId}", middleware.GetUser(false), controller.AddComment)
	app.Get("/api/comment/{kind}", middleware.GetUser(false), controller.GetComment)
	app.Get("/api/comments/{kind}/{refId}", controller.GetComments)
}
