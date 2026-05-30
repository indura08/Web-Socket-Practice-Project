package main

import (
	"github/indura08/web-socket-practice-project/handlers"
	"github/indura08/web-socket-practice-project/managers"
	"log"
	"net/http"
)

func main() {
	hub := managers.NewHub()

	go hub.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handlers.HanldeWS(hub, w, r)
	})

	http.Handle("/", http.FileServer(
		http.Dir("./static"),
	))

	log.Println("Server started")
	http.ListenAndServe(":8080", nil)
}
