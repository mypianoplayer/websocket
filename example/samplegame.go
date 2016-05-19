package samplegame

import (
    "game"
    "math"
    "network"
    "strconv"
)

type SampleGame struct {
    game *game.Game
    cnt float64
}

func Start() {
    samplegame := SampleGame{ game.NewGame(), 1.0 }
	samplegame.game.SetTickFunc( samplegame.tick )
	samplegame.game.Start()
}

func (g* SampleGame) tick() {

	left := 50 + math.Sin(g.cnt) * 30
	top := 70 + math.Sin(g.cnt*0.7+3.4) * 30
	msg := network.Message{"game","player.style.left = " + strconv.Itoa(int(left)) + ";player.style.top=" + strconv.Itoa(int(top))}
	g.cnt += 0.1
	if( g.cnt > 500 ){ g.cnt = 0.0 }
	g.game.SendAll(&msg)

}