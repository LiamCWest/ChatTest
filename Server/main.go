package Server

import (
	"context"
	"log"
	"net"

	pb "github.com/LiamCWest/ChatTest/api/v1/go"
	"google.golang.org/grpc"
)

type gameServer struct {
	pb.GameServiceServer
}

func (s *gameServer) AddPlayer(ctx context.Context, req *pb.Player) (*pb.Player, error) {
	return &pb.Player{}, nil
}

func (s *gameServer) GetPlayer(ctx context.Context, req *pb.PlayerID) (*pb.Player, error) {
	return &pb.Player{}, nil
}

func (s *gameServer) MovePlayer(ctx context.Context, req *pb.PlayerMovement) (*pb.Player, error) {
	return &pb.Player{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, &gameServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
