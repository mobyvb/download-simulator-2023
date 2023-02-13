package main

import (
	"flag"
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"github.com/mobyvb/download-simulator-2023/game"
)

func main() {
	flag.Parse()
	ui := NewUI()
	go func() {
		w := app.NewWindow()
		if err := ui.Run(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type UI struct {
	Theme  *material.Theme
	Modal  *component.ModalLayer
	AppBar *component.AppBar
}

func NewUI() *UI {
	th := material.NewTheme(gofont.Collection())
	ui := &UI{}
	ui.Theme = th
	ui.Modal = component.NewModal()
	ui.AppBar = component.NewAppBar(ui.Modal)
	ui.AppBar.Title = "Download Simulator"
	return ui
}

func (ui *UI) Run(w *app.Window) error {
	var ops op.Ops

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				// I guess thsi will result in re-calling frame
				op.InvalidateOp{At: gtx.Now.Add(30 * time.Millisecond)}.Add(gtx.Ops)

				ui.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}
}

func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ui.AppBar.Layout(gtx, ui.Theme, "nav description", "overflow description")
		}),
		layout.Flexed(1, ui.Canvas),
	)
}

const canvasHeightRatio = 3
const canvasWidthRatio = 12

var (
	started     bool
	currentGame game.GameState
	prevTime    time.Time
)

func (ui *UI) Canvas(gtx layout.Context) layout.Dimensions {
	if !started {
		currentGame = game.NewGame()
		prevTime = gtx.Now
		started = true
	}
	now := gtx.Now
	dt := now.Sub(prevTime)
	prevTime = now
	currentGame = currentGame.Update(dt)

	max := gtx.Constraints.Max

	canvasWidth := max.X - 50
	canvasHeight := canvasWidth / canvasWidthRatio * canvasHeightRatio
	canvasBounds := image.Point{X: canvasWidth, Y: canvasHeight}

	bounds := clip.Rect(image.Rectangle{Max: canvasBounds}).Push(gtx.Ops)
	defer bounds.Pop()
	red := color.NRGBA{R: 0xFF, A: 0xFF}
	paint.ColorOp{Color: red}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	currentGame.Draw(gtx, canvasWidth)

	return layout.Dimensions{Size: canvasBounds}
}
