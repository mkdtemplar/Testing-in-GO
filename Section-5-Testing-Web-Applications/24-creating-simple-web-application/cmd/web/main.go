package main

import (
	"fmt"
	"log"
	"net/http"
)

type application struct {
}

func main() {
	// Set up application configuration
	app := application{}

	// Get application routes
	mux := app.routes()
	// Prion out the message
	fmt.Println("Starting server on port 8080....")

	// Start the server
	err := http.ListenAndServe("localhost:8080", mux)

	if err != nil {
		log.Fatalf("Failed to start server %s", err)
	}
}
