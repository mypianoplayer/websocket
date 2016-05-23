package main

import (
	"log"
	"net/http"
	"ragtime"
	// "os"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// wd,_ := os.Getwd()
	// log.Println(wd)

	http.Handle("/", http.FileServer(http.Dir("../client/")))
	server := ragtime.NewGameServer("/game")
	server.Start()

}
