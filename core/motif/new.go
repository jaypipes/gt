package motif

import (
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a Motif.
//
// You can pass zero or more MotifWithOptions to optionally set certain
// attributes on the returned Motifd.
func New(opts ...types.MotifWithOption) *Motif {
	m := &Motif{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// Empty returns a new empty Motif.
func Empty() *Motif {
	return &Motif{}
}

// WithNormalStyle sets the styling for when the thing is not being hovered
// over with a mouse, does not have the focus and is not disabled.
// WithNormalStyle then returns the Motif.
func WithNormalStyle(s types.Style) types.MotifWithOption {
	return func(m types.Motif) {
		m.SetNormalStyle(s)
	}
}

// WithNormalBorder sets the border for when the thing is not being hovered
// over with a mouse, does not have the focus and is not disabled.
// WithNormalBorder then returns the Motif.
func WithNormalBorder(b types.Border) types.MotifWithOption {
	return func(m types.Motif) {
		m.SetNormalBorder(b)
	}
}

// WithDisabledStyle sets the styling for when the thing is disabled and
// returns the Motif.
func WithDisabledStyle(s types.Style) types.MotifWithOption {
	return func(m types.Motif) {
		m.SetDisabledStyle(s)
	}
}

// WithDisabledBorder sets the border for when the thing is disabled and
// returns the Motif.
func WithDisabledBorder(b types.Border) types.MotifWithOption {
	return func(m types.Motif) {
		m.SetDisabledBorder(b)
	}
}

// WithFocusedStyle sets the styling for when the thing has the focus and
// returns the Motif.
func WithFocusedStyle(s types.Style) types.MotifWithOption {
	return func(m types.Motif) {
		m.SetFocusedStyle(s)
	}
}

// WithFocusedBorder sets the border for when the thing has the focus and
// returns the Motif.
func WithFocusedBorder(b types.Border) types.MotifWithOption {
	return func(m types.Motif) {
		m.SetFocusedBorder(b)
	}
}

// WithHoveredStyle sets the styling for when the mouse is hovering over the
// thing and returns the Motif.
func WithHoveredStyle(s types.Style) types.MotifWithOption {
	return func(m types.Motif) {
		m.SetHoveredStyle(s)
	}
}

// WithHoveredBorder sets the border for when the mouse is hovering over the
// thing and returns the Motif.
func WithHoveredBorder(b types.Border) types.MotifWithOption {
	return func(m types.Motif) {
		m.SetHoveredBorder(b)
	}
}
