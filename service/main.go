package main

import (
	"context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"service/controller/hwebsocket"
	"service/initialize"
	"service/router"
	"time"
)

func main() {

	//go cron.CronStart(cron.GetCron())

	go hwebsocket.Start()

	fastRouter := router.FastRouter()

	s := &fasthttp.Server{
		Concurrency:  100,
		Handler:      fastRouter.Handler,
		ReadTimeout:  initialize.ServerSettings.ReadTimeout,
		WriteTimeout: initialize.ServerSettings.WriteTimeout,
		LogAllErrors: true,
	}

	go func() {
		if err := s.ListenAndServe(initialize.ServerSettings.HttpPort); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	httpRouter := router.HttpRouter()

	ws := &http.Server{
		Addr:           ":8080",
		Handler:        httpRouter,
		ReadTimeout:    initialize.ServerSettings.ReadTimeout,
		WriteTimeout:   initialize.ServerSettings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := ws.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ws.Shutdown(ctx); err != nil {
		log.Fatal("ws Server Shutdown:", err)
	}
	if err := s.Shutdown(); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
