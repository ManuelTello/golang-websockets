package main

type Client struct {
	username string
	posx int 
	posy int 
	color int 
}

type ClientsMap struct {
	size int
	clients map[int]Client 
}

var clients ClientsMap = ClientsMap {
	size:0,
	clients:make(map[int]Client),
}
