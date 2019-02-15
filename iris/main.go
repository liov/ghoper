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

	cron.New().Start()
	defer cron.New().Stop()

	go hwebsocket.Start()

	//go hcache.Start()

	irisRouter := router.IrisRouter()

	// listen and serve on http://0.0.0.0:8080.
	if err := irisRouter.Run(iris.Addr(initialize.Config.Server.HttpPort),
		iris.WithConfiguration(iris.YAML("../config/iris.yml"))); err != nil && err != http.ErrServerClosed {
		log.Printf("Listen: %s\n", err)
	}

	/*	opts := groupcache.HTTPPoolOptions{BasePath: hcache.BasePath}
		peers := groupcache.NewHTTPPoolOpts("", &opts)
		peers.Set("http://localhost:8333", "http://localhost:8222")

		val, err := hcache.GetFromPeer("helloworld", "wjs1", peers)*/

}
