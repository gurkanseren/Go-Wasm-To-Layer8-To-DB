package rest_server

import (
	"fmt"
	"net/http"

	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/config"
	serverInterface "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/interface"
	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8/middleware"
)

type HTTPServer struct {
	conf   *config.Config
	server *http.Server
}

func NewServer(conf *config.Config) serverInterface.ServerImpl {
	return &HTTPServer{
		conf:   conf,
		server: &http.Server{Addr: ":" + fmt.Sprint(conf.ServerPORT)},
	}
}

// Serve registers the services and starts serving
func (s *HTTPServer) Serve() error {
	httpServer := http.NewServeMux()

	// register handlers
	httpServer.HandleFunc("/api/v1/ping", s.pingHandler)
	httpServer.HandleFunc("/api/v1/register-user", s.registerUserHandler)
	httpServer.HandleFunc("/api/v1/login-precheck", s.loginPrecheckHandler)
	httpServer.HandleFunc("/api/v1/login-user", s.loginUserHandler)

	// Apply middleware
	handler := middleware.Cors(httpServer)
	s.server.Handler = handler
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server
func (s *HTTPServer) Shutdown() {
	s.server.Close()
}
