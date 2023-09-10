package api

import (
	"context"
	"log"

	pb "github.com/LiamCWest/ChatTest/api/v1"
	utils "github.com/LiamCWest/ChatTest/api/v1/Utils"
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

func (api API) AddPlayer(name string) *pb.Player {
	player := &pb.Player{Name: name}
	resp, err := api.client.AddPlayer(context.Background(), player)
	if err != nil {
		log.Fatalf("AddPlayer failed: %v", err)
	}

	log.Printf("Added player with ID %s", resp.Id.Id)

	return resp
}

func (api API) GetPlayer(id string) *pb.Player {
	playerID := &pb.PlayerID{Id: id}
	player, err := api.client.GetPlayer(context.Background(), playerID)
	if err != nil {
		log.Fatalf("GetPlayer failed: %v", err)
	}

	return player
}

func (api API) MovePlayer(id string, v utils.Vector2) utils.Vector2 {
	playerID := &pb.PlayerID{Id: id}
	playerMovement := &pb.PlayerMovement{Id: playerID, X: v.X, Y: v.Y}
	player, err := api.client.MovePlayer(context.Background(), playerMovement)
	if err != nil {
		log.Fatalf("MovePlayer failed: %v", err)
	}

	return utils.NewVector2(player.X, player.Y)
}
