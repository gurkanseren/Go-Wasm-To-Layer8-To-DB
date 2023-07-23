package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func serveImage(w http.ResponseWriter, r *http.Request) {
	// Open the image file
	imageFile, err := os.Open("static/image.png")
	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	defer imageFile.Close()

	// Set the appropriate Content-Type header
	w.Header().Set("Content-Type", "image/png")

	// Copy the image data to the response writer
	_, err = io.Copy(w, imageFile)
	if err != nil {
		http.Error(w, "Error serving image", http.StatusInternalServerError)
		return
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go Image Server")
	})

	// Define the endpoint for serving the image
	http.HandleFunc("/image", serveImage)

	// Start the HTTP server on port 9090
	port := 9090
	fmt.Printf("Server started on http://localhost:%d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
