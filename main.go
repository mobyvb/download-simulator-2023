package main

import (
	"flag"
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func main() {
	flag.Parse()
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var ops op.Ops

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)

				// Configure a label styled to be a heading.
				/*
					title := material.Body1(th, "Download Simulator 2023")
					title.Font.Weight = text.Bold
					title.Alignment = text.Start
					title.MaxLines = 1
					title.Layout(gtx)
				*/

				red := color.NRGBA{R: 0xFF, A: 0xFF}
				// ColorOp sets the brush for painting.
				paint.ColorOp{Color: red}.Add(gtx.Ops)
				// PaintOp paints the configured.
				paint.PaintOp{}.Add(gtx.Ops)

				modal := component.NewModal()
				appbar := component.NewAppBar(modal)
				appbar.Title = "Download Simulator 2023"
				//bar.NavigationIcon, _ = widget.NewIcon(icons.ActionHome)
				//bar.Layout(gtx, th, "Download S", "asdf")
				/*
					headingLabel := material.Body1(th, "")
					headingLabel.Font.Weight = text.Bold
					headingLabel.Alignment = text.Middle
					headingLabel.MaxLines = 1
					inset := layout.UniformInset(unit.Dp(6))
				*/

				bar := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return appbar.Layout(gtx, th, "asdf1", "adsf2")
				})

				/*
					content := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						borderWidth := float32(0.5)
						borderColor := color.NRGBA{A: 107}
						return widget.Border{
							Color:        borderColor,
							CornerRadius: unit.Dp(4),
							Width:        unit.Dp(borderWidth),
						}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(4)).Layout(gtx,
								widget.Text("adsfsf"))
						})
					})
				*/

				flex := layout.Flex{Axis: layout.Vertical}
				flex.Layout(gtx, bar)

				e.Frame(gtx.Ops)
			}
		}
	}
}
