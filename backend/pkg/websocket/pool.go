package websocket

import (
	"fmt"
)

type Pool struct {
	// set up a channel to send messages to the pool
	// this channel will be used to send messages to the pool
	// and receive messages from the pool
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

// NewPool function simply returns a new pool
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			//fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				// JSON b/c it's easy to interpret w React frontend
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined!"})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			//fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				// JSON b/c it's easy to interpret w React frontend
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconencted."})
			}
			break
		case message := <-pool.Broadcast:
			//fmt.Println("Sending message to all clients in Pool.")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
