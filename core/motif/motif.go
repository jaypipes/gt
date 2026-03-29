package motif

import "github.com/jaypipes/gt/types"

// Motif is a design pattern for an Element that encapsulates different styles
// and borders when the Element is in various states (focused, being hovered
// over with the mouse, disabled, etc)
type Motif struct {
	// normalStyle contains the styling for when the thing is not being hovered
	// over with a mouse, does not have the focus and is not disabled.
	normalStyle types.Style
	// normalBorder contains the border for when the thing is not being hovered
	// over with a mouse, does not have the focus and is not disabled.
	normalBorder types.Border
	// disabledStyle contains the styling for when the thing is disabled.
	disabledStyle types.Style
	// disabledBorder contains the border for when the thing is disabled.
	disabledBorder types.Border
	// focusedStyle contains the styling for when the thing has the focus.
	focusedStyle types.Style
	// focusedBorder contains the border for when the thing has the focus.
	focusedBorder types.Border
	// hoveredStyle contains the styling for when the mouse is hovering over
	// the thing.
	hoveredStyle types.Style
	// hoveredBorder contains the border for when the mouse is hovering over
	// the thing.
	hoveredBorder types.Border
}

// Unstyled returns true if the Motif has no styling set up.
func (m *Motif) Unstyled() bool {
	return (m.normalStyle == nil || m.normalStyle.Unstyled()) &&
		(m.disabledStyle == nil || m.disabledStyle.Unstyled()) &&
		(m.focusedStyle == nil || m.focusedStyle.Unstyled()) &&
		(m.hoveredStyle == nil || m.hoveredStyle.Unstyled())
}

// NormalStyle returns the styling for when the thing is not being hovered
// over with a mouse, does not have the focus and is not disabled.
func (m *Motif) NormalStyle() types.Style {
	return m.normalStyle
}

// SetNormalStyle sets the styling for when the thing is not being hovered
// over with a mouse, does not have the focus and is not disabled.
func (m *Motif) SetNormalStyle(s types.Style) {
	m.normalStyle = s
}

// WithNormalStyle sets the styling for when the thing is not being hovered
// over with a mouse, does not have the focus and is not disabled.
// WithNormalStyle then returns the Motif.
func (m *Motif) WithNormalStyle(s types.Style) types.Motif {
	m.SetNormalStyle(s)
	return m
}

// NormalBorder returns the border for when the thing is not being hovered
// over with a mouse, does not have the focus and is not disabled.
func (m *Motif) NormalBorder() types.Border {
	return m.normalBorder
}

// SetNormalBorder sets the border for when the thing is not being hovered
// over with a mouse, does not have the focus and is not disabled.
func (m *Motif) SetNormalBorder(b types.Border) {
	m.normalBorder = b
}

// WithNormalBorder sets the border for when the thing is not being hovered
// over with a mouse, does not have the focus and is not disabled.
// WithNormalBorder then returns the Motif.
func (m *Motif) WithNormalBorder(b types.Border) types.Motif {
	m.SetNormalBorder(b)
	return m
}

// DisabledStyle returns the styling for when the thing is disabled.
func (m *Motif) DisabledStyle() types.Style {
	return m.disabledStyle
}

// SetDisabledStyle sets the styling for when the thing is disabled.
func (m *Motif) SetDisabledStyle(s types.Style) {
	m.disabledStyle = s
}

// WithDisabledStyle sets the styling for when the thing is disabled and
// returns the Motif.
func (m *Motif) WithDisabledStyle(s types.Style) types.Motif {
	m.SetDisabledStyle(s)
	return m
}

// DisabledBorder returns the border for when the thing is disabled.
func (m *Motif) DisabledBorder() types.Border {
	return m.disabledBorder
}

// SetDisabledBorder sets the border for when the thing is disabled.
func (m *Motif) SetDisabledBorder(b types.Border) {
	m.disabledBorder = b
}

// WithDisabledBorder sets the border for when the thing is disabled and
// returns the Motif.
func (m *Motif) WithDisabledBorder(b types.Border) types.Motif {
	m.SetDisabledBorder(b)
	return m
}

// FocusedStyle returns the styling for when the thing has the focus.
func (m *Motif) FocusedStyle() types.Style {
	return m.focusedStyle
}

// SetFocusedStyle sets the styling for when the thing has the focus.
func (m *Motif) SetFocusedStyle(s types.Style) {
	m.focusedStyle = s
}

// WithFocusedStyle sets the styling for when the thing has the focus and
// returns the Motif.
func (m *Motif) WithFocusedStyle(s types.Style) types.Motif {
	m.SetFocusedStyle(s)
	return m
}

// FocusedBorder returns the border for when the thing has the focus.
func (m *Motif) FocusedBorder() types.Border {
	return m.focusedBorder
}

// SetFocusedBorder sets the border for when the thing has the focus.
func (m *Motif) SetFocusedBorder(b types.Border) {
	m.focusedBorder = b
}

// WithFocusedBorder sets the border for when the thing has the focus and
// returns the Motif.
func (m *Motif) WithFocusedBorder(b types.Border) types.Motif {
	m.SetFocusedBorder(b)
	return m
}

// HoveredStyle returns the styling for when the mouse is hovering over
// the thing.
func (m *Motif) HoveredStyle() types.Style {
	return m.hoveredStyle
}

// SetHoveredStyle sets the styling for when the mouse is hovering over the
// thing.
func (m *Motif) SetHoveredStyle(s types.Style) {
	m.hoveredStyle = s
}

// WithHoveredStyle sets the styling for when the mouse is hovering over the
// thing and returns the Motif.
func (m *Motif) WithHoveredStyle(s types.Style) types.Motif {
	m.SetHoveredStyle(s)
	return m
}

// HoveredBorder returns the border for when the mouse is hovering over
// the thing.
func (m *Motif) HoveredBorder() types.Border {
	return m.hoveredBorder
}

// SetHoveredBorder sets the border for when the mouse is hovering over the
// thing.
func (m *Motif) SetHoveredBorder(b types.Border) {
	m.hoveredBorder = b
}

// WithHoveredBorder sets the border for when the mouse is hovering over the
// thing and returns the Motif.
func (m *Motif) WithHoveredBorder(b types.Border) types.Motif {
	m.SetHoveredBorder(b)
	return m
}

var _ types.Motif = (*Motif)(nil)
