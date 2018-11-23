package hwebsocket

import (
	"encoding/json"
	"github.com/fasthttp/websocket"
	"github.com/gomodule/redigo/redis"
	"github.com/satori/go.uuid"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"os"
	"service/controller"
	"service/controller/common"
	"service/initialize"
	"service/utils"
	"strings"
	"time"
)

type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte //广播聊天
	register   chan *Client
	unregister chan *Client
}

type Client struct {
	uuid   string
	conn   *websocket.Conn
	send   chan []byte
	device string
}

type Message struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	//SendUser	controller.User `gorm:"ForeignKey:SenderUserID" json:"send_user"`
	SenderUserID uint `json:"sender_user_id,omitempty"`
	//RecipientUser	controller.User `gorm:"ForeignKey:RecipientUserID" json:"recipient_user"`
	RecipientUserID uint   `json:"recipient_user_id,omitempty"`
	Content         string `json:"content,omitempty"`
	Remarks         string `json:"remarks,omitempty"`
}

type SendMessage struct {
	ID            uint            `gorm:"primary_key" json:"id"`
	CreatedAt     time.Time       `json:"created_at"`
	SendUser      controller.User `gorm:"ForeignKey:SenderUserID" json:"send_user"`
	RecipientUser controller.User `gorm:"ForeignKey:RecipientUserID" json:"recipient_user"`
	Content       string          `json:"content"`
	Remarks       string          `json:"remarks"`
	Device        string          `json:"device"`
}

type ReceiveMessage struct {
	ID              uint      `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	SenderUserID    uint      `json:"sender_user_id,omitempty"`
	RecipientUserID uint      `json:"recipient_user_id,omitempty"`
	Content         string    `json:"content,omitempty"`
	Remarks         string    `json:"remarks,omitempty"`
}

var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

var upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.register:
			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Remarks: "/A new conn has connected."})
			manager.send(jsonMessage, conn)
		case conn := <-manager.unregister:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Remarks: "/A conn has disconnected."})
				manager.send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:
			//这里貌似可以做单点发送
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func Start() {
	manager.start()
}

func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		if conn != ignore {
			conn.send <- message
		}
	}
}

func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			}
			break
		}
		var receiveMessage ReceiveMessage
		json.Unmarshal([]byte(msg), &receiveMessage)
		receiveMessage.CreatedAt = time.Now()
		sendUser, _ := controller.UserFromRedis(int(receiveMessage.SenderUserID))
		sendMessage := SendMessage{
			ID:        receiveMessage.ID,
			CreatedAt: receiveMessage.CreatedAt,
			SendUser:  sendUser,
			//RecipientUser:nil,
			Content: receiveMessage.Content,
			Remarks: receiveMessage.Remarks,
			Device:  c.device,
		}
		jsonMessage, _ := json.Marshal(&sendMessage)
		MsgRedis(jsonMessage)
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func Chat(c *fasthttp.RequestCtx) {
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		var dviceName string
		if strings.Contains(utils.ToSting(c.Request.Header.UserAgent()), "iPhone") {
			dviceName = "iPhone"
		} else if strings.Contains(utils.ToSting(c.Request.Header.UserAgent()), "Android") {
			dviceName = "Android"
		} else {
			dviceName = "PC"
		}
		client := &Client{uuid: uuid.NewV4().String(), conn: conn, send: make(chan []byte), device: dviceName}
		manager.register <- client

		go client.write()
		client.read()
	})
	if err != nil {
		log.Println(err)
	}
}

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 10 * time.Second
)

var filename string

func readFileIfModified(lastMod time.Time) ([]byte, time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, lastMod, err
	}
	if !fi.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}
	p, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fi.ModTime(), err
	}
	return p, fi.ModTime(), nil
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func writer(ws *websocket.Conn, lastMod time.Time) {
	lastError := ""
	pingTicker := time.NewTicker(pingPeriod)
	fileTicker := time.NewTicker(filePeriod)
	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()
	for {
		select {
		case <-fileTicker.C:
			var p []byte
			var err error

			p, lastMod, err = readFileIfModified(lastMod)

			if err != nil {
				if s := err.Error(); s != lastError {
					lastError = s
					p = []byte(lastError)
				}
			} else {
				lastError = ""
			}

			if p != nil {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.WriteMessage(websocket.TextMessage, p); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func MsgRedis(data []byte) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	conn.Do("RPUSH", "Chat", data)
}

func GetChat(c *fasthttp.RequestCtx) {
	conn := initialize.RedisPool.Get()
	defer conn.Close()
	data, err := redis.ByteSlices(conn.Do("LRANGE", "Chat", 0, -1))
	if err != nil {
		return
	}
	var messages []SendMessage

	for _, v := range data {
		var message SendMessage
		common.Json.Unmarshal(v, &message)
		messages = append(messages, message)
	}
	common.Response(c, messages)
}
