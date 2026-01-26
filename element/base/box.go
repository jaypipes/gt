package base

import "github.com/jaypipes/gt/types"

// WithBounds sets the Element's outer bounding box and returns the Element.
func (b *Base) WithBounds(bounds types.Rectangle) types.Element {
	b.Box.SetBounds(bounds)
	return b
}

// WithAbsolutePosition sets the Element's outer bounding box's top-left
// coordinates and marks the Element as using absolute positioning and returns
// the Element.
func (b *Base) WithAbsolutePosition(pt types.Point) types.Element {
	b.Box.SetAbsolutePosition(pt)
	return b
}

// WithBorder sets the Element's border and returns the Element.
func (b *Base) WithBorder(border types.Border) types.Element {
	b.Box.SetBorder(border)
	return b
}

// WithBorderForegroundColor sets the Element's border foreground color (i.e
// the color of the border cell's underlying grapheme) and returns the Element.
func (b *Base) WithBorderForegroundColor(c types.Color) types.Element {
	b.Box.SetBorderForegroundColor(c)
	return b
}

// WithBorderBackgroundColor sets the Element's border background color
// (i.e the background color of the border's cells and returns the Element.
func (b *Base) WithBorderBackgroundColor(c types.Color) types.Element {
	b.Box.SetBorderBackgroundColor(c)
	return b
}

// WithPadding sets the Element's padding and returns the Element.
func (b *Base) WithPadding(padding types.Padding) types.Element {
	b.Box.SetPadding(padding)
	return b
}
