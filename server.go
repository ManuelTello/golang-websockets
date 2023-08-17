package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleServer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	
	if err != nil {
		return
	} else {
		fmt.Print("New connection\n")
		connections.conns[connections.size] = conn
		connections.size = connections.size + 1
	}
	go Reader(conn)
}
