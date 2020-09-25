package websocket

import (
	"fmt"
	"io"
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

//Upgrade takes a connection request and upgrades the connection to a websocket
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return ws, err
	}
	return ws, nil
}

//Reader takes a websocket connection and runs indefinitely until an error occurs.  It handles reading and writing messsages
func Reader(conn *websocket.Conn) {
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

//Writer takes a websocket connection and maintains
func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")

		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println("Error getting next reader: ", err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println("Error getting next writer: ", err)
			return
		}

		if _, err := io.Copy(w, r); err != nil {
			fmt.Println("Error copying reader/writer: ", err)
			return
		}

		if err := w.Close(); err != nil {
			fmt.Println("Error closing writer: ", err)
			return
		}
	}
}
