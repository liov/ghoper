package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller/upload"
	"hoper/client/middleware"
)

func Upload(app *iris.Application) {

	app.Post("/api/upload/exist/{md5:string}", middleware.JWT, upload.MD5)
	app.Post("/api/upload/{classify:string}", middleware.JWT, iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.Upload(ctx)
	})
	app.Post("/api/upload_multiple/{classify:string}", middleware.JWT, iris.LimitRequestBodySize(10<<20), func(ctx iris.Context) {
		upload.UploadMultiple(ctx)
	})
}
