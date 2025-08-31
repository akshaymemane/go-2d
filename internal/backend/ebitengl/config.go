package ebitengl

import "image/color"

type Config struct {
	Title         string
	Width, Height int
	VSync         bool
	Scale         float64
	BgColor       color.RGBA
}
