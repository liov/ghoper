package router

import (
	"github.com/kataras/iris"
	"hoper/controller/upload"
	"hoper/middleware"
)

func Upload(app *iris.Application) {

	app.Post("/api/upload/exist/{md5:string}", middleware.GetUser(false), upload.MD5)
	app.Post("/api/upload/{classify:string}", middleware.GetUser(false), iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.Upload(ctx)
	})
	app.Post("/api/upload_multiple/{classify:string}", middleware.GetUser(false), iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.UploadMultiple(ctx)
	})
}
