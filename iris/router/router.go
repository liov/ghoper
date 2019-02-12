package router

//go:generate qtc -dir=../template
import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/middleware/pprof"
	"github.com/kataras/iris/middleware/recover"
	"service/controller"
	"service/controller/common/logging"
	"service/controller/hnsq"
	"service/controller/hwebsocket"
	"service/controller/tmpl/html"
	"service/controller/tmpl/markdown"
	"service/controller/tmpl/pug"
	"service/middleware"
	"time"
)

func IrisRouter() *iris.Application {
	app := iris.New()

	app.StaticWeb("/api/static", "../static")

	tmpl := iris.HTML("./template", ".html").Reload(true)

	tmplPug := iris.Pug("./template", ".pug").Reload(true)

	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)
	app.RegisterView(tmplPug)

	app.Use(recover.New())

	globalLocale := i18n.New(i18n.Config{
		Default:      "en-US",
		URLParameter: "lang",
		Languages: map[string]string{
			"en-US": "../locales/locale_en-US.ini",
			"zh-CN": "../locales/locale_zh-CN.ini"}})
	app.Use(globalLocale)

	app.Logger().Printer.SetOutput(logging.F)

	//auth
	authConfig := basicauth.Config{
		Users:   map[string]string{"admin": "lby604"},
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)
	app.Any("/tpl/pprof/{action:path}", authentication, pprof.New())

	app.Get("/auth/{provider}/callback", func(ctx iris.Context) {

		user, err := middleware.CompleteUserAuth(ctx)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef("%v", err)
			return
		}
		ctx.ViewData("", user)
		if err := ctx.View("user.html"); err != nil {
			ctx.Writef("%v", err)
		}
	})

	//缓存10s
	tplRouter := app.Party("/tpl")
	{
		tplRouter.Get("/hi", cache.Handler(10*time.Second), html.HtmlTest)
		tplRouter.Get("/pug", pug.PugTest)
		tplRouter.Get("/markdown", markdown.MarkdownTest)
		tplRouter.Get("/time", iris.Cache304(10*time.Second), html.Time)
		tplRouter.Get("/auth", authentication, html.Auth)
		tplRouter.Get("/logout/{provider}", func(ctx iris.Context) {
			middleware.Logout(ctx)
			ctx.Redirect("/", iris.StatusTemporaryRedirect)
		})

		tplRouter.Get("/auth/{provider}", func(ctx iris.Context) {
			// try to get the user without re-authenticating
			if gothUser, err := middleware.CompleteUserAuth(ctx); err == nil {
				ctx.ViewData("", gothUser)
				if err := ctx.View("user.html"); err != nil {
					ctx.Writef("%v", err)
				}
			} else {
				middleware.BeginAuthHandler(ctx)
			}
		})

		tplRouter.Get("/", func(ctx iris.Context) {

			ctx.ViewData("", middleware.NewAuth())

			if err := ctx.View("index.html"); err != nil {
				ctx.Writef("%v", err)
			}
		})
	}

	articleRouter := app.Party("/api/article")
	{
		//获取文章列表
		articleRouter.Get("", controller.GetArticle)
		//获取指定文章
		articleRouter.Get("/{id:uint64}", controller.GetArticles)
		//新建文章
		articleRouter.Post("", controller.AddArticle)
		//更新指定文章
		articleRouter.Put("/{id:uint64}", controller.EditArticle)
		//删除指定文章
		articleRouter.Delete("/{id:uint64}", controller.DeleteArticle)
	}

	momentRouter := app.Party("/api/moment")
	{
		//获取文章列表
		momentRouter.Get("", controller.GetMoments)
		//获取文章列表
		momentRouter.Get("/{id:uint64}", middleware.GetUser(), controller.GetMoment)
		//新建文章
		momentRouter.Post("", middleware.JWT, controller.AddMoment)
		//更新指定文章
		momentRouter.Put("/{id:uint64}", middleware.JWT, controller.EditMoment)
		//删除指定文章
		momentRouter.Delete("/{id:uint64}", middleware.JWT, controller.DeleteMoment)
	}

	//获取标签
	app.Get("/api/tag", controller.GetTags)

	userRouter := app.Party("/api/user")
	{
		userRouter.Get("/active/{id:uint64}/{secret:string}", controller.ActiveAccount)
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

	app.Get("/api/chat/ws", hwebsocket.Chat)
	return app
}
