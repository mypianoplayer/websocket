package server

import (
	"log"
	"net/http"
	// "time"
	// "strconv"
	// "math"
//	"sync"
	"golang.org/x/net/websocket"
)

// type MessageReceiver interface {
// 	OnMessage(m *Message)
// }

type Server struct {
	pattern   string
	connections   map[int]*Connection

	AddConnectionCh     chan *Connection
	DeleteConnectionCh     chan *Connection
	SendAllCh chan *Message
	RecvCh chan* Message
	ErrorCh     chan error

	MsgCh chan *Message
}

func New(pattern string, msgch chan *Message) *Server {
	return &Server{
		pattern,
		make(map[int]*Connection),
		make(chan *Connection),
		make(chan *Connection),
		make(chan *Message),
		make(chan *Message),
		make(chan error),
		make(chan *Message),
	}
}

func (s *Server) sendAll(msg *Message) {
	for _, c := range s.connections {
		c.SendCh <- msg
	}
}

func (s * Server) Start() {
	go s.listen()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (s *Server) listen() {

	log.Println("Listening server...")

	onConnected := func(ws *websocket.Conn) {

		defer func() {
			err := ws.Close()
			if err != nil {
				s.ErrorCh <- err
			}
		}()

		connection := NewConnection(ws, s.RecvCh)
		s.AddConnectionCh <- connection
		connection.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))

	for {
		select {

		case c := <-s.AddConnectionCh:
			log.Println("Added new client")
			s.connections[c.id] = c
			log.Println("Now", len(s.connections), "connections connected.")

		case c := <-s.DeleteConnectionCh:
			log.Println("Delete client")
			delete(s.connections, c.id)

		case msg := <-s.RecvCh:
			log.Println("received:")
			s.MsgCh <- msg

		case msg := <-s.SendAllCh:
			s.sendAll(msg)

		case err := <-s.ErrorCh:
			log.Println("Error:", err.Error())
		}
	}
}


