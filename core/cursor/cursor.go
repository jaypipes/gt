package cursor

import (
	"image/color"

	"github.com/gdamore/tcell/v3"
	tccolor "github.com/gdamore/tcell/v3/color"

	"github.com/jaypipes/gt/types"
)

// Cursor represents the style, shape and position of a cursor on the [Screen].
type Cursor struct {
	// screen is the underlying terminal screen the Cursor is displayed on.
	screen types.Screen
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

// Screen returns the Screen the Cursor will display on.
func (c *Cursor) Screen() types.Screen {
	return c.screen
}

// SetScreen sets the Screen the Cursor will display on.
func (c *Cursor) SetScreen(s types.Screen) {
	c.screen = s
}

// Position returns the coordinates of the Cursor on the Screen. This will be (-1,
// -1) when the Cursor is hidden.
func (c *Cursor) Position() types.Point {
	return c.pos
}

// SetPosition sets the coordinates of the Cursor.
func (c *Cursor) SetPosition(pt types.Point) {
	c.pos = pt
	if c.screen != nil {
		c.screen.ShowCursor(pt.X, pt.Y)
	}
}

// Hide is a shortcut for SetPosition(Point(-1, -1))
func (c *Cursor) Hide() {
	c.pos = types.Point{X: -1, Y: -1}
	if c.screen != nil {
		c.screen.HideCursor()
	}
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
	c.tcellSetCursorStyle()
}

// Blink returns whether the Cursor is set to blink.
func (c *Cursor) Blink() bool {
	return c.blink
}

// SetBlink sets the Cursor's blink behaviour.
func (c *Cursor) SetBlink(blink bool) {
	c.blink = blink
	c.tcellSetCursorStyle()
}

// Color returns the color of the Cursor's cell.
func (c *Cursor) Color() color.Color {
	return c.color
}

// SetColor sets the color of the Cursor's cell.
func (c *Cursor) SetColor(color color.Color) {
	c.color = color
	c.tcellSetCursorStyle()
}

// tcellSetCursorStyle translates the types.CursorShape and colors/blinking to
// the tcell.CursorStyle and colors.
func (c *Cursor) tcellSetCursorStyle() {
	if c.screen == nil {
		return
	}
	cols := []tccolor.Color{}
	col := c.Color()
	if col != nil {
		cols = append(cols, tccolor.FromImageColor(col))
	}
	blink := c.Blink()
	cs := tcell.CursorStyleDefault
	shape := c.Shape()
	switch shape {
	case types.CursorShapeBlock:
		cs = tcell.CursorStyleSteadyBlock
		if blink {
			cs = tcell.CursorStyleBlinkingBlock
		}
	case types.CursorShapeBar:
		cs = tcell.CursorStyleSteadyBar
		if blink {
			cs = tcell.CursorStyleBlinkingBar
		}
	case types.CursorShapeUnderline:
		cs = tcell.CursorStyleSteadyUnderline
		if blink {
			cs = tcell.CursorStyleBlinkingUnderline
		}
	}
	c.screen.SetCursorStyle(cs, cols...)
}

var _ types.Cursor = (*Cursor)(nil)
