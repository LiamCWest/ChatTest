package main

import (
	"context"
	"log"
	"net"

	"strconv"

	pb "github.com/LiamCWest/ChatTest/api/v1"
	utils "github.com/LiamCWest/ChatTest/api/v1/Utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gameServer struct {
	pb.GameServiceServer
	players map[string]*utils.Player // map of player names to player objects
	nextID  int32
}

var gameData struct {
	players map[string]*utils.Player // map of player names to player objects
}

func (s *gameServer) AddPlayer(ctx context.Context, req *pb.Player) (*pb.Player, error) {
	if s.players == nil {
		s.players = make(map[string]*utils.Player)
	}
	s.nextID++ // increment the counter variable
	req.Id = &pb.PlayerID{Id: strconv.Itoa(int(s.nextID))}
	log.Printf("Adding player %s with ID %s", req.Name, req.Id.Id)
	s.players[req.Id.Id] = utils.NewPlayerFromMessage(req)
	return req, nil
}

func (s *gameServer) GetPlayer(ctx context.Context, req *pb.PlayerID) (*pb.Player, error) {
	player := s.players[req.Id]
	if player == nil {
		return nil, status.Errorf(codes.NotFound, "Player with ID %s not found", req.Id) // use req.Id.Id instead of req.Id
	}
	return player.ToMessage(), nil
}

func (s *gameServer) MovePlayer(ctx context.Context, req *pb.PlayerMovement) (*pb.Player, error) {
	player := s.players[req.Id.Id]
	player.Move(utils.NewVector2(req.X, req.Y))

	return player.ToMessage(), nil
}

func (s *gameServer) GetPlayers(req *pb.Empty, stream pb.GameService_GetPlayersServer) error {
	players := make([]*pb.Player, 0, len(s.players))
	for _, player := range s.players {
		players = append(players, player.ToMessage())
	}

	// Create a Players message that contains the list of players
	playersMessage := &pb.Players{Players: players}

	// Send the Players message to the client
	if err := stream.Send(playersMessage); err != nil {
		return err
	}

	return nil
}

func (s *gameServer) RemovePlayer(ctx context.Context, req *pb.PlayerID) (*pb.Empty, error) {
	delete(s.players, req.Id)
	log.Printf("Removed player with ID %s", req.Id)
	return &pb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, &gameServer{})

	log.Printf("Starting server on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
