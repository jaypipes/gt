package base

import "github.com/jaypipes/gt/types"

// SetPadding sets the Element's padding.
func (b *Base) SetPadding(padding types.Padding) types.Element {
	b.padding = padding
	return b
}

// Padding returns the padding for the Element.
func (b *Base) Padding() types.Padding {
	return b.padding
}
