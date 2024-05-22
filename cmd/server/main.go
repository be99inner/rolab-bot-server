package main

import (
	"log"
	"net/http"

	"github.com/be99inner/rolab-bot-server/internal/handler"
	// "github.com/gorilla/websocket"
)

var addr = "localhost:3000"

func main() {
	log.Printf("Server is running on %s\n", addr)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWs(w, r)
	})

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
