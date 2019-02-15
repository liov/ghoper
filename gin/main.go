package main

import (
	"context"
	"ginService/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	httpRouter := router.HttpRouter()
	gin := &http.Server{
		Addr:           ":8080",
		Handler:        httpRouter,
		ReadTimeout:    1024,
		WriteTimeout:   1024,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := gin.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := gin.Shutdown(ctx); err != nil {
		log.Fatal("gin Server Shutdown:", err)
	}
}
