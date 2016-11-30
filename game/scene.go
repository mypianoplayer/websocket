package game

import (
	"log"
	"time"
	"sort"
	// "reflect"
	"github.com/mypianoplayer/ragtime/server"
)


type MessageReceiver interface {
	OnMessage(m *server.Message)
}

type ComponentArray []Component

func (ca ComponentArray) updateOrder() int {
	if len(ca) <= 0 {
		log.Println("ComponentArray size 0")
	}
	return ca[0].UpdateOrder()
}

type Components []ComponentArray

func (comps Components) Len() int {
	return len(comps)
}

func (comps Components) Swap(i, j int) {
	comps[i],comps[j] = comps[j],comps[i]
}

func (comps Components) Less(i, j int) bool {
	return comps[i].updateOrder() < comps[j].updateOrder()
}

func (comps Components) Add(c Component) Components {

	var exist bool
	for i := 0; i < len(comps); i++ {
		if comps[i].updateOrder() == c.UpdateOrder() {
			comps[i] = append( comps[i], c )
			exist = true
		}
	}
	
	if !exist {
		ca := make(ComponentArray, 0, 10)
		ca = append(ca, c)
		comps = append( comps, ca )
		sort.Sort(comps)
	}
	
	return comps
}

// Scene
type Scene struct {
	objects    []Object
	startComponents Components
	components Components

	recvCh chan *server.Message
	receiver MessageReceiver
}

func NewScene(recvCh chan *server.Message, receiver MessageReceiver) *Scene {
	s := &Scene{
		objects:make([]Object, 0, 50),
		startComponents:make(Components, 0, 10),
		components:make(Components, 0, 10),
		recvCh:recvCh,
		receiver:receiver,
	}
	return s
}

func (s *Scene) SetReceiver(r MessageReceiver) {
	s.receiver = r
}

func (s *Scene) Update() {
//	log.Println("scene update")

	for _,comps := range s.startComponents {
		for _,c := range comps {
			c.Start()
		}
	}
	
	s.startComponents = nil

	for _,comps := range s.components {
		for _,c := range comps {
			c.Update()
		}
	}

}

func (s *Scene) AddObject(o Object) {
	s.objects = append(s.objects, o)
	
	for _,c := range o.Components() {
		s.startComponents = s.startComponents.Add(c)
	}
	
	for _,c := range o.Components() {
		s.components = s.components.Add(c)
	}
}

func (s *Scene) Start() {
    log.Println("Scene Start")
    go s.listen()
}

func (s *Scene) listen() {
	ticker := time.NewTicker(time.Second/60.0)

	for {
		select {

        case <-ticker.C:
        	s.Update()

        case msg := <- s.recvCh:
        	s.receiver.OnMessage(msg)

        }
	}
}






