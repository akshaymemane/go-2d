package main

import (
	"io/ioutil"
	"log"

	"go-2d/core"
	"go-2d/gfx"
)

type Demo struct {
	img gfx.Image
	txt gfx.Text
}

func (d *Demo) Init(ctx *core.Context) error {
	// load a PNG (put a file at examples/text-and-image/assets/ship.png)
	if img, err := ctx.Gfx.LoadImage("examples/text-and-image/assets/ship.png"); err == nil {
		d.img = img
	}

	// load a TTF and create text
	data, err := ioutil.ReadFile("examples/text-and-image/assets/Inter-Regular.ttf")
	if err != nil {
		return err
	}
	font, err := ctx.Gfx.NewFont(data, 24)
	if err != nil {
		return err
	}
	t, err := ctx.Gfx.NewText(font, "Hello, Go 2D!")
	if err != nil {
		return err
	}
	d.txt = t
	return nil
}

func (d *Demo) Update(dt float64) error { return nil }

func (d *Demo) Draw(g gfx.Device) error {
	g.Clear(gfx.Color{0.1, 0.1, 0.12, 1})
	if d.txt != nil {
		g.DrawText(d.txt, 100, 80)
	}
	if d.img != nil {
		g.Draw(d.img, &gfx.DrawOptions{X: 100, Y: 120, ScaleX: 1, ScaleY: 1})
	}
	return nil
}

func main() {
	if err := core.Run(&Demo{}, core.Options{Title: "Text & Image", Width: 1280, Height: 720, VSync: true}); err != nil {
		log.Fatal(err)
	}
}
