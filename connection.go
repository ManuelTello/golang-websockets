package main 

import "github.com/gorilla/websocket"

type ConnMap struct {
	size int 
	conns map[int]*websocket.Conn 
}

var connections ConnMap = ConnMap{
	size:0,
	conns:make(map[int]*websocket.Conn),
}

