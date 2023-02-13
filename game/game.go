package game

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

type Vector struct {
	X, Y float64
}

func NewVector(x, y float64) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vector) Scale(s float64) Vector {
	return Vector{
		X: v.X * s,
		Y: v.Y * s,
	}
}

func (v Vector) Point() image.Point {
	return image.Point{
		X: int(v.X),
		Y: int(v.Y),
	}
}

func (v Vector) String() string {
	return fmt.Sprintf("<%f, %f>", v.X, v.Y)
}

var (
	bounds = Vector{X: float64(1200), Y: float64(300)}
)

type Entity struct {
	pos, vel, acc, size Vector
}

func (e Entity) Update(dt time.Duration) Entity {
	e.vel = e.vel.Add(e.acc.Scale(dt.Seconds()))
	e.pos = e.pos.Add(e.vel.Scale(dt.Seconds()))
	return e
}

type GameState struct {
	Player Entity
}

func NewGame() GameState {
	return GameState{
		Player: Entity{
			pos:  NewVector(bounds.X/2, bounds.Y/2),
			vel:  NewVector(100, 0),
			size: NewVector(20, 20),
		},
	}
}

func (gs GameState) Update(dt time.Duration) GameState {
	gs.Player = gs.Player.Update(dt)
	return gs
}

func (gs GameState) Draw(gtx layout.Context, pixelWidth int) {
	// TODO calculate scale outside so we are not constantly recomputing
	scale := float64(pixelWidth) / bounds.X
	scaledMin := gs.Player.pos.Scale(scale)
	scaledSize := gs.Player.size.Scale(scale)
	scaledMax := scaledMin.Add(scaledSize)

	entityBounds := clip.Rect(image.Rectangle{
		Min: scaledMin.Point(),
		Max: scaledMax.Point(),
	}).Push(gtx.Ops)
	defer entityBounds.Pop()

	green := color.NRGBA{G: 0xFF, A: 0xFF}
	paint.ColorOp{Color: green}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)
}
