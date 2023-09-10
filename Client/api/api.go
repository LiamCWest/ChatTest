package api

import (
	"context"
	"log"

	pb "github.com/LiamCWest/ChatTest/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type API struct {
	client pb.GameServiceClient
}

func New() API {
	api := API{}

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	api.client = pb.NewGameServiceClient(conn)

	return api
}

func (api API) GetClient() pb.GameServiceClient {
	return api.client
}

func (api API) AddPlayer(name string) (id string) {
	player := &pb.Player{Name: name}
	resp, err := api.client.AddPlayer(context.Background(), player)
	if err != nil {
		log.Fatalf("AddPlayer failed: %v", err)
	}

	log.Printf("Added player with ID %s", resp.Id.Id)

	return resp.Id.Id
}

func (api API) GetPlayer(id string) *pb.Player {
	playerID := &pb.PlayerID{Id: id}
	player, err := api.client.GetPlayer(context.Background(), playerID)
	if err != nil {
		log.Fatalf("GetPlayer failed: %v", err)
	}

	return player
}

func (api API) MovePlayer(id string, x float32, y float32) (xOut float32, yOut float32) {
	playerID := &pb.PlayerID{Id: id}
	playerMovement := &pb.PlayerMovement{Id: playerID, X: x, Y: y}
	player, err := api.client.MovePlayer(context.Background(), playerMovement)
	if err != nil {
		log.Fatalf("MovePlayer failed: %v", err)
	}

	log.Printf("Player position: (%f, %f)", player.X, player.Y)

	return player.X, player.Y
}
