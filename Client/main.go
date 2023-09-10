package main

import (
	pb "github.com/LiamCWest/ChatTest/api/v1"

	api "github.com/LiamCWest/ChatTest/Client/api"
)

func main() {
	API := api.New()

	// Create a client instance
	client := pb.NewGameServiceClient(API.conn)

	// Call the AddPlayer method
	playerID := API.AddPlayer(client, "Liam")

	// Get player position
	API.GetPlayer(client, playerID)

	// Move player
	API.MovePlayer(client, "1", 1.0, 1.0)
	API.MovePlayer(client, playerID, 1.0, 1.0)
}
