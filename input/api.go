package input

// Key is a keyboard key.
type Key int

const (
	Left Key = iota
	Right
	Up
	Down
	Space
	Esc
)

// Device captures per-frame input state.
type Device interface {
	KeyDown(k Key) bool
}
