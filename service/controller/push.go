package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Push(c *gin.Context) {
	if pusher := c.Writer.Pusher(); pusher != nil {
		// use pusher.Push() to do server push
		if err := pusher.Push("/assets/app.js", nil); err != nil {
			log.Printf("Failed to push: %v", err)
		}
	}
	c.HTML(200, "https", gin.H{
		"status": "success",
	})
}
