package ragtime

import (
	"log"
	"net/http"
	"time"
	"strconv"
	"math"
//	"sync"
	"golang.org/x/net/websocket"
)

// Chat server.
type GameServer struct {
	pattern   string
	connections   map[int]*Connection
	addCh     chan *Connection
	delCh     chan *Connection
	sendAllCh chan *Message
	doneCh    chan bool
	errCh     chan error
}

// Create new chat server.
func NewGameServer(pattern string) *GameServer {
	connections := make(map[int]*Connection)
	addCh := make(chan *Connection)
	delCh := make(chan *Connection)
	sendAllCh := make(chan *Message)
	doneCh := make(chan bool)
	errCh := make(chan error)

	return &GameServer{
		pattern,
		connections,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
	}
}

func (s *GameServer) Add(c *Connection) {
	s.addCh <- c
}

func (s *GameServer) Del(c *Connection) {
	s.delCh <- c
}

func (s *GameServer) SendAll(msg *Message) {
	s.sendAllCh <- msg
}

func (s *GameServer) Done() {
	s.doneCh <- true
}

func (s *GameServer) Err(err error) {
	s.errCh <- err
}

// func (s *Server) sendPastMessages(c *Client) {
// 	for _, msg := range s.messages {
// 		c.Write(msg)
// 	}
// }

func (s *GameServer) sendAll(msg *Message) {
	for _, c := range s.connections {
		c.Write(msg)
	}
}

// Listen and serve.
// It serves client connection and broadcast request.
func (s *GameServer) Start() {

	log.Println("Listening server...")

	onConnected := func(ws *websocket.Conn) {

		log.Println("CONNECTION>>>")

		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- err
			}
		}()

		connection := NewConnection(ws, s)
		s.Add(connection)
		connection.Listen()
	}
//	http.Handle(s.pattern, websocket.Handler(onConnected))

	http.HandleFunc(s.pattern,
        func(w http.ResponseWriter, req *http.Request) {
        	log.Println("HANDLE")
            s := websocket.Server{Handler: websocket.Handler(onConnected)}
            s.ServeHTTP(w, req)
        })

	log.Println("Created handler " + s.pattern)


	go s.Tick()

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.connections[c.id] = c
			log.Println("Now", len(s.connections), "connections connected.")

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")
			delete(s.connections, c.id)

		// broadcast message for all connections
		case msg := <-s.sendAllCh:
//			log.Println("Send all:", msg)
			s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}


func (s* GameServer) Tick() {
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