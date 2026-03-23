package types

import (
	"image/color"
)

type CursorShape int

const (
	CursorShapeDefault CursorShape = iota
	CursorShapeBlock
	CursorShapeUnderline
	CursorShapeBar
)

// Cursor represents the style, shape and position of a cursor on the [Screen].
type Cursor interface {
	// Screen returns the Screen object the Cursor will display on.
	Screen() Screen
	// SetScreen sets the Screen object the Cursor will display on.
	SetScreen(Screen)
	// Position returns the coordinates of the Cursor on the Screen. Will be
	// (-1, -1) when the Cursor is hidden.
	Position() Point
	// SetPosition sets the coordinates of the Cursor.
	SetPosition(Point)
	// Hide is a shortcut for SetPosition(Point(-1, -1))
	Hide()
	// Visible returns whether the Cursor is shown on the Screen.
	Visible() bool
	// Shape returns the CursorShape for the Cursor's cell.
	Shape() CursorShape
	// SetShape sets the CursorShape for the Cursor's cell.
	SetShape(CursorShape)
	// Blink returns whether the Cursor is set to blink.
	Blink() bool
	// SetBlink sets the Cursor's blink behaviour.
	SetBlink(bool)
	// Color returns the color of the Cursor's cell.
	Color() color.Color
	// SetColor sets the color of the Cursor's cell.
	SetColor(color.Color)
}

// CursorWithOption describes an optional varg parameter to [core.cursor.New]
// that modifies the returned Cursor.
type CursorWithOption func(Cursor)
