package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebsocketHandler is an HTTP handler function that manages WebSocket connections.
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		// Log the received message
		log.Printf("Received message: %s", message)

		// Send the message back to the client
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}
	}
}

// upgrader is a configured Upgrader object to upgrade WebSocket connections.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin validates the origin of the requesting client
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
