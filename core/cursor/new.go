package cursor

import (
	"github.com/jaypipes/gt/types"
)

// New returns a new Cursor.
func New(
	opts ...types.CursorWithOption,
) *Cursor {
	e := &Cursor{
		pos: types.Point{X: -1, Y: -1}, // hidden by default
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// WithScreen modifies the returned Cursor, setting its screen to the
// supplied value.
func WithScreen(s types.Screen) types.CursorWithOption {
	return func(e types.Cursor) {
		e.SetScreen(s)
	}
}

// WithPosition modifies the returned Cursor, setting its position to the
// supplied value.
func WithPosition(pos types.Point) types.CursorWithOption {
	return func(e types.Cursor) {
		e.SetPosition(pos)
	}
}

// WithShape modifies the returned Cursor, setting its shape to the supplied
// value.
func WithShape(shape types.CursorShape) types.CursorWithOption {
	return func(e types.Cursor) {
		e.SetShape(shape)
	}
}

// WithBlink modifies the returned Cursor, setting its blink to the
// supplied value.
func WithBlink(on bool) types.CursorWithOption {
	return func(e types.Cursor) {
		e.SetBlink(on)
	}
}

// WithColor modifies the returned Cursor, setting its color to the supplied
// value.
func WithColor(color types.Color) types.CursorWithOption {
	return func(e types.Cursor) {
		e.SetColor(color)
	}
}
