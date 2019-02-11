package router

//go:generate qtc -dir=../template
import (
	"github.com/kataras/iris"
	"service/controller"
	"service/controller/hnsq"
	"service/controller/hwebsocket"
	"service/middleware"
)

func IrisRouter() *iris.Application {
	app := iris.New()

	app.StaticWeb("/iris/static", "../static")

	articleRouter := app.Party("/api/article")
	{
		//获取文章列表
		articleRouter.Get("", controller.GetArticle)
		//获取指定文章
		articleRouter.Get("/:id", controller.GetArticles)
		//新建文章
		articleRouter.Post("", controller.AddArticle)
		//更新指定文章
		articleRouter.Put("/:id", controller.EditArticle)
		//删除指定文章
		articleRouter.Delete("/:id", controller.DeleteArticle)
	}

	momentRouter := app.Party("/api/moment")
	{
		//获取文章列表
		momentRouter.Get("", controller.GetMoments)
		//获取文章列表
		momentRouter.Get("/:id", middleware.GetUser(), controller.GetMoment)
		//新建文章
		momentRouter.Post("", middleware.JWT, controller.AddMoment)
		//更新指定文章
		momentRouter.Put("/:id", middleware.JWT, controller.EditMoment)
		//删除指定文章
		momentRouter.Delete("/:id", middleware.JWT, controller.DeleteMoment)
	}

	//获取标签
	app.Get("/api/tag", controller.GetTags)

	userRouter := app.Party("/api/user")
	{
		userRouter.Get("/active/:id/:secret", controller.ActiveAccount)
		userRouter.Post("/signup", controller.Signup)
		userRouter.Post("/login", controller.Login)
		userRouter.Get("/logout", middleware.JWT, controller.Logout)
		userRouter.Post("/active", controller.ActiveSendMail)
		userRouter.Get("/Get", middleware.JWT, controller.SignInFlag)
	}
	app.Post("/api/comment/:classify", middleware.JWT, controller.AddComment)

	//app.Get("/api/push",controller.Push)

	app.Get("/api/chat/getChat", hwebsocket.GetChat)

	app.Post("/api/nsq", hnsq.Start)

	app.Get("/tpl/index", func(ctx iris.Context) {
		ctx.ContentType("text/html; charset=utf-8")

	})
	app.Get("/api/chat/ws", hwebsocket.Chat)
	return app
}
