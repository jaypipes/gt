package element

import "github.com/jaypipes/gt/types"

// WithBounds sets the Element's outer bounding box and returns the Element.
func (e *Element) WithBounds(bounds types.Rectangle) types.Element {
	e.Box.SetBounds(bounds)
	return e
}

// WithAbsolutePosition sets the Element's outer bounding box's top-left
// coordinates and marks the Element as using absolute positioning and returns
// the Element.
func (e *Element) WithAbsolutePosition(pt types.Point) types.Element {
	e.Box.SetAbsolutePosition(pt)
	return e
}

// WithBorder sets the Element's border and returns the Element.
func (e *Element) WithBorder(border types.Border) types.Element {
	e.Box.SetBorder(border)
	return e
}

// WithBorderForegroundColor sets the Element's border foreground color (i.e
// the color of the border cell's underlying grapheme) and returns the Element.
func (e *Element) WithBorderForegroundColor(c types.Color) types.Element {
	e.Box.SetBorderForegroundColor(c)
	return e
}

// WithBorderBackgroundColor sets the Element's border background color
// (i.e the background color of the border's cells and returns the Element.
func (e *Element) WithBorderBackgroundColor(c types.Color) types.Element {
	e.Box.SetBorderBackgroundColor(c)
	return e
}

// WithPadding sets the Element's padding and returns the Element.
func (e *Element) WithPadding(padding types.Padding) types.Element {
	e.Box.SetPadding(padding)
	return e
}
