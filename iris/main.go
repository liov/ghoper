package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"log"
	"net/http"
	"service/controller/cron"
	"service/controller/hwebsocket"
	"service/initialize"
	"service/router"
)

func main() {

	go cron.StartCron()

	go hwebsocket.Start()

	irisRouter := router.IrisRouter()

	// listen and serve on http://0.0.0.0:8080.
	if err := irisRouter.Run(iris.Addr(initialize.Config.Server.HttpPort)); err != nil && err != http.ErrServerClosed {
		log.Printf("Listen: %s\n", err)
	}

}
