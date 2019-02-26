package hwebsocket

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

func SetupWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		// These are low-level optionally fields,
		// user/client can't see those values.
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// only javascript client-side code has the same rule,
		// which you serve using the ws.ClientSource (see below).
		EvtMessagePrefix: []byte("JYB:"),
	})
	ws.OnConnection(handleConnection)

	app.Get("/tpl/ws", func(ctx iris.Context) {
		// try to get the user without re-authenticating
		if err := ctx.View("ws.html"); err != nil {
			ctx.Writef("%v", err)
		}
	})

	// register the server on an endpoint.
	// see the inline javascript code in the websockets.html, this endpoint is used to connect to the server.
	app.Get("/ws/echo", ws.Handler())

	// serve the javascript built'n client-side library,
	// see websockets.html script tags, this path is used.
	app.Any("/tpl/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(ws.ClientSource)
	})
}
func handleConnection(c websocket.Connection) {
	// Read events from browser
	c.On("chat", func(msg string) {
		// Print the message to the console, c.Context() is the iris's http context.
		// Write message back to the client message owner with:
		// c.Emit("chat", msg)
		// Write message to all except this client with:
		c.To(websocket.Broadcast).Emit("chat", msg)
	})
}
