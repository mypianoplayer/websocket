package game

import (
    "network"
    "time"
)


type TickFunc func()

type Game struct {
    server *network.Server
    tickfunc TickFunc
}

func NewGame() *Game {

    server := network.NewServer("/game")

	return &Game{
		server,
		nil,
	}
}

func (g* Game) SetTickFunc(tickfunc TickFunc) {
    g.tickfunc = tickfunc
}

func (g* Game) Start() {
	go g.Tick()
    go g.server.Listen()
}

func (g* Game) SendAll(msg *network.Message) {
    g.server.SendAll(msg)
}

func (g* Game) Tick() {
	ticker := time.NewTicker(time.Second/60.0)

	for{
		select {
			case <-ticker.C:
			g.tickfunc()
		}
	}

}