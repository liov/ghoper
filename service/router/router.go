package router

//go:generate qtc -dir=../template
import (
	"github.com/fasthttp/router"
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
	"service/controller"
	"service/controller/common/logging"
	"service/controller/hnsq"
	"service/controller/hwebsocket"
	"service/initialize"
	"service/middleware"
	"service/template"
)

func FastRouter() *router.Router {

	jwt := middleware.JWT
	getUser := middleware.GetUser

	r := router.New()

	/*	store := memstore.NewStore([]byte("hoper_session"))

		r.Use(sessions.Sessions("session_id", store))*/

	//r.Use(middleware.Cors())

	r.GET("/api", controller.Index)

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

	r.POST("/api/comment/:classify", jwt(controller.AddComment))

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

func HttpRouter() *gin.Engine {
	logFile := logging.GetIOWrite()

	gin.SetMode(initialize.Config.Server.Env)

	gin.DefaultWriter = io.MultiWriter(logFile)

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.LoadHTMLGlob("template/*")

	r.GET("/gin/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.Static("/gin/static", "../static")
	v1 := r.Group("/gin")
	v1.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "gin",
		})
	})
	//r.GET("/api/chat/ws", hwebsocket.Chat)

	r.GET("/api/push", controller.Push)
	return r
}

func IrisRouter() *iris.Application {
	app := iris.New()

	app.StaticWeb("/iris/static", "../static")

	v1 := app.Party("/iris")
	{
		v1.Get("/ping", func(ctx iris.Context) {
			ctx.JSON(iris.Map{
				"message": "iris",
			})
		})

	}

	app.Macros().Get("string").RegisterFunc("range", func(minLength, maxLength int) func(string) bool {
		return func(paramValue string) bool {
			return len(paramValue) >= minLength && len(paramValue) <= maxLength
		}
	})

	app.Get("/limitchar/{name:string range(1,200) else 400}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | the name should be between 1 and 200 characters length
    otherwise this handler will not be executed`, name)
	})

	app.Macros().Get("string").RegisterFunc("has", func(validNames []string) func(string) bool {
		return func(paramValue string) bool {
			for _, validName := range validNames {
				if validName == paramValue {
					return true
				}
			}

			return false
		}
	})

	app.Get("/static_validation/{name:string has([kataras,gerasimos,maropoulos])}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef(`Hello %s | the name should be "kataras" or "gerasimos" or "maropoulos"
    otherwise this handler will not be executed`, name)
	})

	app.Get("/api/chat/ws", hwebsocket.Chat)
	return app
}
