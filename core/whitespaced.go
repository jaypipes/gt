package core

import "github.com/jaypipes/gt/core/types"

// Whitespaced describes something that has a whitespace mode
type Whitespaced struct {
	whitespace types.Whitespace
}

// SetWhitespace sets the Whitespaced's whitespace mode
func (w *Whitespaced) SetWhitespace(whitespace types.Whitespace) {
	w.whitespace = whitespace
}

// Whitespace returns the Whitespaced's whitespace mode
func (w *Whitespaced) Whitespace() types.Whitespace {
	return w.whitespace
}

var _ types.Whitespaced = (*Whitespaced)(nil)
