package router

//go:generate qtc -dir=../template
import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"service/controller"
	"service/controller/hnsq"
	"service/controller/hwebsocket"
	"service/middleware"
	"service/template"
)

func InitializeRouter() *router.Router {

	jwt := middleware.JWT
	getUser := middleware.GetUser

	r := router.New()

	/*	store := memstore.NewStore([]byte("hoper_session"))

		r.Use(sessions.Sessions("session_id", store))*/

	//r.Use(middleware.Cors())

	r.GET("/", controller.Index)
	r.GET("/api/a", jwt(controller.Index2))

	//获取文章列表
	r.GET("/api/article", controller.GetArticle)
	//获取指定文章
	r.GET("/api/article/:id", controller.GetArticles)
	//新建文章
	r.POST("/api/article", controller.AddArticle)
	//更新指定文章
	r.PUT("/api/article/:id", controller.EditArticle)
	//删除指定文章
	r.DELETE("/api/article/:id", controller.DeleteArticle)

	//获取文章列表
	r.GET("/api/moment", controller.GetMoments)
	//获取文章列表
	r.GET("/api/moment/:id", getUser(controller.GetMoment))
	//新建文章
	r.POST("/api/moment", jwt(controller.AddMoment))
	//更新指定文章
	r.PUT("/api/moment/:id", jwt(controller.EditMoment))
	//删除指定文章
	r.DELETE("/api/moment/:id", jwt(controller.DeleteMoment))

	//获取标签
	r.GET("/api/tag", controller.GetTags)

	r.GET("/api/user/active/:id/:secret", controller.ActiveAccount)
	r.POST("/api/user/signup", controller.Signup)
	r.POST("/api/user/signin", controller.Signin)
	r.GET("/api/user/signout", jwt(controller.Signout))
	r.POST("/api/user/active", controller.ActiveSendMail)
	r.GET("/api/user/get", jwt(controller.SigninFlag))

	//r.GET("/api/push",controller.Push)

	r.GET("/api/chat/getChat", hwebsocket.GetChat)

	r.POST("/api/nsq", hnsq.Start)

	r.GET("/tpl/index", func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("text/html; charset=utf-8")
		p := &template.IndexPage{
			CTX: ctx,
		}
		template.WritePageTemplate(ctx, p)
	})
	return r
}
