package controller

import (
	"log"

	"github.com/kataras/iris"
)

//可用的，前提是页面需要
func Push(c iris.Context) {
	if pusher := c.ResponseWriter(); pusher != nil {
		// use pusher.Push() to do server push
		if err := pusher.Push("/api/static/images/6cbeb5c8-7160-4b6f-a342-d96d3c00367a.jpg", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	c.HTML(`<html><body><img src="/api/static/images/6cbeb5c8-7160-4b6f-a342-d96d3c00367a.jpg"></body></html>`)
}
