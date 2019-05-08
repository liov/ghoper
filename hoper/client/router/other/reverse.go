package other

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
)

func Reverse(app *iris.Application) {
	rv := router.NewRoutePathReverser(app)

	myroute := app.Get("/api/reverse/anything/{anythingparameter:path}", func(ctx iris.Context) {
		paramValue := ctx.Params().Get("anythingparameter")
		ctx.Writef("The path after /anything is: %s", paramValue)
	})

	myroute.Name = "myroute"

	// useful for links, although iris' view engine has the {{ urlpath "routename" "path values"}} already.
	app.Get("/api/reverse/reverse_myroute", func(ctx iris.Context) {
		myrouteRequestPath := rv.Path(myroute.Name, "any/path")
		ctx.HTML("Should be <b>/anything/any/path</b>: " + myrouteRequestPath)
	})

	// execute a route, similar to redirect but without redirect :)
	app.Get("/api/reverse/execute_myroute", func(ctx iris.Context) {
		ctx.Exec("GET", "/api/reverse/anything/any/path") // like it was called by the client.
	})

}
