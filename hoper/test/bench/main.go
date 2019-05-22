package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kataras/iris"
	"net/http"
)

func Gin() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	
	r.Run(":8000")
}

func Iris(){
	app := iris.New()
    app.Get("/ping", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "pong",
        })
    })
    // listen and serve on http://0.0.0.0:8080.
    app.Run(iris.Addr(":8000"))
}

func main(){
Iris()
}
