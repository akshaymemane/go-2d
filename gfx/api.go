package gfx

import "math"

// Color with [0,1] components.
type Color struct{ R, G, B, A float64 }

func Clamp01(x float64) float64 {
	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	return x
}

// DrawOptions describes how to draw an Image.
type DrawOptions struct {
	X, Y             float64
	ScaleX, ScaleY   float64
	Rotation         float64 // radians
	OriginX, OriginY float64 // pivot
	Tint             Color   // multiply color (defaults to white)
}

// Image is an engine-owned texture handle.
type Image interface{ IsImage() }

// Font is a loaded font face resource.
type Font interface{ IsFont() }

// Text is pre-shaped text tied to a font.
type Text interface{ IsText() }

// Device is the 2D drawing façade used by games.
type Device interface {
	NewImage(w, h int) (Image, error)
	LoadImage(path string) (Image, error)

	Clear(c Color)
	Draw(img Image, opts *DrawOptions)

	// Text (M1)
	NewFont(ttf []byte, size float64) (Font, error)
	NewText(font Font, s string) (Text, error)
	DrawText(t Text, x, y float64)
}

// Helper for degrees → radians (nice for samples).
func Deg(v float64) float64 { return v * math.Pi / 180 }
