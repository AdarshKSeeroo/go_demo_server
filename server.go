package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Serve static HTML files from the "static" directory
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	// Start the server on port 8080
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}