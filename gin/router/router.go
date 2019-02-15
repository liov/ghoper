package router

//go:generate qtc -dir=../template
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpRouter() *gin.Engine {

	gin.SetMode("release")

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

	return r
}
