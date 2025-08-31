package ebitengl

import (
	"go-2d/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type ebInput struct{}

func newInput() *ebInput { return &ebInput{} }

func (i *ebInput) KeyDown(k input.Key) bool {
	switch k {
	case input.Left:
		return ebiten.IsKeyPressed(ebiten.KeyArrowLeft)
	case input.Right:
		return ebiten.IsKeyPressed(ebiten.KeyArrowRight)
	case input.Up:
		return ebiten.IsKeyPressed(ebiten.KeyArrowUp)
	case input.Down:
		return ebiten.IsKeyPressed(ebiten.KeyArrowDown)
	case input.Space:
		return ebiten.IsKeyPressed(ebiten.KeySpace)
	case input.Esc:
		return ebiten.IsKeyPressed(ebiten.KeyEscape)
	default:
		return false
	}
}

func (i *ebInput) MousePosition() (x, y float64) {
	xi, yi := ebiten.CursorPosition()
	return float64(xi), float64(yi)
}

func (i *ebInput) MouseDown(b input.MouseButton) bool {
	switch b {
	case input.MouseLeft:
		return ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	case input.MouseRight:
		return ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	case input.MouseMiddle:
		return ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle)
	default:
		return false
	}
}
