package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"hoper/controller/tmpl/html"
	"hoper/controller/tmpl/markdown"
	"hoper/controller/tmpl/pug"
	"time"
)

func TPLRouter(app *iris.Application) {

	tmpl := iris.HTML("../../..static/template", ".html").Reload(true)

	tmplPug := iris.Pug("../../..static/template", ".pug").Reload(true)

	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)
	app.RegisterView(tmplPug)

	//缓存10s
	tplRouter := app.Party("/api/tpl")
	{
		//这里的pprof有问题，访问profile返回的是文件
		tplRouter.Get("/hi", cache.Handler(10*time.Second), html.HtmlTest)
		tplRouter.Get("/pug", pug.PugTest)
		tplRouter.Get("/markdown", markdown.MarkdownTest)
		tplRouter.Get("/time", iris.Cache304(10*time.Second), html.Time)
		tplRouter.Get("/auth", html.Auth)
		tplRouter.Get("/values", html.Values)
		tplRouter.Get("/logout/{provider}", func(ctx iris.Context) {
			ctx.Redirect("/", iris.StatusTemporaryRedirect)
		})

		tplRouter.Get("/ws", func(ctx iris.Context) {
			// try to get the user without re-authenticating
			if err := ctx.View("ws.html"); err != nil {
				ctx.Writef("%v", err)
			}
		})

		tplRouter.Get("/wasm", func(ctx iris.Context) {
			// try to get the user without re-authenticating
			if err := ctx.View("wasm.html"); err != nil {
				ctx.Writef("%v", err)
			}
		})

	}

}
