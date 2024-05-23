package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/be99inner/rolab-bot-server/internal/handler"
)

var addr = flag.String("addr", "localhost:3000", "HTTP service address")

func main() {
	log.Printf("Server is running on %s\n", *addr)

	http.HandleFunc("/ws", handler.ServeWs)
	http.HandleFunc("/preview", handler.PreviewImage)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}
