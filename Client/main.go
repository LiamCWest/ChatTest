package main

import (
	"context"
	"log"

	pb "github.com/LiamCWest/ChatTest/api/v1"
	"google.golang.org/grpc"
)

func main() {
	// Create a connection to the gRPC server
	conn, err := grpc.Dial("localhost:50051")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := pb.NewGameServiceClient(conn)

	// Call the AddPlayer RPC method
	player := &pb.Player{Name: "Alice"}
	resp, err := client.AddPlayer(context.Background(), player)
	if err != nil {
		log.Fatalf("AddPlayer failed: %v", err)
	}

	log.Printf("Added player with ID %d", resp.Id)
}
