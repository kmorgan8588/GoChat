package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kmorgan8588/go-chat-react/pkg/websocket"
)

// serveWs handles requests to the /ws endpoint by upgrading the connection and passing off to the reader function
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Println("Error upgrading connection: ", err)
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
