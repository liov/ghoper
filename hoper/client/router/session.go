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

	sessRouter := app.Party("/api/sess")

	sessRouter.Get("/set", controller.SessSet)

	sessRouter.Get("/get", controller.SessGet)

	sessRouter.Get("/hoper", controller.SessTest)

	sessRouter.Get("/delete", controller.SessDelete)

	sessRouter.Get("/clear", controller.SessClear)

	sessRouter.Get("/destroy", controller.SessDestroy)

	sessRouter.Get("/update", controller.SessUpdate)

	//sessRouter.Get("/gset", controller.GsessSet)

	//sessRouter.Get("/gget", controller.GsessGet)
}
