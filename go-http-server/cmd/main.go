package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	Ctl "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-http-server/controller"
)

func main() {
	// Load environment variables
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("SERVER_TWO_PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Image Server")
	})

	// Define the endpoint for serving the image URL
	http.HandleFunc("/image", Ctl.ServeImage)

	// Start the HTTP server on port 9090
	fmt.Printf("Server listening on localhost:%s\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
