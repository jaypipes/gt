package box

import "github.com/jaypipes/gt/types"

// SetWhitespace sets the Box's whitespace mode.
func (b *Box) SetWhitespace(whitespace types.Whitespace) {
	b.whitespace = whitespace
}

// Whitespace returns the Box's whitespace mode.
func (b *Box) Whitespace() types.Whitespace {
	return b.whitespace
}
