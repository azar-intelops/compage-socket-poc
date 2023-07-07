package main

import (
	"log"
	"net/http"

	"github.com/azar-intelops/websockets/pkg/websockets/server"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	chatServer := server.NewChatServer()
	router.HandleFunc("/ws", chatServer.WebSocketHandler)

	log.Println("Starting chat application...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
