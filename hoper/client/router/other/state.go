package other

import "github.com/kataras/iris"

func State(app *iris.Application) {
	none := app.None("/api/invisible/{username}", func(ctx iris.Context) {
		ctx.Writef("Hello %s with method: %s", ctx.Params().Get("username"), ctx.Method())

		if from := ctx.Values().GetString("from"); from != "" {
			ctx.Writef("\nI see that you're coming from %s", from)
		}
	})

	app.Get("/api/change", func(ctx iris.Context) {

		if none.IsOnline() {
			none.Method = iris.MethodNone
		} else {
			none.Method = iris.MethodGet
		}

		// refresh re-builds the router at serve-time in order to be notified for its new routes.
		app.RefreshRouter()
	})

	app.Get("/api/execute", func(ctx iris.Context) {
		// same as navigating to "http://localhost:8080/invisible/iris" when /change has being invoked and route state changed
		// from "offline" to "online"
		ctx.Values().Set("from", "/api/execute") // values and session can be shared when calling Exec from a "foreign" context.
		// 	ctx.Exec("NONE", "/invisible/iris")
		// or after "/change":
		ctx.Exec("GET", "/api/invisible/iris")
	})
}
