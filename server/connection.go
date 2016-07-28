package server

import (
	// "fmt"
	"golang.org/x/net/websocket"
	"io"
	"log"
)

const channelBufSize = 100

var maxId int = 0

type Connection struct {
	id      int
	conn    *websocket.Conn
	sendCh  chan *Message
	recvCh  chan *Message
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

func (c *Connection) SendCh() chan *Message {
	return c.sendCh
}

func (c *Connection) RecvCh() chan *Message {
	return c.recvCh
}

func (c *Connection) Listen() {
	log.Println("conn listen")
	go c.listenWrite()
	c.listenRead()
}

func (c *Connection) listenWrite() {

	for {
		select {

		case msg := <-c.sendCh:
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
			if err == io.EOF {
				c.deleted = true
				return

			} else if err != nil {
				log.Println(err)
				// c.connectionManager.ErrorCh <- err
			} else {
				log.Println("msg")
				c.recvCh <- &msg
			}
		}
	}
}
