package main

import (
	"log"
	"net/http"
	"ragtime/game"
	"ragtime/server"
	// "os"
)


const (
    ComponentType_Position = iota
)

func main() {
	log.SetFlags(log.Lshortfile)

	// wd,_ := os.Getwd()
	// log.Println(wd)
	// sm := ragtime.NewSceneManager()
	// ts := NewTitleScene()

	// sm.Start()
	// sm.SetSceneCh <- ts

	player := NewPlayer()

	scene := game.NewScene()
	scene.AddObject(player)


	http.Handle("/", http.FileServer(http.Dir("../client/")))
	sv := server.New("/game", scene.MsgCh )

	scene.Start()
	sv.Start()

}
