package base

import "github.com/jaypipes/gt/types"

// WithWhitespace sets the Element's whitespace mode and returns the Element.
func (b *Base) WithWhitespace(whitespace types.Whitespace) types.Element {
	b.SetWhitespace(whitespace)
	return b
}
