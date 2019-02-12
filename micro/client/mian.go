package main

import (
	"github.com/kataras/iris"
	"log"
	"micro/client/router"
	"micro/common/controller/cron"
	"net/http"
)

func main() {

	cron.New().Start()
	defer cron.New().Stop()

	app := router.Router()

	// listen and serve on http://0.0.0.0:8080.
	if err := app.Run(iris.Addr(":8000")); err != nil && err != http.ErrServerClosed {
		log.Printf("Listen: %s\n", err)
	}

}
