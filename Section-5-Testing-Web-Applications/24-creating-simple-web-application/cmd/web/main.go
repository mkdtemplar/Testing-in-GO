package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

type application struct {
	Session *scs.SessionManager
}

func main() {
	// Set up application configuration
	app := application{}

	// get session manager
	app.Session = getSession()

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
