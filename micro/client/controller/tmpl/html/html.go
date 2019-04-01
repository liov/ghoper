package html

import (
	"github.com/kataras/iris"
	"hoper/initialize"
	"hoper/model/ov"
	"time"
)

func HtmlTest(ctx iris.Context) {
	ctx.Application().Logger().Infof("Request path: %s", ctx.Path())
	//源码如果ViewData的key为"",那么只能设置一次值，第二次设置key""会覆盖第一次的值，
	//其后设置key不为空的值无效
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

func Values(ctx iris.Context) {
	user := ov.User{Name: ctx.URLParam("name")}
	ctx.Values().Set("a", user)

	initialize.Cache.Set(ctx.URLParam("key"), user)

	user2, _ := initialize.Cache.Get("a")
	ctx.JSON(iris.Map{
		"user1": ctx.Values().Get("a"),

		"user2": user2,
	})

}
