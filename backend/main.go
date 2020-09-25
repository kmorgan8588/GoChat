package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrader used to set Read/Write bufffer sizes and a function to handle checking where requests come from
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

// reader takes a websocket connection and runs indefinitely until an error occurs.  It handles reading and writing messsages
func reader(conn *websocket.Conn) {
	for {

		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error reading message: ", err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("Error writing message", err)
			return
		}
	}
}

// serveWs handles requests to the /ws endpoint by upgrading the connection and passing off to the reader function
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error upgrading connection: ", err)
	}

	reader(ws)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home")
}

func setupRoutes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01 listening on port :8080")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
