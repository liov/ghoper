package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"os/signal"
	"service/controller/hwebsocket"
	"service/initialize"
	"service/router"
)

func main() {

	//go cron.CronStart(cron.GetCron())

	go hwebsocket.Start()

	router := router.InitializeRouter()

	/*	s := &http.Server{
			Addr:           fmt.Sprintf(":%d", initialize.ServerSettings.HttpPort),
			Handler:        routersInit,
			ReadTimeout:    initialize.ServerSettings.ReadTimeout,
			WriteTimeout:   initialize.ServerSettings.WriteTimeout,
			MaxHeaderBytes: 1 << 20,
		}

		go func() {
			if err := s.ListenAndServe(); err != nil&& err != http.ErrServerClosed {
				log.Printf("Listen: %s\n", err)
			}
		}()
	*/

	s := &fasthttp.Server{
		Concurrency:  100,
		Handler:      router.Handler,
		ReadTimeout:  initialize.ServerSettings.ReadTimeout,
		WriteTimeout: initialize.ServerSettings.WriteTimeout,
		LogAllErrors: true,
	}

	go func() {
		if err := s.ListenAndServe(initialize.ServerSettings.HttpPort); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	//log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")
	if err := s.Shutdown(); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
