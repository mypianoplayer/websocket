package ragtime

import (
	"log"
	"net/http"
	// "time"
	// "strconv"
	// "math"
//	"sync"
	"golang.org/x/net/websocket"
)

// Chat server.
type ConnectionManager struct {
	pattern   string
	connections   map[int]*Connection
	AddConnectionCh     chan *Connection
	DeleteConnectionCh     chan *Connection
	SendAllCh chan *Message
	ErrorCh     chan error
}

// Create new chat server.
func NewConnectionManager(pattern string) *ConnectionManager {
	connections := make(map[int]*Connection)
	addCh := make(chan *Connection)
	delCh := make(chan *Connection)
	sendAllCh := make(chan *Message)
	errCh := make(chan error)

	return &ConnectionManager{
		pattern,
		connections,
		addCh,
		delCh,
		sendAllCh,
		errCh,
	}
}


func (cm *ConnectionManager) sendAll(msg *Message) {
	for _, c := range cm.connections {
		c.SendMessageCh <- msg
	}
}

func (cm * ConnectionManager) Start() {
	go cm.listen()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (cm *ConnectionManager) listen() {

	log.Println("Listening server...")

	onConnected := func(ws *websocket.Conn) {

		defer func() {
			err := ws.Close()
			if err != nil {
				cm.ErrorCh <- err
			}
		}()

		connection := NewConnection(ws, cm)
		cm.AddConnectionCh <- connection
		connection.Listen()
	}
	http.Handle(cm.pattern, websocket.Handler(onConnected))

	// http.HandleFunc(s.pattern,
 //       func(w http.ResponseWriter, req *http.Request) {
 //       	log.Println("HANDLE")
 //           s := websocket.Server{Handler: websocket.Handler(onConnected)}
 //           s.ServeHTTP(w, req)
 //       })

	for {
		select {

		// Add new a client
		case c := <-cm.AddConnectionCh:
			log.Println("Added new client")
			cm.connections[c.id] = c
			log.Println("Now", len(cm.connections), "connections connected.")

		// del a client
		case c := <-cm.DeleteConnectionCh:
			log.Println("Delete client")
			delete(cm.connections, c.id)

		// broadcast message for all connections
		case msg := <-cm.SendAllCh:
//			log.Println("Send all:", msg)
			cm.sendAll(msg)

		case err := <-cm.ErrorCh:
			log.Println("Error:", err.Error())
		}
	}
}


// func (cm* ConnEctionManager) tick() {
// 	ticker := time.NewTicker(time.Second/60.0)
// 	cnt := 1.0
// 	for{
// 		select {
// 			case <-ticker.C:
// 			left := 50 + math.Sin(cnt) * 30
// 			top := 70 + math.Sin(cnt*0.7+3.4) * 30
// 			msg := Message{"server","all","javascript",[]string{"player.style.left = " + strconv.Itoa(int(left)) + ";player.style.top=" + strconv.Itoa(int(top))}}
// 			cnt += 0.1
// 			if( cnt > 500 ){ cnt = 0.0 }
// 			s.SendAllCh <- &msg
// 		}
// 	}

// }