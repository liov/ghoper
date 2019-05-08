package main

import (
	"github.com/gin-gonic/gin"
	"hoper/client"
	"hoper/initialize"
	"net/http"
)

func main() {

	//go server.Service()
	client.Client()
}

//查看日志
//tail -n 10 *.log
/**
*发现了好玩的地方,一个go就这么好玩
*js也挺好玩的，ts不知道好不好玩
*java除了生态还是有优点的，热更新，但是吧我觉得启动比编译+启动时间久呢
*Rust一定更好玩了
*Python只能归为实用类语言了
 */

func Gin()  {
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
}
