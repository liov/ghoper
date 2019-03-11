package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/middleware"
	"hoper/client/router/user"
)

func UserRouter(app *iris.Application) {

	userRouter := app.Party("/api/user")
	{
		userRouter.Get("/active/{id:int}/{secret:string}", controller.ActiveAccount)
		userRouter.Post("/signup", controller.Signup)
		userRouter.Post("/login", controller.Login)
		userRouter.Get("/logout", middleware.JWT, controller.Logout)
		userRouter.Post("/active", controller.ActiveSendMail)
		userRouter.Get("/get", middleware.Login, controller.LoginFlag)
		userRouter.Post("/rpc/login", user.Login)
		userRouter.Get("/rpc/logout", user.Logout)
		userRouter.Post("/rpc/signup", user.Signup)
		userRouter.Get("/edit", middleware.JWT, controller.GetUserSelf)
		userRouter.Get("/{id:uint64}", user.GetUser)
		userRouter.Put("/", middleware.JWT, controller.UpdateUser)
	}
}
