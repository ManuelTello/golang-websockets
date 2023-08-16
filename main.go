package main 

import (
	"net/http"
)

func main(){
	http.HandleFunc("/echo",HandleServer)
	http.ListenAndServe(":8080",nil)
}
