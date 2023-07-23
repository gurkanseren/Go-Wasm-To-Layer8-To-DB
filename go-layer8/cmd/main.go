package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/middleware"
	router "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/router"
	"github.com/joho/godotenv"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
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

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		c, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Printf("failed to accept websocket connection: %v", err)
			return
		}
		defer c.Close(websocket.StatusInternalError, "the sky is falling")

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		var v interface{}
		err = wsjson.Read(ctx, c, &v)
		if err != nil {
			log.Printf("failed to read message: %v", err)
			return
		}

		log.Printf("received: %v", v)

		c.Close(websocket.StatusNormalClosure, "")
	})

	// Register the routes using the RegisterRoutes() function with logger middleware
	http.HandleFunc("/api/v1/", middleware.LogRequest(middleware.Cors(router.RegisterRoutes())))

	fmt.Printf("Server listening on localhost:%s\n", serverPort)

	// Start the server on localhost and log any errors
	err = http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server stopped")
}
