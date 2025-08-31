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

// MouseButton identifies a mouse button.
type MouseButton int

const (
	MouseLeft MouseButton = iota
	MouseRight
	MouseMiddle
)

// Device captures per-frame input state.
type Device interface {
	KeyDown(k Key) bool
	MousePosition() (x, y float64)
	MouseDown(b MouseButton) bool
}

// Global device reference set by backend runner.
var Current Device

func SetDevice(d Device) { Current = d }

// keyDown helpers forward to the Current device.
func KeyDown(k Key) bool {
	if Current == nil {
		return false
	}
	return Current.KeyDown(k)
}

func MousePosition() (float64, float64) {
	if Current == nil {
		return 0, 0
	}
	return Current.MousePosition()
}

func MouseDown(b MouseButton) bool {
	if Current == nil {
		return false
	}
	return Current.MouseDown(b)
}
