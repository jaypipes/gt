package base

import "github.com/jaypipes/gt/types"

// SetWhitespace sets the Element's whitespace mode
func (b *Base) SetWhitespace(whitespace types.Whitespace) types.Element {
	b.whitespace = whitespace
	return b
}

// Whitespace returns the Element's whitespace mode
func (b *Base) Whitespace() types.Whitespace {
	return b.whitespace
}
