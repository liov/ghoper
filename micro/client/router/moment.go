package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/middleware"
)

func MomentRouter(app *iris.Application) {

	momentRouter := app.Party("/api/moment")
	{
		//获取文章列表
		momentRouter.Get("/", middleware.GetUserId, controller.GetMoments)
		//获取文章列表
		momentRouter.Get("/{id:uint64}", middleware.GetUserId, controller.GetMoment)
		//新建文章
		momentRouter.Post("", middleware.GetUser(), controller.AddMoment)
		//更新指定文章
		momentRouter.Put("/{id:uint64}", middleware.Login, controller.EditMoment)
		//删除指定文章
		momentRouter.Delete("/{id:uint64}", middleware.Login, controller.DeleteMoment)
	}
}
