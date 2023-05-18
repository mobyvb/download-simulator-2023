package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

type Player struct {
	positionX, positionY, velocityX, velocityY float32
	isOnGround, isJumping                      bool
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	fmt.Println("Welcome to Download Simulator 2D Platform Game!")

	// Define core physics variables
	var gravity, momentum, jumpForce float32
	gravity = 9.81
	momentum = 0
	jumpForce = 20

	player := Player{0, 0, 0, 0, true, false}

	// Set up player input
	go func() {
		for {
			event := termbox.PollEvent()
			if event.Type == termbox.EventKey {
				playerInput(event.Key, &player, jumpForce)
			}
		}
	}()

	// Create a function to simulate a game loop
	gameLoop := time.Tick(time.Millisecond * 16)

	for range gameLoop {
		updateGame(&player, gravity, momentum, jumpForce)
		renderGame(player)
	}
}

func playerInput(key termbox.Key, player *Player, jumpForce float32) {
	switch key {
	case termbox.KeyArrowLeft, termbox.KeyCtrlB:
		player.velocityX -= 5
	case termbox.KeyArrowRight, termbox.KeyCtrlF:
		player.velocityX += 5
	case termbox.KeyArrowUp, termbox.KeySpace, termbox.KeyCtrlP:
		if player.isOnGround {
			player.velocityY = -jumpForce
			player.isOnGround = false
			player.isJumping = true
		}
	}
}

func updateGame(player *Player, gravity, momentum, jumpForce float32) {
	// TO-DO: Implement game logic to update the state of the game
}

func renderGame(player Player) {
	// TODO: Clear the screen
	// TODO: Draw the player character at its current position
	// TODO: Draw the platforms and other game objects
	// TODO: Update the screen with the new drawings
}
