package main

import (
	"go-crud-learning/myproject/api"
	"log"
	"net/http"
)

func main() {
	// Initialize API handlers
	api.InitHandlers()

	// Start the HTTP server on port 8080
	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
