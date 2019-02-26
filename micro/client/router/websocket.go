package router

import (
	"github.com/kataras/iris"
	"hoper/client/controller/hwebsocket"
	"hoper/client/middleware"
)

func WS(app *iris.Application) {
	app.Get("/ws/chat", hwebsocket.Chat)

	ws := hwebsocket.GetWebsocket()

	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html, this endpoint is used to connect to the server.
	app.Get("/ws/echo", middleware.JWT, ws.Handler())

	// serve the javascript built'n client-side library,
	// see websockets.html script tags, this path is used.
	app.Any("/tpl/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(ws.ClientSource)
	})
}
