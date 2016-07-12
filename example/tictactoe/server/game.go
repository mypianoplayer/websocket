package main

import (
    // "log"
    "ragtime/game"
    "ragtime/server"
    )

type GameScene struct {
    game.Scene
}


func (gs *GameScene) OnMessage(msg *server.Message) {

    // put O or X when clicked

}
