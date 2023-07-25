package controller

import (
	"encoding/json"
	"net/http"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-http-server/models"
)

func ServeImage(w http.ResponseWriter, r *http.Request) {

	// Get the image ID from the request parameters
	imageID := r.URL.Query().Get("id")

	// Create the response object containing the image URL according to the image ID
	var imageURLResponse models.ImageURLResponse
	switch imageID {
	case "1":
		imageURLResponse = models.ImageURLResponse{
			URL: models.ImageURLOne,
		}
	case "2":
		imageURLResponse = models.ImageURLResponse{
			URL: models.ImageURLTwo,
		}
	case "3":
		imageURLResponse = models.ImageURLResponse{
			URL: models.ImageURLThree,
		}
	case "4":
		imageURLResponse = models.ImageURLResponse{
			URL: models.ImageURLFour,
		}
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
