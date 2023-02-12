package main

import (
	"flag"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"
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
				ui.Layout(gtx)
				e.Frame(gtx.Ops)
			}
		}
	}
}

func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return ui.AppBar.Layout(gtx, ui.Theme, "asdf1", "adsf2")
		}),
		layout.Flexed(1, ui.Canvas),
	)
}

const canvasHeightRatio = 3
const canvasWidthRatio = 12

func (ui *UI) Canvas(gtx layout.Context) layout.Dimensions {
	max := gtx.Constraints.Max
	canvasWidth := max.X - 50
	canvasHeight := canvasWidth / canvasWidthRatio * canvasHeightRatio
	canvasBounds := image.Point{X: canvasWidth, Y: canvasHeight}
	bounds := clip.Rect(image.Rectangle{Max: canvasBounds}).Push(gtx.Ops)
	defer bounds.Pop()

	red := color.NRGBA{R: 0xFF, A: 0xFF}
	paint.ColorOp{Color: red}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Dimensions{Size: canvasBounds}
}
