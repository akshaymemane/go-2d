package core

import (
	"image/color"
	"log"

	"go-2d/gfx"
	"go-2d/input"
	backend "go-2d/internal/backend/ebitengl"
)

// Game is what your users implement.
type Game interface {
	Init(ctx *Context) error
	Update(dt float64) error // seconds
	Draw(g gfx.Device) error
}

// Options configures the window and runtime.
type Options struct {
	Title         string
	Width, Height int
	VSync         bool
	Scale         float64
	Background    gfx.Color
}

// Context exposes subsystems (no globals).
type Context struct {
	Gfx   gfx.Device
	Input input.Device
}

// Run boots the engine with the given game and options.
func Run(game Game, opts Options) error {
	// Translate Options → backend config.
	bg := color.RGBA{
		R: uint8(gfx.Clamp01(opts.Background.R) * 255),
		G: uint8(gfx.Clamp01(opts.Background.G) * 255),
		B: uint8(gfx.Clamp01(opts.Background.B) * 255),
		A: uint8(gfx.Clamp01(opts.Background.A) * 255),
	}

	cfg := backend.Config{
		Title:   opts.Title,
		Width:   opts.Width,
		Height:  opts.Height,
		VSync:   opts.VSync,
		Scale:   opts.Scale,
		BgColor: bg,
	}

	adapter := &backend.Adapter{
		OnInit: func(ctx backend.RuntimeContext) error {
			c := &Context{Gfx: ctx.Gfx, Input: ctx.Input}
			return game.Init(c)
		},
		OnUpdate: func(dt float64) error { return game.Update(dt) },
		OnDraw:   func(g gfx.Device) error { return game.Draw(g) },
	}

	if err := backend.Run(cfg, adapter); err != nil {
		log.Printf("engine terminated: %v", err)
		return err
	}
	return nil
}
