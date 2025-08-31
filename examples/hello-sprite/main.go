package main

import (
	"image/color"

	"go-2d/core"
	"go-2d/gfx"
	"go-2d/input"
)

type Hello struct {
	x      float64
	y      float64
	square gfx.Image
	in     input.Device // <-- add this
}

func (h *Hello) Init(ctx *core.Context) error {
	h.in = ctx.Input // <-- keep handle to input
	img, _ := ctx.Gfx.(interface {
		NewSolid(int, int, color.RGBA) (gfx.Image, error)
	}).
		NewSolid(64, 64, color.RGBA{200, 100, 240, 255})
	h.square = img
	return nil
}

func (h *Hello) Update(dt float64) error {
	const speed = 200.0
	if h.in.KeyDown(input.Right) {
		h.x += speed * dt
	}
	if h.in.KeyDown(input.Left) {
		h.x -= speed * dt
	}
	if h.in.KeyDown(input.Up) {
		h.y -= speed * dt
	}
	if h.in.KeyDown(input.Down) {
		h.y += speed * dt
	}
	if h.x > 1200 {
		h.x = 0
	}
	if h.x < 0 {
		h.x = 1200
	}

	if h.y > 700 {
		h.y = 0
	}
	if h.y < 0 {
		h.y = 700
	}

	return nil
}

func (h *Hello) Draw(g gfx.Device) error {
	// Background is set by engine; we just draw our square moving on a sine wave.
	// y := 300 + 50*math.Sin(h.x*0.02)
	g.Draw(h.square, &gfx.DrawOptions{X: h.x, Y: h.y})
	return nil
}

func main() {
	game := &Hello{}
	_ = core.Run(game, core.Options{
		Title: "go-2d — hello sprite",
		Width: 1280, Height: 720,
		VSync:      true,
		Background: gfx.Color{R: 0.08, G: 0.09, B: 0.12, A: 1},
	})
}
