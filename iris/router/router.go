package router

//go:generate qtc -dir=../template
import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/i18n"
	"github.com/kataras/iris/middleware/recover"
	"service/controller"
	"service/controller/common/logging"
	"service/controller/hnsq"
	"service/controller/hwebsocket"
	"service/middleware"
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

	TPLRouter(app)

	ArticleRouter(app)

	MomentRouter(app)

	//获取标签
	app.Get("/api/tag", controller.GetTags)

	UserRouter(app)

	app.Post("/api/comment/:classify", middleware.JWT, controller.AddComment)

	//app.Get("/api/push",controller.Push)

	app.Get("/api/chat/getChat", hwebsocket.GetChat)

	app.Post("/api/nsq", hnsq.Start)

	app.Get("/api/chat/ws", hwebsocket.Chat)
	return app
}
