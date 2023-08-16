package interfaces

import (
	"context"
	"os"

	pb "github.com/globe-and-citizen/Go-Wasm-To-Layer8-To-DB/go-layer8-master/pkg/service"
)

type Server struct {
	pb.UnimplementedLayer8MasterServiceServer
}

func (s *Server) GetJwtSecret(ctx context.Context, empty *pb.Empty) (*pb.JwtSecretResponse, error) {
	return &pb.JwtSecretResponse{JwtSecret: os.Getenv("JWT_SECRET")}, nil
}

func (s *Server) GetPublicKey(ctx context.Context, empty *pb.Empty) (*pb.PublicKeyResponse, error) {
	return &pb.PublicKeyResponse{PublicKey: os.Getenv("PUBLIC_KEY")}, nil
}
