package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/middleware"
	router "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/router"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the port to listen on from .env
	serverPort := os.Getenv("SERVER_PORT")

	// Set up File Server
	fs := http.FileServer(http.Dir("./assets"))

	// Set up route for File Server
	http.Handle("/", fs)

	// Register the routes using the RegisterRoutes() function with logger middleware
	http.HandleFunc("/api/v1/", middleware.LogRequest(middleware.Cors(router.RegisterRoutes())))

	fmt.Printf("Server listening on localhost:%s\n", serverPort)

	// Start the server on localhost and log any errors
	err = http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
	if err != nil {
		log.Fatal(err)
	}
}
