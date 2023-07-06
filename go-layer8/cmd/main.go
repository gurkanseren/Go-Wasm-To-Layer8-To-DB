package main

import (
	"log"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/config"
	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/internal/rest_server"
)

func main() {
	// Load configuration
	conf := config.LoadConfig()

	// Create new server instance
	server := rest_server.NewServer(conf)

	// Start the server
	log.Printf("Starting server on port %d...", conf.RESTPort)
	err := server.Serve()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
