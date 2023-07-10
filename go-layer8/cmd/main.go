package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/middleware"
	router "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/router"
)

func main() {
	serverPort := 8080

	// Create a new server mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", router.HandleRequest)

	// Wrap the server mux with the CORS middleware
	handler := middleware.Cors(mux)

	// Start the server
	fmt.Printf("Starting server on port %v\n", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), handler)
	if err != nil {
		log.Fatal(err)
	}
}
