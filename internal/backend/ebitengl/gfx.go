package ebitengl

import (
	"image"
	"image/color"

	"go-2d/gfx"

	"github.com/hajimehoshi/ebiten/v2"
)

type ebImage struct{ img *ebiten.Image }

// Implement gfx.Image
func (ebImage) IsImage() {}

type ebGfx struct{}

func newGfx() *ebGfx { return &ebGfx{} }

func (g *ebGfx) NewImage(w, h int) (gfx.Image, error) {
	img := ebiten.NewImage(w, h)
	return ebImage{img: img}, nil
}

func (g *ebGfx) Clear(c gfx.Color) {
	if currentScreen == nil {
		return
	}
	eb := color.RGBA{
		R: uint8(gfx.Clamp01(c.R) * 255),
		G: uint8(gfx.Clamp01(c.G) * 255),
		B: uint8(gfx.Clamp01(c.B) * 255),
		A: uint8(gfx.Clamp01(c.A) * 255),
	}
	currentScreen.Fill(eb)
}

func (g *ebGfx) Draw(img gfx.Image, opts *gfx.DrawOptions) {
	eimg, ok := img.(ebImage)
	if !ok {
		return
	}
	op := &ebiten.DrawImageOptions{}
	// Normalize options with sensible defaults
	n := gfx.DrawOptions{ScaleX: 1, ScaleY: 1, Tint: gfx.Color{R: 1, G: 1, B: 1, A: 1}}
	if opts != nil {
		n = *opts
		if n.ScaleX == 0 {
			n.ScaleX = 1
		}
		if n.ScaleY == 0 {
			n.ScaleY = 1
		}
		if n.Tint == (gfx.Color{}) {
			n.Tint = gfx.Color{1, 1, 1, 1}
		}
	}
	// Translate origin → scale/rotate → translate to position
	op.GeoM.Translate(-n.OriginX, -n.OriginY)
	op.GeoM.Scale(n.ScaleX, n.ScaleY)
	op.GeoM.Rotate(n.Rotation)
	op.GeoM.Translate(n.X, n.Y)
	// Tint
	op.ColorScale.Scale(float32(n.Tint.R), float32(n.Tint.G), float32(n.Tint.B), float32(n.Tint.A))
	if currentScreen == nil {
		return
	}
	currentScreen.DrawImage(eimg.img, op)
}

// A global render target set by the game loop for the current frame.
var currentScreen *ebiten.Image

// Utility used by samples/tests to make a solid-colored image.
func (g *ebGfx) NewSolid(w, h int, c color.RGBA) (gfx.Image, error) {
	img := ebiten.NewImage(w, h)
	img.Fill(c)
	return ebImage{img: img}, nil
}

// For future: LoadImage(path) using image.Decode.
func decodeToEbitImage(src image.Image) *ebiten.Image {
	b := ebiten.NewImage(src.Bounds().Dx(), src.Bounds().Dy())
	// Convert via DrawImage from an *image.RGBA; omitted for brevity in M0.
	return b
}
