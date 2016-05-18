package main

import (
	"log"
	"net/http"
	"samplegame"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// websocket server
//	server := chat.NewServer("/entry")
	// server := network.NewServer("/game")
	// go server.Listen()
	samplegame.Start()

	// static files
	http.Handle("/", http.FileServer(http.Dir("webroot")))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
