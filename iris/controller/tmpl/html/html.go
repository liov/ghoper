package html

import (
	"github.com/kataras/iris"
	"time"
)

func HtmlTest(ctx iris.Context) {
	ctx.Application().Logger().Infof("Request path: %s", ctx.Path())
	ctx.ViewData("Title", "Hi Page")
	ctx.ViewData("Name", "iris")
	ctx.View("hi.html")
}

func Time(ctx iris.Context) {
	ctx.Header("X-Custom", "my  custom header")
	ctx.Writef("Hello World! %s", time.Now())
}

func Auth(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()

	ctx.Writef("%s:%s", username, password)
}
