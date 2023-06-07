package router

import (
	"github.com/kataras/iris"
	"hoper/controller/hwebsocket"
	"hoper/middleware"
)

func WS(app *iris.Application) {

	app.Get("/api/chat/getChat", middleware.GetUser(false), hwebsocket.GetChat)

	app.Get("/ws/chat", middleware.GetUser(true), hwebsocket.Chat)

	ws := hwebsocket.GetWebsocket()

	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html, this endpoint is used to connect to the server.
	app.Get("/ws/echo", middleware.GetUser(true), ws.Handler())

	// serve the javascript built'n client-side library,
	// see websockets.html script tags, this path is used.
	app.Any("/api/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(ws.ClientSource)
	})
}
