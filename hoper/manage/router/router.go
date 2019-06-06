package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"net/http"
)

func GinRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	//r.Use(gin.Logger())

	r.Use(gin.Recovery())


	v1 := r.Group("/gin")

	v1.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "gin",
		})
	})
	//当前目录为工作目录
	box := packr.New("packr","../../../static/template/packr")

	v1.StaticFS("/packr",http.FileSystem(box))

	return r
}
