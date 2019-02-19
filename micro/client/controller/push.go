package controller

import (
	"github.com/kataras/iris"
	"log"
)

func Push(c iris.Context) {
	if pusher := c.ResponseWriter(); pusher != nil {
		// use pusher.Push() to do server push
		if err := pusher.Push("/assets/app.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	c.HTML("status:success")
}
