package router

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
	"hoper/client/controller"
	"hoper/client/controller/hnsq"
	"os"
	"time"
)

/**
 * @author     ：lbyi
 * @date       ：Created in 2019/4/2
 * @description：
 */

func System(app *iris.Application) {

	//auth
	authConfig := basicauth.Config{
		Users:   map[string]string{"admin": "lby604"},
		Realm:   "Authorization Required", // defaults to "Authorization Required"
		Expires: time.Duration(30) * time.Minute,
	}

	authentication := basicauth.New(authConfig)

	app.Get("/api/init", authentication, controller.DBInit)

	app.Post("/api/nsq", hnsq.Start)

	app.Get("/api/restart", func(c iris.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		app.Shutdown(ctx)
		c.WriteString("重启了")
	})
	app.Get("/api/shutdown", func(c iris.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		Ch <- os.Kill
		app.Shutdown(ctx)

	})
}
