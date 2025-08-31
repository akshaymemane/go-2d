package ebitengl

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	ebitext "github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"go-2d/gfx"
)

type ebImage struct{ img *ebiten.Image }

func (ebImage) IsImage() {}

type ebFont struct{ face font.Face }

func (ebFont) IsFont() {}

type ebText struct {
	s    string
	face font.Face
}

func (ebText) IsText() {}

type ebGfx struct{}

func newGfx() *ebGfx { return &ebGfx{} }

func (g *ebGfx) NewImage(w, h int) (gfx.Image, error) {
	img := ebiten.NewImage(w, h)
	return ebImage{img: img}, nil
}

func (g *ebGfx) LoadImage(path string) (gfx.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	im, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	eimg := ebiten.NewImageFromImage(im)
	return ebImage{img: eimg}, nil
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
	// Transform pipeline: origin → scale → rotate → position
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

// Text APIs (M1)
func (g *ebGfx) NewFont(ttf []byte, size float64) (gfx.Font, error) {
	ft, err := opentype.Parse(ttf)
	if err != nil {
		return nil, err
	}
	face, err := opentype.NewFace(ft, &opentype.FaceOptions{Size: size, DPI: 72})
	if err != nil {
		return nil, err
	}
	return ebFont{face: face}, nil
}

func (g *ebGfx) NewText(font gfx.Font, s string) (gfx.Text, error) {
	ef, ok := font.(ebFont)
	if !ok {
		return nil, fmt.Errorf("invalid font type")
	}
	return ebText{s: s, face: ef.face}, nil
}

func (g *ebGfx) DrawText(t gfx.Text, x, y float64) {
	et, ok := t.(ebText)
	if !ok || currentScreen == nil {
		return
	}
	ebitext.Draw(currentScreen, et.s, et.face, int(x), int(y), color.White)
}

// A global render target set by the game loop for the current frame.
var currentScreen *ebiten.Image

// Utility used by samples/tests to make a solid-colored image.
func (g *ebGfx) NewSolid(w, h int, c color.RGBA) (gfx.Image, error) {
	img := ebiten.NewImage(w, h)
	img.Fill(c)
	return ebImage{img: img}, nil
}
