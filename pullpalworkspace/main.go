package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Download Simulator 2D Platform Game!")

	// Define core physics variables
	var gravity, momentum, jumpForce float32
	gravity = 9.81
	momentum = 0
	jumpForce = 20

	// Create player character
	type Player struct {
		positionX, positionY, velocityX, velocityY float32
		isOnGround, isJumping                      bool
	}
	player := Player{0, 0, 0, 0, true, false}

	// TO-DO: Implement main game loop

	// TO-DO: Implement player character controller
}
