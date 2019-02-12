package html

import "github.com/kataras/iris"

func HtmlTest(ctx iris.Context) {
	ctx.Application().Logger().Infof("Request path: %s", ctx.Path())
	ctx.ViewData("Title", "Hi Page")
	ctx.ViewData("Name", "iris")
	ctx.View("hi.html")
}
