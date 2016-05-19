package ragtime

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const channelBufSize = 100

var maxId int = 0

// Chat Connection.
type Connection struct {
	id     int
	ws     *websocket.Conn
	server *GameServer
	ch     chan *Message
	doneCh chan bool
}

// Create new chat Connection.
func NewConnection(ws *websocket.Conn, server *GameServer) *Connection {

	if ws == nil {
		panic("ws cannot be nil")
	}

	if server == nil {
		panic("server cannot be nil")
	}

	log.Println("new Connection")

	maxId++
	ch := make(chan *Message, channelBufSize)
	doneCh := make(chan bool)

	return &Connection{maxId, ws, server, ch, doneCh}
}

func (c *Connection) Conn() *websocket.Conn {
	return c.ws
}

func (c *Connection) Write(msg *Message) {
	select {
	case c.ch <- msg:
	default:
		c.server.Del(c)
		err := fmt.Errorf("Connection %d is disconnected.", c.id)
		c.server.Err(err)
	}
}

func (c *Connection) Done() {
	c.doneCh <- true
}

// Listen Write and Read request via chanel
func (c *Connection) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Connection) listenWrite() {
	log.Println("Listening write to Connection")
	for {
		select {

		// send message to the Connection
		case msg := <-c.ch:
//			log.Println("Send:", msg)
			websocket.JSON.Send(c.ws, msg)

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *Connection) listenRead() {
	log.Println("Listening read from Connection")
	for {
		select {

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenWrite method
			return

		// read data from websocket connection
		default:
			var msg Message
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.server.Err(err)
			} else if msg.Body == "close" {
				c.doneCh <- true
			} else {
				c.server.SendAll(&msg)
			}
		}
	}
}
