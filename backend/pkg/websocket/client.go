package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		msgType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: msgType, Body: string(p)}
		c.Pool.Broadcast <- message // write the message on the channel (broadcast))
		fmt.Printf("Message Received: %+v\n", message)
	}
}
