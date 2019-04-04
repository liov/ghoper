package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/initialize"
)

func Session(app *iris.Application) {

	iris.RegisterOnInterrupt(func() {
		initialize.BoltDB.Close()
	})

	app.Get("/set", controller.SessSet)

	app.Get("/get", controller.SessGet)

	app.Get("/hoper", controller.SessTest)

	app.Get("/delete", controller.SessDelete)

	app.Get("/clear", controller.SessClear)

	app.Get("/destroy", controller.SessDestroy)

	app.Get("/update", controller.SessUpdate)

	//app.Get("/gset", controller.GsessSet)

	//app.Get("/gget", controller.GsessGet)
}
