package main

import (
    // "log"
    "ragtime/game"
    "ragtime/server"
    "net/http"
    )

type Game struct {
    scene *game.Scene
    server *server.Server
}

func NewGame() *Game {
    sc := game.NewScene()
    sv := server.New("/game", sc.MsgCh)
    return &Game{sc,sv}
}

func (g *Game) Scene() *game.Scene {
    return g.scene
}

func (g *Game) Server() *server.Server {
    return g.server
}

func (g* Game) Start() {


	player := NewPlayer()

	g.scene.AddObject(player)

	http.Handle("/", http.FileServer(http.Dir("../client/")))

	g.scene.Start()
	g.server.Start()

}
