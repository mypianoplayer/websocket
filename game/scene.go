package game

import (
	"log"
	"time"
	// "reflect"
	"ragtime/server"
)

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

type Scene struct {
	holders    []ComponentHolder
	components ComponentArray

	MsgCh chan *server.Message
	// get object by name,id,tag...
}

func NewScene() *Scene {
	ch := make([]ComponentHolder, 0, 50)
	ca := ComponentArray{nil}
	return &Scene{ch, ca, make(chan *server.Message)}
}

func (s *Scene) Update() {
	log.Println("scene update")
	for _, comps := range s.components.Get() {
		for _, c := range comps {
			c.Update()
		}
	}

}

func (s *Scene) OnMessage(msg *server.Message) {

}

func (s *Scene) AddObject(holder ComponentHolder) {
	s.holders = append(s.holders, holder)
	for c := range EachComponent(holder) {
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

        case msg := <- s.MsgCh:
        	s.OnMessage(msg)

        }
	}
}






