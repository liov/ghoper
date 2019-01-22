package main

import (
	"github.com/kataras/iris"
	"log"
	"net/http"
)

func main() {
	app := iris.Default()

	app.Get("/graphql", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"message": "iris",
		})
	})

	// listen and serve on http://0.0.0.0:8080.
	if err := app.Run(iris.Addr(":8888")); err != nil && err != http.ErrServerClosed {
		log.Printf("Listen: %s\n", err)
	}

}
