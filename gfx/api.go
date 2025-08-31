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

func (o *DrawOptions) Norm() DrawOptions {
	n := *o
	if n.ScaleX == 0 {
		n.ScaleX = 1
	}
	if n.ScaleY == 0 {
		n.ScaleY = 1
	}
	return n
}

// Image is an engine-owned texture handle.
type Image interface{ IsImage() }

// Device is the 2D drawing façade used by games.
type Device interface {
	NewImage(w, h int) (Image, error)
	Clear(c Color)
	Draw(img Image, opts *DrawOptions)
}

// Helper for degrees → radians (nice for samples).
func Deg(v float64) float64 { return v * math.Pi / 180 }
