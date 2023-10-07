package main

import (
	"golang-websocket/handler"
	"log"
	"net/http"
)

func main() {
	// Add WebsocketHandler function for the "/websocket" endpoint
	http.HandleFunc("/websocket", handler.WebsocketHandler)

	// Start listening for incoming HTTP requests on port :8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
