package websocket

import (
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
