package hwebsocket

import (
	"github.com/kataras/iris/websocket"
	"github.com/satori/go.uuid"
	"hoper/client/controller"
	"hoper/client/controller/common"
	"time"
)

type ClientManagerI struct {
	clients    map[*ClientI]bool
	broadcast  chan []byte //广播聊天
	register   chan *ClientI
	unregister chan *ClientI
}

type ClientI struct {
	uuid   string
	conn   websocket.Connection
	send   chan []byte
	device string
}

var managerI = ClientManagerI{
	clients:    make(map[*ClientI]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *ClientI),
	unregister: make(chan *ClientI),
}

func GetWebsocket() *websocket.Server {
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

	return ws
}
func handleConnection(c websocket.Connection) {
	// Read events from browser
	c.On("chat", func(msg string) {
		// Print the message to the console, c.Context() is the iris's http context.
		// Write message back to the client message owner with:

		var deviceName string
		/*			if strings.Contains(c.Context().Request().UserAgent(), "iPhone") {
						dviceName = "iPhone"
					} else if strings.Contains(c.Context().Request().UserAgent(), "Android") {
						dviceName = "Android"
					} else {
						dviceName = "PC"
					}*/
		client := &ClientI{uuid: uuid.NewV4().String(), conn: c, send: make(chan []byte), device: deviceName}

		managerI.clients[client] = true
		user := c.Context().Values().Get("user").(controller.User)
		var receiveMessage ReceiveMessage
		common.Json.UnmarshalFromString(msg, &receiveMessage)
		sendMessage := SendMessage{
			ID:        receiveMessage.ID,
			CreatedAt: time.Now(),
			SendUser:  user,
			//RecipientUser:nil,
			Content: receiveMessage.Content,
			Remarks: receiveMessage.Remarks,
			Device:  deviceName,
		}

		ms, _ := common.Json.Marshal(sendMessage)
		MsgRedis(ms)
		c.Emit("chat", ms)
		// Write message to all except this client with:
		//c.To(websocket.Broadcast).Emit("chat",msg)
	})
}
