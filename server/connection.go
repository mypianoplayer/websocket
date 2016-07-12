package server

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
	SendCh  chan *Message
	RecvCh  chan *Message
	deleted bool
}

func NewConnection(conn *websocket.Conn, recvCh chan *Message) *Connection {

	if conn == nil {
		panic("conn cannot be nil")
	}

	log.Println("new Connection")

	maxId++
	sendCh := make(chan *Message, channelBufSize)

	return &Connection{maxId, conn, sendCh, recvCh, false}
}

func (c *Connection) Listen() {
	log.Println("conn listen")
	go c.listenWrite()
	c.listenRead()
}

func (c *Connection) listenWrite() {

	for {
		select {

		case msg := <-c.SendCh:
			websocket.JSON.Send(c.conn, msg)

		}
	}
}

func (c *Connection) listenRead() {

	for {
		log.Println("listenRead...")
		select {

		default:
			var msg Message
			err := websocket.JSON.Receive(c.conn, &msg)
			log.Println("msg")
			if err == io.EOF {
				c.deleted = true
				return

			} else if err != nil {
				// c.connectionManager.ErrorCh <- err
			} else {
				c.RecvCh <- &msg
			}
		}
	}
}
