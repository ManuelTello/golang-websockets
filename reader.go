package main

import (
	"github.com/gorilla/websocket"
	"fmt"
	"encoding/json"
)

func Reader(connection *websocket.Conn){
	for{
		msgtype,msg,msgerr := connection.ReadMessage()

		if msgerr != nil{
			fmt.Printf("Error on message: %s\n",msgerr)
		}

		var usrm Client
		err:=json.Unmarshal(msg,&usrm)
	
		if err != nil{
			fmt.Printf("Error parsing json: %s\n",err)
		}

		var user_index = IdentifyClientByUsername(usrm.username)

		if user_index != -1 {
			var user_copy Client = clients.clients[user_index]
			user_copy.username = usrm.username
			user_copy.color = usrm.color
			user_copy.px = usrm.px
			user_copy.py = usrm.py
			clients.clients[user_index] = user_copy
		}else{			
			clients.clients[clients.size] = Client{
				username:usrm.username,
				color:usrm.color,
				px:usrm.px,
				py:usrm.py,
			}
			clients.size = clients.size + 1
		}				

		for i:=0;i<connections.size;i++ {
			data,_ := json.Marshal(clients.clients)
			connections.conns[i].WriteMessage(msgtype,data)
		}
	}
}
