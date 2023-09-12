package main

import (
	"log"

	serverApi "github.com/LiamCWest/ChatTest/Client/api"
	graphics "github.com/LiamCWest/ChatTest/Client/gui"
	utils "github.com/LiamCWest/ChatTest/api/v1/Utils"
)

func main() {
	API := serverApi.New()

	// Call the AddPlayer method
	player := utils.NewPlayerFromMessage(API.AddPlayer("Liam"))

	// Get player position
	log.Printf("Player %s is at X: %f, Y: %f", player.GetID(), player.GetPos().X, player.GetPos().Y)

	// Move player
	API.MovePlayer(player.GetID(), utils.NewVector2(50.0, 50.0))

	// Start the graphics thread
	go graphics.NewGUI(player)

	// Wait for the input thread to finish
	select {}
}
