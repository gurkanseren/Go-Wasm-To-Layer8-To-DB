package routes

import (
	"net/http"

	Ctl "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8-master/controller"
)

func RegisterRoutes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set up route for API
		switch r.URL.Path {

		case "/api/v1/jwt-secret":
			Ctl.GetJwtSecret(w, r)

		default:
			// Return a 404 Not Found error for unknown routes
			http.NotFound(w, r)
		}
	}
}
