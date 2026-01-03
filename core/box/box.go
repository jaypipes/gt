package box

import (
	"context"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/types"
)

// New returns a new instance of a Box.
func New() *Box {
	return &Box{}
}

// Box is the simplest types.Renderable element. All components inherit from
// Box.
type Box struct {
	core.Node
	core.Sized
	core.Positioned
	core.Bordered
	core.Padded
}

// Bounds returns the minimum bounding box within which the Box's content is
// contained.
func (b *Box) Bounds() types.Rectangle {
	bounds := b.Bounded.Bounds()
	size := b.Sized.Size()
	if bounds.Empty() {
		bounds.Max.X = size.W
		bounds.Max.Y = size.H
	} else if !size.Empty() {
		// Set the width and height of the bounding box to the smaller of the
		// bounds width/height itself or the size that has been set on the box.
		bw := bounds.Dx()
		if bw > size.W {
			bounds.Max.X = bounds.Min.X + size.W
		}
		bh := bounds.Dy()
		if bh > size.H {
			bounds.Max.Y = bounds.Min.Y + size.H
		}
	}

	if b.Fixed() {
		fixed := b.FixedPosition()
		bounds.Min.X = fixed.X
		bounds.Min.Y = fixed.Y
	} else {
		offset := b.RelativePosition()
		bounds.Min.X += offset.X
		bounds.Min.Y += offset.Y
		bounds.Max.X += offset.X
		bounds.Max.Y += offset.Y
	}
	return bounds
}

// InnerBounds returns the inner bounding box for the Box, which accounts for
// any border and padding.
func (b *Box) InnerBounds() types.Rectangle {
	outer := b.Bounds()
	border := b.Border()
	if border != nil {
		outer.Min.X++
		outer.Min.Y++
		outer.Max.X--
		outer.Max.Y--
	}

	return b.Padding().Bounds(outer)
}

func (b *Box) Draw(screen types.Screen, area types.Rectangle) {
	// determine the overlapping bounding element and clear its cells before
	// rendering the element.
	bb := render.Overlapping(area, b.Bounds())
	render.Clear(screen, bb)

	// If we have a border, draw it around the outer bounding box.
	border := b.Border()
	if border != nil {
		border.Draw(screen, bb)
	}
}

// Render renders the Element to the given buffer at the specified area.
func (b *Box) Render(
	ctx context.Context,
	screen types.Screen,
) {
	bounds := b.Bounds()
	if bounds.Empty() {
		bounds := screen.Bounds()
		b.SetBounds(bounds)
	}
	b.Draw(screen, bounds)
}

var _ types.Renderable = (*Box)(nil)
