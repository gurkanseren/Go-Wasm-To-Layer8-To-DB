package routes

import (
	"net/http"

	Ctl "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/api/controller"
)

func RegisterRoutes() {
	// Register the routes
	http.HandleFunc("/api/v1/ping", Ctl.Ping)
	http.HandleFunc("/api/v1/register-user", Ctl.RegisterUserHandler)
}
