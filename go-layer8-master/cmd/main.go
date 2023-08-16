package main

import (
	"log"
	"net"
	"os"

	interfaces "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8-master/pkg/interface"
	pb "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8-master/pkg/service"
	utils "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8-master/utils"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the port to listen on from .env
	serverPort := os.Getenv("LAYER8_MASTER_PORT")

	lis, err := net.Listen("tcp", "localhost:"+serverPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new logger instance for more detailed logging
	logger := log.New(os.Stdout, "[Layer8-M-gRPC] ", log.LstdFlags)

	s := grpc.NewServer(
		// You can add interceptor here to log incoming requests
		grpc.UnaryInterceptor(utils.UnaryInterceptor(logger)),
	)
	pb.RegisterLayer8MasterServiceServer(s, &interfaces.Server{})

	logger.Printf("Layer8 Master gRPC server listening on port %s", serverPort)
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve: %v", err)
	}
}
