package router

//go:generate qtc -dir=../template
import (
	"github.com/kataras/iris"
	"micro/client/router/user"
)

func Router() *iris.Application {

	app := iris.New()

	app.StaticWeb("/iris/static", "../static")

	userRouter := app.Party("/api/user")

	//userRouter.Use(middleware.JWT)

	{
		userRouter.Post("/signin", user.Signup)
		userRouter.Post("/login", user.Login)
		userRouter.Get("/logout", user.Logout)

	}

	return app
}
