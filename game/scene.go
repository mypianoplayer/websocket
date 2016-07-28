package game

import (
	"log"
	"time"
	// "reflect"
	"github.com/mypianoplayer/ragtime/server"
)


type MessageReceiver interface {
	OnMessage(m *server.Message)
}

// ComponentArray
type ComponentArray struct {
	elms [][]Component
}

func (a ComponentArray) Get() [][]Component {
	return a.elms
}

func (a *ComponentArray) Add(idx int, c Component) {

	l := len(a.elms)
	if l < idx+1 {
		a.elms = append(a.elms, make([][]Component, idx+1-l)...)
	}
	a.elms[idx] = append(a.elms[idx], c)
}

func (a *ComponentArray) Clear() {
	a.elms = nil
}

// Scene
type Scene struct {
	objects    []Object
	startComponents ComponentArray
	components ComponentArray

	recvCh chan *server.Message
	receiver MessageReceiver
	// get object by name,id,tag...
}

func NewScene(recvCh chan *server.Message, receiver MessageReceiver) *Scene {
	s := &Scene{
		objects:make([]Object, 0, 50),
		startComponents:ComponentArray{nil},
		components:ComponentArray{nil},
		recvCh:recvCh,
		receiver:receiver,
	}
	return s
}

func (s *Scene) SetReceiver(r MessageReceiver) {
	s.receiver = r
}

func (s *Scene) Update() {
	log.Println("scene update")

	for _, comps := range s.startComponents.Get() {
		for _, c := range comps {
			c.Start()
		}
	}
	
	s.startComponents.Clear()

	for _, comps := range s.components.Get() {
		for _, c := range comps {
			c.Update()
		}
	}

}

func (s *Scene) AddObject(o Object) {
	s.objects = append(s.objects, o)
	
	for c := range EachComponent(o) {
		order := c.UpdateOrder()
		s.startComponents.Add(order, c)
	}
	
	for c := range EachComponent(o) {
		order := c.UpdateOrder()
		s.components.Add(order, c)
	}
}

func (s *Scene) Start() {
    log.Println("Scene Start")
    go s.listen()
}

func (s *Scene) listen() {
	ticker := time.NewTicker(time.Second)///60.0)

	for {
		select {

        case <-ticker.C:
        	s.Update()

        case msg := <- s.recvCh:
        	s.receiver.OnMessage(msg)

        }
	}
}






