package main

import (
	"github.com/kataras/iris"
	"hoper/client/controller/cron"
	"hoper/client/controller/hwebsocket"
	"hoper/client/router"
	"hoper/initialize"
	"log"
	"net/http"
)

func main() {

	//crawler.M131()
	//go crawler.MM131()

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
