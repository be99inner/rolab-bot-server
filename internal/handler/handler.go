package handler

import (
	"log"
	"net/http"

	// "github.com/be99inner/rolab-bot-server/internal/processing"
	"github.com/be99inner/rolab-bot-utility/networking"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ServeWs serves WebSocket endpoint to client
func ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade error: %v\n", err)
	}
	defer conn.Close()

	Handle(conn)
}

// Handle handles the GameData from Websocket
func Handle(conn *websocket.Conn) {
	for {
		data, err := networking.ReceiveData(conn)
		if err != nil {
			log.Printf("Receive err: %v\n", err)
		}

		log.Printf("Received data: %+v\n", data)

		// // Decode the base64 image from the payload
		// img, err := processing.DecodeImage(data.Payload)
		// if err != nil {
		// 	log.Printf("Error decoding image: %v\n", err)
		// }
		//
		// // Save the received image for preview purposes
		// err = processing.SaveImage("received_img.jpg", img)
		// if err != nil {
		// 	log.Printf("Error saving image: %v\n", err)
		// 	break
		// }

		response := networking.GameData{
			EventType: "image_received",
			Payload:   "Image received and saved successfully",
		}

		err = networking.SendData(conn, response)
		if err != nil {
			log.Printf("Send error: %v\n", err)
			break
		}
	}
}
