package main

import (
	serverApi "github.com/LiamCWest/ChatTest/Client/api"
)

func main() {
	API := serverApi.New()

	// Call the AddPlayer method
	playerID := API.AddPlayer("Liam")

	// Get player position
	API.GetPlayer(playerID)

	// Move player
	API.MovePlayer("1", 1.0, 1.0)
	API.MovePlayer(playerID, 1.0, 1.0)
}
