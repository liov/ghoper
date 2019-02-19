package router

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
	"github.com/kataras/iris/middleware/basicauth"
	"github.com/kataras/iris/middleware/pprof"
	"hoper/client/controller/tmpl/html"
	"hoper/client/controller/tmpl/markdown"
	"hoper/client/controller/tmpl/pug"
	"hoper/client/middleware"
	"time"
)

func TPLRouter(app *iris.Application) {

	tmpl := iris.HTML("./client/template", ".html").Reload(true)

	tmplPug := iris.Pug("./client/template", ".pug").Reload(true)

	tmpl.AddFunc("greet", func(s string) string {
		return "Greetings " + s + "!"
	})

	app.RegisterView(tmpl)
	app.RegisterView(tmplPug)

	//auth
	authConfig := basicauth.Config{
		Users:   map[string]string{"admin": "lby604"},
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	//缓存10s
	tplRouter := app.Party("/tpl")
	{
		tplRouter.Any("/tpl/pprof/{action:path}", authentication, pprof.New())
		tplRouter.Get("/hi", cache.Handler(10*time.Second), html.HtmlTest)
		tplRouter.Get("/pug", pug.PugTest)
		tplRouter.Get("/markdown", markdown.MarkdownTest)
		tplRouter.Get("/time", iris.Cache304(10*time.Second), html.Time)
		tplRouter.Get("/auth", authentication, html.Auth)
		tplRouter.Get("/values", authentication, html.Values)
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
		tplRouter.Get("/auth/{provider}/callback", func(ctx iris.Context) {

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
		tplRouter.Get("/", func(ctx iris.Context) {

			ctx.ViewData("", middleware.NewAuth())

			if err := ctx.View("index.html"); err != nil {
				ctx.Writef("%v", err)
			}
		})
	}
}