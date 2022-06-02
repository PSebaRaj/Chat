package main

import (
	"fmt"
	"golang-chat/pkg/websocket"
	"net/http"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("web socket endpoint reached")
	conn, err := websocket.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	// after creating the client, we add it to the pool register channel
	pool.Register <- client

	// after the client has been read, we add it to the pool unregister channel
	client.Read()

}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	// web socket endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}

func main() {
	fmt.Println("yer")
	setupRoutes()
	http.ListenAndServe(":9000", nil)

}
