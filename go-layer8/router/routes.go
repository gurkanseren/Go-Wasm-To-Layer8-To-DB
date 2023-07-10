package routes

import (
	"net/http"

	Ctl "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/api/controller"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/api/v1/ping":
		// Handle the "/api/v1/ping" route
		Ctl.Ping(w, r)
	case "/api/v1/register-user":
		// Handle the "/api/v1/register-user" route
		Ctl.RegisterUserHandler(w, r)
	case "/api/v1/login-precheck":
		// Handle the "/api/v1/login-precheck" route
		Ctl.LoginPrecheckHandler(w, r)
	case "/api/v1/login-user":
		// Handle the "/api/v1/login-user" route
		Ctl.LoginUserHandler(w, r)
	default:
		// Return a 404 Not Found error for unknown routes
		http.NotFound(w, r)
	}
}
