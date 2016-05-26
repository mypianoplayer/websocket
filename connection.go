package ragtime

import (
	// "fmt"
	"io"
	"log"
	"golang.org/x/net/websocket"
)

const channelBufSize = 100

var maxId int = 0

type Connection struct {
	id     int
	conn   *websocket.Conn
	server *GameServer
	SendMessageCh  chan *Message
}

func NewConnection(conn *websocket.Conn, server *GameServer) *Connection {

	if conn == nil {
		panic("conn cannot be nil")
	}

	if server == nil {
		panic("server cannot be nil")
	}

	log.Println("new Connection")

	maxId++
	ch := make(chan *Message, channelBufSize)

	return &Connection{maxId, conn, server, ch}
}

// func (c *Connection) Conn() *websocket.Conn {
// 	return c.conn
// }

// func (c *Connection) Write(msg *Message) {
// 	select {
// 	case c.SendMessageCh <- msg:
// 	default:
// 		c.server.DeleteConnectionCh <- c
// 		err := fmt.Errorf("Connection %d is disconnected.", c.id)
// 		c.server.ErrorCh <- err
// 	}
// }

func (c *Connection) Listen() {
	go c.listenWrite()
	c.listenRead()
}

func (c *Connection) listenWrite() {

	for {
		select {

		case msg := <-c.SendMessageCh:
			websocket.JSON.Send(c.conn, msg)

		}
	}
}

func (c *Connection) listenRead() {

	for {
		select {

		default:
			var msg Message
			err := websocket.JSON.Receive(c.conn, &msg)
			if err == io.EOF {
				c.server.DeleteConnectionCh <- c
				return

			} else if err != nil {
				c.server.ErrorCh <- err
			} else {
				c.server.SendAllCh <- &msg
			}
		}
	}
}
