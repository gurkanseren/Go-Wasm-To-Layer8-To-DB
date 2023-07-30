package controller

import (
	"log"
	"net/http"
	"os"
)

// PingHandler handles ping requests
func Ping(w http.ResponseWriter, r *http.Request) {
	// Send response to client
	_, err := w.Write([]byte("ping successful"))
	if err != nil {
		log.Printf("Error sending response: %v", err)
	}
}

func GetJwtSecret(w http.ResponseWriter, r *http.Request) {
	// Send response to client
	_, err := w.Write([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Printf("Error sending response: %v", err)
	}
}
