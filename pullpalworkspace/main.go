package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Player struct {
	positionX, positionY, velocityX, velocityY float32
	isOnGround, isJumping                      bool
}

func main() {
	fmt.Println("Welcome to Download Simulator 2D Platform Game!")

	// Define core physics variables
	var gravity, momentum, jumpForce float32
	gravity = 9.81
	momentum = 0
	jumpForce = 20

	player := Player{0, 0, 0, 0, true, false}

	// Create a function to simulate a game loop
	gameLoop := time.Tick(time.Millisecond * 16)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Download Simulator 2D Platform Game")

	game := &Game{player: player, gravity: gravity, momentum: momentum, jumpForce: jumpForce}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func (g *Game) playerInput() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.velocityX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.velocityX += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyW) {
		if g.player.isOnGround {
			g.player.velocityY = -g.jumpForce
			g.player.isOnGround = false
			g.player.isJumping = true
		}
	}
}

func updateGame(player *Player, gravity, momentum, jumpForce float32) {
	// TO-DO: Implement game logic to update the state of the game
}

func renderGame(player Player) {
	// Clear the screen
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw the player character at its current position
	drawPlayer(player)

	// Draw the platforms and other game objects
	// TODO: Implement platform and game object rendering

	// Update the screen with the new drawings
	termbox.Flush()
}

func drawPlayer(player Player) {
	x := int(player.positionX)
	y := int(player.positionY)
	termbox.SetCell(x, y, '@', termbox.ColorWhite, termbox.ColorDefault)
}
