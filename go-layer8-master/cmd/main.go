package main

import (
	"fmt"
	"log"
	"net"
	"os"

	interfaces "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8-master/pkg/interface"
	pb "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8-master/pkg/service"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Get the port to listen on from .env
	serverPort := os.Getenv("LAYER8_MASTER_PORT")

	lis, err := net.Listen("tcp", "localhost:"+serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLayer8MasterServiceServer(s, &interfaces.Server{})

	fmt.Println("Layer8 Master gRPC server listening on port " + serverPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
