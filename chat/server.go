package chat

import (
	"log"
	"net/http"
	"time"
	"strconv"
	"math"
	"sync"
	"golang.org/x/net/websocket"
)

type Server struct {
	pattern   string
	messages  []*Message
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *Message
	doneCh    chan bool
	errCh     chan error
	mtx       *sync.Mutex
}

// Create new chat server.
func NewServer(pattern string) *Server {
	messages := []*Message{}
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *Message)
	doneCh := make(chan bool)
	errCh := make(chan error)
	mtx := new(sync.Mutex)

	return &Server{
		pattern,
		messages,
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
		mtx,
	}
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) SendAll(msg *Message) {
	s.sendAllCh <- msg
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

func (s *Server) sendPastMessages(c *Client) {
	for _, msg := range s.messages {
		c.Write(msg)
	}
}

func (s *Server) sendAll(msg *Message) {
	s.mtx.Lock()
	for _, c := range s.clients {
		c.Write(msg)
	}
	s.mtx.Unlock()
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {

	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {

		log.Println("ON CONNECT CALLED")

		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		client := NewClient(ws, s)
		s.Add(client)
		client.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))
	log.Println("Created handler " + s.pattern)

	go s.Tick()

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.mtx.Lock()
			s.clients[c.id] = c
			s.mtx.Unlock()
			log.Println("Now", len(s.clients), "clients connected.")
			s.sendPastMessages(c)

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")
			delete(s.clients, c.id)

		// broadcast message for all clients
		case msg := <-s.sendAllCh:
//			log.Println("Send all:", msg)
			//s.messages = append(s.messages, msg)
			s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}


func (s* Server) Tick() {
	ticker := time.NewTicker(time.Second/60.0)
	cnt := 1.0
	for{
		select {
			case <-ticker.C:
			left := 50 + math.Sin(cnt) * 30
			top := 70 + math.Sin(cnt*0.7+3.4) * 30
			msg := Message{"game","player.style.left = " + strconv.Itoa(int(left)) + ";player.style.top=" + strconv.Itoa(int(top))}
			cnt += 0.1
			if( cnt > 500 ){ cnt = 0.0 }
			s.SendAll(&msg)
		}
	}

}