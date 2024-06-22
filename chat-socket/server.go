package main

import (
	"chatsocket/hub"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	go hub.Center.Run()
	router.HandleFunc("/ws", )

}
