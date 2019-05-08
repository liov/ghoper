package router

//go:generate qtc -dir=../template
import (
	"github.com/fasthttp/router"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/middleware"
	"github.com/valyala/fasthttp"
	"hoper/client/controller"
	"hoper/initialize"
	"io"
	"net/http"
)

func FastRouter() *router.Router {

	r := router.New()

	/*	store := memstore.NewStore([]byte("hoper_session"))

		r.Use(sessions.Sessions("session_id", store))*/

	//r.Use(middleware.Cors())


	r.GET("/tpl/index", func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("text/html; charset=utf-8")
		p := &template.IndexPage{
			CTX: ctx,
		}
		template.WritePageTemplate(ctx, p)
	})

	return r
}



