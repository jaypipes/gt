package base

import "github.com/jaypipes/gt/types"

// SetBounds sets the Element's outer bounding box.
func (b *Base) SetBounds(bounds types.Rectangle) types.Element {
	b.Box.SetBounds(bounds)
	return b
}

// SetAbsolutePosition sets the Element's outer bounding box's top-left
// coordinates and marks the Element as using absolute positioning.
func (b *Base) SetAbsolutePosition(pt types.Point) types.Element {
	b.Box.SetAbsolutePosition(pt)
	return b
}

// SetBorder sets the Element's border.
func (b *Base) SetBorder(border types.Border) types.Element {
	b.Box.SetBorder(border)
	return b
}

// SetBorderForegroundColor sets the Element's border foreground color
// (i.e the color of the border cell's underlying grapheme).
func (b *Base) SetBorderForegroundColor(c types.Color) types.Element {
	b.Box.SetBorderForegroundColor(c)
	return b
}

// SetBorderBackgroundColor sets the Element's border background color
// (i.e the background color of the border's cells.
func (b *Base) SetBorderBackgroundColor(c types.Color) types.Element {
	b.Box.SetBorderBackgroundColor(c)
	return b
}

// SetPadding sets the Element's padding.
func (b *Base) SetPadding(padding types.Padding) types.Element {
	b.Box.SetPadding(padding)
	return b
}
