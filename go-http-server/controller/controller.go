package controller

import (
	"encoding/json"
	"net/http"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-http-server/models"
)

func ServeImage(w http.ResponseWriter, r *http.Request) {
	// Create the response object containing the image URL
	imageURLResponse := models.ImageURLResponse{
		URL: models.ImageURL,
	}

	// Set the appropriate Content-Type header for JSON response
	w.Header().Set("Content-Type", "application/json")

	// Encode the response object to JSON and write it to the response writer
	err := json.NewEncoder(w).Encode(imageURLResponse)
	if err != nil {
		http.Error(w, "Error serving image URL", http.StatusInternalServerError)
		return
	}
}
