package main

import (
	"fmt"
	"net/http"

	interfaces "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-load-balancer/interface"
	"github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-load-balancer/utils"
)

func main() {
	servers := []interfaces.Server{
		utils.NewSimpleServer("http://localhost:8001"),
		utils.NewSimpleServer("http://localhost:8002"),
		utils.NewSimpleServer("http://localhost:8003"),
	}

	lb := utils.NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	}

	// register a proxy handler to handle all requests
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.Port)
	http.ListenAndServe(":"+lb.Port, nil)
}
