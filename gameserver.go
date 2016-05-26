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
	AddConnectionCh     chan *Connection
	DeleteConnectionCh     chan *Connection
	SendAllCh chan *Message
	ErrorCh     chan error
}

// Create new chat server.
func NewGameServer(pattern string) *GameServer {
	connections := make(map[int]*Connection)
	addCh := make(chan *Connection)
	delCh := make(chan *Connection)
	sendAllCh := make(chan *Message)
	errCh := make(chan error)

	return &GameServer{
		pattern,
		connections,
		addCh,
		delCh,
		sendAllCh,
		errCh,
	}
}


func (s *GameServer) sendAll(msg *Message) {
	for _, c := range s.connections {
		c.SendMessageCh <- msg
	}
}

func (s * GameServer) Start() {
	go s.listen()
	go s.tick()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (s *GameServer) listen() {

	log.Println("Listening server...")

	onConnected := func(ws *websocket.Conn) {

		defer func() {
			err := ws.Close()
			if err != nil {
				s.ErrorCh <- err
			}
		}()

		connection := NewConnection(ws, s)
		s.AddConnectionCh <- connection
		connection.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))

	// http.HandleFunc(s.pattern,
 //       func(w http.ResponseWriter, req *http.Request) {
 //       	log.Println("HANDLE")
 //           s := websocket.Server{Handler: websocket.Handler(onConnected)}
 //           s.ServeHTTP(w, req)
 //       })

	for {
		select {

		// Add new a client
		case c := <-s.AddConnectionCh:
			log.Println("Added new client")
			s.connections[c.id] = c
			log.Println("Now", len(s.connections), "connections connected.")

		// del a client
		case c := <-s.DeleteConnectionCh:
			log.Println("Delete client")
			delete(s.connections, c.id)

		// broadcast message for all connections
		case msg := <-s.SendAllCh:
//			log.Println("Send all:", msg)
			s.sendAll(msg)

		case err := <-s.ErrorCh:
			log.Println("Error:", err.Error())
		}
	}
}


func (s* GameServer) tick() {
	ticker := time.NewTicker(time.Second/60.0)
	cnt := 1.0
	for{
		select {
			case <-ticker.C:
			left := 50 + math.Sin(cnt) * 30
			top := 70 + math.Sin(cnt*0.7+3.4) * 30
			msg := Message{"server","all","javascript",[]string{"player.style.left = " + strconv.Itoa(int(left)) + ";player.style.top=" + strconv.Itoa(int(top))}}
			cnt += 0.1
			if( cnt > 500 ){ cnt = 0.0 }
			s.SendAllCh <- &msg
		}
	}

}