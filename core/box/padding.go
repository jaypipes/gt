package box

import (
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// SetPadding sets the Box's padding.
func (b *Box) SetPadding(padding types.Padding) {
	b.padding = padding
}

// Padding returns the padding for the Box.
func (b *Box) Padding() types.Padding {
	return b.padding
}

// HorizontalSpace returns the number of cells consumed by the Box's
// left-right padding and border.
func (b *Box) HorizontalSpace() types.Dimension {
	space := b.padding.HorizontalSpace()
	if b.border != nil {
		space += render.BorderHorizontalSpace(b.border)
	}
	return space
}

// VerticalSpace returns the number of lines consumed by the Box's
// top-bottom padding and border
func (b *Box) VerticalSpace() types.Dimension {
	space := b.padding.VerticalSpace()
	if b.border != nil {
		space += render.BorderVerticalSpace(b.border)
	}
	return space
}
