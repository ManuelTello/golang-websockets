package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

type ConnMap struct {
	size int 
	conns map[int]*websocket.Conn 
}

var cm ConnMap = ConnMap{
	size:0,
	conns:make(map[int]*websocket.Conn),
}

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
		cm.conns[cm.size] = conn
		cm.size = cm.size + 1
	}
	go Reader(conn)
}
