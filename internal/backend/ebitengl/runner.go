package ebitengl

import (
	"go-2d/gfx"
	"go-2d/input"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Adapter bridges engine-agnostic Game to ebiten.Game.
type Adapter struct {
	OnInit   func(ctx RuntimeContext) error
	OnUpdate func(dt float64) error
	OnDraw   func(g gfx.Device) error
}

// RuntimeContext passes devices into Init.
type RuntimeContext struct {
	Gfx   gfx.Device
	Input input.Device
}

func Run(cfg Config, a *Adapter) error {
	ebiten.SetWindowTitle(cfg.Title)
	if cfg.Width <= 0 {
		cfg.Width = 1280
	}
	if cfg.Height <= 0 {
		cfg.Height = 720
	}
	ebiten.SetWindowSize(cfg.Width, cfg.Height)
	if cfg.Scale <= 0 {
		cfg.Scale = 1
	}
	// VSync is on by default in ebiten.

	game := &ebGame{
		adapter: a,
		bg:      cfg.BgColor,
		gfx:     newGfx(),
		input:   newInput(),
	}
	input.SetDevice(game.input)
	return ebiten.RunGame(game)
}

type ebGame struct {
	adapter     *Adapter
	initialized bool
	bg          color.RGBA
	last        time.Time
	gfx         *ebGfx
	input       *ebInput
}

func (g *ebGame) Update() error {
	if !g.initialized {
		g.last = time.Now()
		if err := g.adapter.OnInit(RuntimeContext{Gfx: g.gfx, Input: g.input}); err != nil {
			return err
		}
		g.initialized = true
	}
	now := time.Now()
	dt := now.Sub(g.last).Seconds()
	g.last = now
	if dt > 0.25 {
		dt = 0.25
	} // clamp huge pauses
	return g.adapter.OnUpdate(dt)
}

func (g *ebGame) Draw(screen *ebiten.Image) {
	currentScreen = screen
	// Clear background
	screen.Fill(g.bg)
	_ = g.adapter.OnDraw(g.gfx)
	currentScreen = nil
	// Debug fps
	// ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.1f", ebiten.CurrentFPS()))
	_ = ebitenutil.DebugPrint // keep import for later; removed text in M0
}

func (g *ebGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
