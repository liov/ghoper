package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
	"hoper/client/middleware"
)

func ArticleRouter(app *iris.Application) {

	//获取标签
	app.Get("/api/articleSet", controller.GetTags)

	articleRouter := app.Party("/api/article")
	{
		//获取文章列表
		articleRouter.Get("", controller.GetArticles)
		//获取指定文章
		articleRouter.Get("/{id:uint64}", controller.GetArticle)
		//新建文章
		articleRouter.Post("", middleware.GetUser(), controller.AddArticle)
		//更新指定文章
		articleRouter.Put("/{id:uint64}", middleware.Login, controller.EditArticle)
		//删除指定文章
		articleRouter.Delete("/{id:uint64}", middleware.Login, controller.DeleteArticle)
	}
}
