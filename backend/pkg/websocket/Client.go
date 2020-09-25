package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

//Client struct for holding a websocket connection and a pointer to the Pool that it is stored in
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

//Message structs stores the message type and text from a successful connection
type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

//Read takes a client and attempts to read a message and passes it into the pool for other clients to read
func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println("Error reading message", err)
			return
		}
		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message received %+v\n", message)
	}
}
