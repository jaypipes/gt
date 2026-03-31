package theme

import (
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a Theme.
//
// You can pass zero or more ThemeWithOptions to optionally set certain
// attributes on the returned Theme.
func New(opts ...types.ThemeWithOption) *Theme {
	p := &Theme{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

// WithMotif sets the Theme's Motif for a given ThemeClass.
func WithMotif(class types.ThemeClass, m types.Motif) types.ThemeWithOption {
	return func(t types.Theme) {
		t.SetMotif(class, m)
	}
}

// WithStyle sets the Theme's Style for a given ThemeClass.
func WithStyle(class types.ThemeClass, s types.Style) types.ThemeWithOption {
	return func(t types.Theme) {
		t.SetStyle(class, s)
	}
}

// WithBorder sets the Theme's Border for a given ThemeClass.
func WithBorder(class types.ThemeClass, b types.Border) types.ThemeWithOption {
	return func(t types.Theme) {
		t.SetBorder(class, b)
	}
}
