package element

import (
	"github.com/jaypipes/gt/core/motif"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

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
	m := e.motif
	if m == nil {
		return nil
	}
	if e.disabled {
		db := m.DisabledBorder()
		if db != nil {
			return db
		}
	}
	if e.focused {
		fb := m.FocusedBorder()
		if fb != nil {
			return fb
		}
	}
	if e.hovered {
		hb := m.HoveredBorder()
		if hb != nil {
			return hb
		}
	}
	return m.NormalBorder()
}

// SetBorder sets the Element's normal border. The normal border is Element's
// border when the Element does not have the focus, is not disabled and does
// not have the mouse hovering over it.
func (e *Element) SetBorder(border types.Border) {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetNormalBorder(border)
	e.Box.SetBorder(border)
}

// WithBorder sets the Element's normal border and returns the Element.
func (e *Element) WithBorder(border types.Border) types.Element {
	e.SetBorder(border)
	return e
}

// DisabledBorder returns the Element's border when the Element is disabled.
func (e *Element) DisabledBorder() types.Border {
	if e.motif == nil {
		return nil
	}
	return e.motif.DisabledBorder()
}

// SetDisabledBorder sets the Element's border when the Element is disabled.
func (e *Element) SetDisabledBorder(border types.Border) {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetDisabledBorder(border)
}

// WithDisabledBorder sets the Element's border when the Element is disabled
// and returns the Element.
func (e *Element) WithDisabledBorder(border types.Border) types.Element {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetDisabledBorder(border)
	return e
}

// FocusedBorder returns the Element's border when the Element has the focus.
func (e *Element) FocusedBorder() types.Border {
	if e.motif == nil {
		return nil
	}
	return e.motif.FocusedBorder()
}

// SetFocusedBorder sets the Element's border when the Element has the focus.
func (e *Element) SetFocusedBorder(border types.Border) {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetFocusedBorder(border)
}

// WithFocusedBorder sets the Element's border when the Element has the focus
// and returns the Element.
func (e *Element) WithFocusedBorder(border types.Border) types.Element {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetFocusedBorder(border)
	return e
}

// HoveredBorder returns the Element's border when the mouse is hovering over
// the Element.
func (e *Element) HoveredBorder() types.Border {
	if e.motif == nil {
		return nil
	}
	return e.motif.HoveredBorder()
}

// SetHoveredBorder sets the Element's border when the mouse is hovering over
// the Element.
func (e *Element) SetHoveredBorder(border types.Border) {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetHoveredBorder(border)
}

// WithHoveredBorder sets the Element's border when the mouse is hovering over
// the Element and returns the Element.
func (e *Element) WithHoveredBorder(border types.Border) types.Element {
	if e.motif == nil {
		e.motif = motif.Empty()
	}
	e.motif.SetHoveredBorder(border)
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

// HorizontalSpace returns the number of cells consumed by the Element's
// left-right padding and border.
func (e *Element) HorizontalSpace() types.Dimension {
	padding := e.Box.Padding()
	space := padding.HorizontalSpace()
	border := e.Border()
	if border != nil {
		space += render.BorderHorizontalSpace(border)
	}
	return space
}

// VerticalSpace returns the number of lines consumed by the Element's
// top-bottom padding and border
func (e *Element) VerticalSpace() types.Dimension {
	padding := e.Box.Padding()
	space := padding.VerticalSpace()
	border := e.Border()
	if border != nil {
		space += render.BorderVerticalSpace(border)
	}
	return space
}
