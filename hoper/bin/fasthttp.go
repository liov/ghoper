package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"hoper/initialize"
)

func main() {


	fastRouter := FastRouter()

	s := &fasthttp.Server{
		Concurrency:  100,
		Handler:      fastRouter.Handler,
		ReadTimeout:  initialize.Config.Server.ReadTimeout,
		WriteTimeout: initialize.Config.Server.WriteTimeout,
		LogAllErrors: true,
	}

	go func() {
		if err := s.ListenAndServe(initialize.Config.Server.HttpPort); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")


	if err := s.Shutdown(); err != nil {
		log.Fatal("fasthttp Server Shutdown:", err)
	}

	log.Println("Server exiting")
}

func FastRouter() *router.Router {

	r := router.New()

	return r
}
