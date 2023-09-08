package controller

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
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

func InitializeECDHKeyExchange(w http.ResponseWriter, r *http.Request) {

	// Unmarshal the request body into the ECDHKeyExchangeRequest object
	var ecdhKeyExchangeRequest models.ECDHKeyExchangeRequest
	err := json.NewDecoder(r.Body).Decode(&ecdhKeyExchangeRequest)
	if err != nil {
		http.Error(w, "Error decoding request body", http.StatusInternalServerError)
		return
	}

	privKeyContentServer, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pubKeyContentServerX, pubKeyContentServerY := privKeyContentServer.PublicKey.Curve.ScalarBaseMult(privKeyContentServer.D.Bytes())
	pubKeyContentServer := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     pubKeyContentServerX,
		Y:     pubKeyContentServerY,
	}
	sharedX, sharedY := elliptic.P256().ScalarMult(ecdhKeyExchangeRequest.PubKeyWasmX, ecdhKeyExchangeRequest.PubKeyWasmY, privKeyContentServer.D.Bytes())
	sharedKeyContentServer := &ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     sharedX,
		Y:     sharedY,
	}

	fmt.Printf("\nShared key (Content Server) (%x, %x)\n", sharedKeyContentServer.X, sharedKeyContentServer.Y)

	// Send back the response object containing the server's private key
	ecdhKeyExchangeOutput := models.ECDHKeyExchangeOutput{
		PubKeyServerX: pubKeyContentServer.X,
		PubKeyServerY: pubKeyContentServer.Y,
	}

	// Set the appropriate Content-Type header for JSON response
	w.Header().Set("Content-Type", "application/json")

	// Encode the response object to JSON and write it to the response writer
	err = json.NewEncoder(w).Encode(ecdhKeyExchangeOutput)
	if err != nil {
		http.Error(w, "Error sending server's public key", http.StatusInternalServerError)
		return
	}
}
