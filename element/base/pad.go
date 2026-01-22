package base

import (
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// SetPadding sets the Element's padding.
func (b *Base) SetPadding(padding types.Padding) types.Element {
	b.padding = padding
	return b
}

// Padding returns the padding for the Element.
func (b *Base) Padding() types.Padding {
	return b.padding
}

// HorizontalSpace returns the number of cells consumed by the element's
// left-right padding and border.
func (b *Base) HorizontalSpace() types.Dimension {
	space := b.padding.HorizontalSpace()
	if b.border != nil {
		space += render.BorderHorizontalSpace(*b.border)
	}
	return space
}

// VerticalSpace returns the number of lines consumed by the element's
// top-bottom padding and border
func (b *Base) VerticalSpace() types.Dimension {
	space := b.padding.VerticalSpace()
	if b.border != nil {
		space += render.BorderVerticalSpace(*b.border)
	}
	return space
}
