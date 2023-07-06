package main

import (
	"log"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/config"
	http "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/internal/httpServer"
)

func main() {
	// Load configuration
	conf := config.LoadConfig()
	// Create new server instance
	server := http.NewServer(conf)
	// Start the server
	log.Printf("Starting server on port %d...", conf.RESTPort)
	err := server.Serve()
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
