package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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

type Game struct {
	player    Player
	gravity   float32
	momentum  float32
	jumpForce float32
}
func (g *Game) Update() error {
	// Update game logic here
	g.playerInput()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw game objects here
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
