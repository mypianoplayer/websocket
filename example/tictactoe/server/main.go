package main

import (
	"log"
)



const (
    ComponentType_Position = iota
)

func main() {
	log.SetFlags(log.Lshortfile)

	game := NewGame()

	game.Start()

	// wd,_ := os.Getwd()
	// log.Println(wd)
	// sm := ragtime.NewSceneManager()
	// ts := NewTitleScene()

	// sm.Start()
	// sm.SetSceneCh <- ts


}
