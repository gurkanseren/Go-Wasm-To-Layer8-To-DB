package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/router"
)

func main() {
	serverPort := 8080
	// Register the routes
	router.RegisterRoutes()
	// Start the server
	fmt.Printf("Starting server on port %v\n", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}
