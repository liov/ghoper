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

	app.StaticWeb("/api/static", "../static")

	tmpl := iris.HTML("./template", ".html").Reload(true)

	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)

	tplRouter := app.Party("/tpl")

	tplRouter.Get("/hi", func(ctx iris.Context) {
		ctx.ViewData("Title", "Hi Page")
		ctx.ViewData("Name", "iris")
		ctx.View("hi.html")
	})

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

	tplRouter.Get("/", func(ctx iris.Context) {
		var markdownContents = []byte(`## Hello Markdown
This is a sample of Markdown contents
 
Features
--------
All features of Sundown are supported, including:
*   **Compatibility**. The Markdown v1.0.3 test suite passes with
    the --tidy option.  Without --tidy, the differences are
    mostly in whitespace and entity escaping, where blackfriday is
    more consistent and cleaner.
*   **Common extensions**, including table support, fenced code
    blocks, autolinks, strikethroughs, non-strict emphasis, etc.
*   **Safety**. Blackfriday is paranoid when parsing, making it safe
    to feed untrusted user input without fear of bad things
    happening. The test suite stress tests this and there are no
    known inputs that make it crash.  If you find one, please let me
    know and send me the input that does it.
    NOTE: "safety" in this context means *runtime safety only*. In order to
    protect yourself against JavaScript injection in untrusted content, see
    [this example](https://github.com/russross/blackfriday#sanitize-untrusted-content).
*   **Fast processing**. It is fast enough to render on-demand in
    most web applications without having to cache the output.
*   **Routine safety**. You can run multiple parsers in different
    goroutines without ill effect. There is no dependence on global
    shared state.
*   **Minimal dependencies**. Blackfriday only depends on standard
    library packages in Go. The source code is pretty
    self-contained, so it is easy to add to any project, including
    Google App Engine projects.
*   **Standards compliant**. Output successfully validates using the
    W3C validation tool for HTML 4.01 and XHTML 1.0 Transitional.
	[this is a link](https://github.com/kataras/iris) `)

		ctx.Markdown(markdownContents)
	})
	app.Get("/api/chat/ws", hwebsocket.Chat)
	return app
}
