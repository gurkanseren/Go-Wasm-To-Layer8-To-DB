package routes

import (
	"net/http"

	Ctl "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/api/controller"
)

func RegisterRoutes() {
	// Register the routes
	http.HandleFunc("/ping", Ctl.Ping)
}
