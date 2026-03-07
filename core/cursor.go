package core

import (
	"image/color"

	"github.com/jaypipes/gt/types"
)

// Cursor represents the style, shape and position of a cursor on the [Screen].
type Cursor struct {
	// pos is the coordinates of the Cursor on the Screen. This will be (-1,
	// -1) when the Cursor is hidden.
	pos types.Point
	// shape is the CursorShape for the Cursor's cell.
	shape types.CursorShape
	// blink is true if the Cursor should blink.
	blink bool
	// color is the color of the Cursor's cell.
	color color.Color
}

// Position returns the coordinates of the Cursor on the Screen. This will be (-1,
// -1) when the Cursor is hidden.
func (c *Cursor) Position() types.Point {
	return c.pos
}

// SetPosition sets the coordinates of the Cursor.
func (c *Cursor) SetPosition(pt types.Point) {
	c.pos = pt
}

// Hide is a shortcut for SetPosition(Point(-1, -1))
func (c *Cursor) Hide() {
	c.pos = types.Point{X: -1, Y: -1}
}

// Visible returns whether the Cursor is shown on the Screen.
func (c *Cursor) Visible() bool {
	return c.pos.X > -1 && c.pos.Y > -1
}

// Shape returns the CursorShape for the Cursor's cell.
func (c *Cursor) Shape() types.CursorShape {
	return c.shape
}

// SetShape sets the CursorShape for the Cursor's cell.
func (c *Cursor) SetShape(shape types.CursorShape) {
	c.shape = shape
}

// Blink returns whether the Cursor is set to blink.
func (c *Cursor) Blink() bool {
	return c.blink
}

// SetBlink sets the Cursor's blink behaviour.
func (c *Cursor) SetBlink(blink bool) {
	c.blink = blink
}

// Color returns the color of the Cursor's cell.
func (c *Cursor) Color() color.Color {
	return c.color
}

// SetColor sets the color of the Cursor's cell.
func (c *Cursor) SetColor(color color.Color) {
	c.color = color
}

var _ types.Cursor = (*Cursor)(nil)
