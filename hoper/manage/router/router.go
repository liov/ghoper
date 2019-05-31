package router

import (
	"github.com/gin-gonic/gin"
	"hoper/initialize"
	"net/http"
)

func GinRouter() *gin.Engine {
	gin.SetMode(initialize.Config.Server.Env)

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	v1 := r.Group("/gin")

	v1.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "gin",
		})
	})

	return r
}
