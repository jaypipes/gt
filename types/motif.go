package types

// Motif is a design pattern for an Element that encapsulates different styles
// and borders when the Element is in various states (focused, being hovered
// over with the mouse, disabled, etc)
type Motif interface {
	// Unstyled returns true if the Motif has no styling set up.
	Unstyled() bool
	// NormalStyle returns the styling for when the thing is not being hovered
	// over with a mouse, does not have the focus and is not disabled.
	NormalStyle() Style
	// SetNormalStyle sets the styling for when the thing is not being hovered
	// over with a mouse, does not have the focus and is not disabled.
	SetNormalStyle(Style)
	// WithNormalStyle sets the styling for when the thing is not being hovered
	// over with a mouse, does not have the focus and is not disabled.
	// WithNormalStyle returns the Mofif.
	WithNormalStyle(Style) Motif
	// NormalBorder returns the border for when the thing is not being hovered
	// over with a mouse, does not have the focus and is not disabled.
	NormalBorder() Border
	// SetNormalBorder sets the border for when the thing is not being hovered
	// over with a mouse, does not have the focus and is not disabled.
	SetNormalBorder(Border)
	// WithNormalBorder sets the border for when the thing is not being hovered
	// over with a mouse, does not have the focus and is not disabled.
	// WithNormalBorder returns the Motif.
	WithNormalBorder(Border) Motif
	// DisabledStyle returns the styling for when the thing is disabled.
	DisabledStyle() Style
	// SetDisabledStyle sets the styling for when the thing is disabled.
	SetDisabledStyle(Style)
	// WithDisabledStyle sets the styling for when the thing is disabled and
	// returns the Motif.
	WithDisabledStyle(Style) Motif
	// DisabledBorder returns the border for when the thing is disabled.
	DisabledBorder() Border
	// SetDisabledBorder sets the border for when the thing is disabled.
	SetDisabledBorder(Border)
	// WithDisabledBorder sets the border for when the thing is disabled and
	// returns the Motif.
	WithDisabledBorder(Border) Motif
	// FocusedStyle returns the styling for when the thing has the focus.
	FocusedStyle() Style
	// SetFocusedStyle sets the styling for when the thing has the focus.
	SetFocusedStyle(Style)
	// WithFocusedStyle sets the styling for when the thing has the focus and
	// returns the Motif.
	WithFocusedStyle(Style) Motif
	// FocusedBorder returns the border for when the thing has the focus.
	FocusedBorder() Border
	// SetFocusedBorder sets the border for when the thing has the focus.
	SetFocusedBorder(Border)
	// WithFocusedBorder sets the border for when the thing has the focus and
	// returns the Motif.
	WithFocusedBorder(Border) Motif
	// HoveredStyle returns the styling for when the mouse is hovering over
	// the thing.
	HoveredStyle() Style
	// SetHoveredStyle sets the styling for when the mouse is hovering over the
	// thing.
	SetHoveredStyle(Style)
	// WithHoveredStyle sets the styling for when the mouse is hovering over
	// the thing and returns the Motif.
	WithHoveredStyle(Style) Motif
	// HoveredBorder returns the border for when the mouse is hovering over
	// the thing.
	HoveredBorder() Border
	// SetHoveredBorder sets the border for when the mouse is hovering over the
	// thing.
	SetHoveredBorder(Border)
	// WithHoveredBorder sets the border for when the mouse is hovering over
	// the thing and returns the Motif.
	WithHoveredBorder(Border) Motif
}

// MotifWithOption describes an optional varg parameter to [motif.New] that
// modifies the returned Motif.
type MotifWithOption func(Motif)
