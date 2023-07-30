package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-wasm/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get the port to listen on from .env
	serverPort := os.Getenv("WASM_SERVER_PORT")

	// Set up File Server
	fs := http.FileServer(http.Dir("./assets"))

	// Set up route for File Server
	http.Handle("/", middleware.LogRequest(fs))

	fmt.Printf("Server listening on localhost:%s\n", serverPort)

	// Start the server on localhost and log any errors
	err = http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server stopped")
}
