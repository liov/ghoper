package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller"
)

func ArticleRouter(app *iris.Application) {

	articleRouter := app.Party("/api/article")
	{
		//获取文章列表
		articleRouter.Get("", controller.GetArticle)
		//获取指定文章
		articleRouter.Get("/{id:uint64}", controller.GetArticles)
		//新建文章
		articleRouter.Post("", controller.AddArticle)
		//更新指定文章
		articleRouter.Put("/{id:uint64}", controller.EditArticle)
		//删除指定文章
		articleRouter.Delete("/{id:uint64}", controller.DeleteArticle)
	}
}
