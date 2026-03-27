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

// Border returns the Element's appropriate Border. If the Element has the
// focus and a FocusBorder, this returns the FocusBorder. If the mouse is
// hovering over the Element and there is a non-nil HoverBorder, this returns
// the HoverBorder. Otherwise, this returns the normal Border, if any.
func (e *Element) Border() types.Border {
	border := e.border
	if e.focused && e.focusBorder != nil {
		border = e.focusBorder
	} else if e.hovered && e.hoverBorder != nil {
		border = e.hoverBorder
	}
	return border
}

// SetBorder sets the Element's normal border. The normal border is Element's
// border when there is no FocusBorder or HoverBorder active.
func (e *Element) SetBorder(border types.Border) {
	e.border = border
	e.Box.SetBorder(border)
}

// WithBorder sets the Element's border and returns the Element.
func (e *Element) WithBorder(border types.Border) types.Element {
	e.SetBorder(border)
	return e
}

// FocusBorder returns the Element's border when the Element has the focus.
func (e *Element) FocusBorder() types.Border {
	return e.focusBorder
}

// SetFocusBorder sets the Element's border when the Element has the focus.
func (e *Element) SetFocusBorder(border types.Border) {
	e.focusBorder = border
}

// WithFocusBorder sets the Element's border when the Element has the focus and
// returns the Element.
func (e *Element) WithFocusBorder(border types.Border) types.Element {
	e.SetFocusBorder(border)
	return e
}

// HoverBorder returns the Element's border when the mouse is hovering over the
// Element.
func (e *Element) HoverBorder() types.Border {
	return e.hoverBorder
}

// SetHoverBorder sets the Element's border when the mouse is hovering over the
// Element.
func (e *Element) SetHoverBorder(border types.Border) {
	e.hoverBorder = border
}

// WithHoverBorder sets the Element's border when the mouse is hovering over
// the Element and returns the Element.
func (e *Element) WithHoverBorder(border types.Border) types.Element {
	e.SetHoverBorder(border)
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
