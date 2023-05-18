package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	positionX, positionY, velocityX, velocityY float64
	isOnGround, isJumping                      bool
}

func main() {
	fmt.Println("Welcome to Download Simulator 2D Platform Game!")

	// Define core physics variables
	var gravity, momentum, jumpForce float64
	gravity = 9.81
	momentum = 0
	jumpForce = 20

	player := Player{positionX: 320, positionY: 240, velocityX: 0, velocityY: 0, isOnGround: true, isJumping: false}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Download Simulator 2D Platform Game")

	game := &Game{player: player, gravity: gravity, momentum: momentum, jumpForce: jumpForce}

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

func (g *Game) playerInput() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.player.velocityX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.player.velocityX += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if g.player.isOnGround {
			g.player.velocityY = -g.jumpForce
			g.player.isOnGround = false
			g.player.isJumping = true
		}
	}
}

type Game struct {
	player    Player
	gravity   float64
	momentum  float64
	jumpForce float64
}

func (g *Game) Update() error {
	// Update game logic here
	g.playerInput()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw game objects here
	playerRect := ebiten.NewImage(20, 20)
	playerRect.Fill(color.RGBA{255, 0, 0, 255})
	screen.Fill(color.RGBA{245, 245, 245, 255})
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.player.positionX-10, g.player.positionY-10)
	screen.DrawImage(playerRect, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}
