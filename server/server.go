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

type Server struct {
	pattern     string
	connections map[int]*Connection

	addConnCh chan *Connection
	delConnCh chan *Connection
	sendAllCh chan *Message
	recvCh    chan *Message
	errCh     chan error
}

func (s *Server) AddConnCh() chan *Connection {
	return s.addConnCh
}

func (s *Server) DelConnCh() chan *Connection {
	return s.delConnCh
}

func (s *Server) SendAllCh() chan *Message {
	return s.sendAllCh
}

func (s *Server) RecvCh() chan *Message {
	return s.recvCh
}


func New(pattern string) *Server {
	return &Server{
		pattern,
		make(map[int]*Connection),
		make(chan *Connection),
		make(chan *Connection),
		make(chan *Message),
		make(chan *Message),
		make(chan error),
	}
}

func (s *Server) sendAll(msg *Message) {
	for _, c := range s.connections {
		c.SendCh() <- msg
	}
}

func (s *Server) Start() {
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
				s.errCh <- err
			}
		}()

		connection := NewConnection(ws, s.recvCh)
		s.addConnCh <- connection
		connection.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))

	for {
		select {

		case c := <-s.addConnCh:
			log.Println("Added new client")
			s.connections[c.id] = c
			log.Println("Now", len(s.connections), "connections connected.")

		case c := <-s.delConnCh:
			log.Println("Delete client")
			delete(s.connections, c.id)

		// case msg := <-s.recvCh:
		// 	log.Println("received:")
		// 	s.recvCh <- msg

		case msg := <-s.sendAllCh:
			s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())
		}
	}
}
