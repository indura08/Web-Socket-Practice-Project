package handlers

import (
	"fmt"
	"github/indura08/web-socket-practice-project/managers"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HanldeWS(hub *managers.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		return
	}

	hub.Register <- conn

	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			hub.Unregister <- conn
			break
		}

		hub.Broadcast <- msg
	}

	fmt.Print("Successfully disonnected to web socket")
}
