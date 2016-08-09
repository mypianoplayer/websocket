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
	server *Server
}

func NewConnection(conn *websocket.Conn, s *Server) *Connection {

	if conn == nil {
		panic("conn cannot be nil")
	}

	log.Println("new Connection")

	maxId++
	sendCh := make(chan *Message, channelBufSize)

	return &Connection{
		id:maxId,
		conn:conn,
		sendCh:sendCh,
		server:s}
}

func (c *Connection) SendCh() chan *Message {
	return c.sendCh
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

	log.Println("listenRead...")
	for {
		select {

		default:
			var msg Message
			err := websocket.JSON.Receive(c.conn, &msg)
			msg.ID = c.id
			if err == io.EOF {
				c.server.DelConnCh() <- c
				return

			} else if err != nil {
				log.Println(err)
				// c.connectionManager.ErrorCh <- err
			} else {
				log.Println("msg")
				c.server.RecvCh() <- &msg
			}
		}
	}
}
