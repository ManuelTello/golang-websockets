package main

type Client struct {
	username string  `json:"username"`
	px int32 `json:"px"`
	py int32 `json:"py"`
	color int32 `json:"color"`
}

type ClientsMap struct {
	size int
	clients map[int]Client 
}

var clients ClientsMap = ClientsMap {
	size:0,
	clients:make(map[int]Client),
}

func IdentifyClientByUsername(username string) int {
	var result int = -1

	for i:=0;i<clients.size;i++{
		if clients.clients[i].username == username{
			result = i
		}
	}
	return result
}
