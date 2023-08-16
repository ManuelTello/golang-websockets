package main

import (
	"github.com/gorilla/websocket"
	"fmt"
)

func Reader(connection *websocket.Conn){
	for{
		_,msg,_ := connection.ReadMessage()
		
		for i:=0;i<cm.size;i++ {
			
		}
		
		for i:=0;i<cm.size;i++ {
			//cm.conns[i].WriteMessage(msgt,msg)
			fmt.Print(string(msg))		
		}
	}
}
