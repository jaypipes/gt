package core

import (
	"fmt"

	"github.com/jaypipes/gt/core/types"
)

// Plotted describes something that has a position on a Screen. The zero value
// for a Plotted describes something that has no border, no padding and uses
// relative positioning with no relative offset from the previous sibling or
// parent (i.e. a (0,0) relative anchor coordinate.
//
// A Plotted can be thought of as a pair of bounding boxes that represent the
// outer box around the Plotted's padding and content and an inner box around
// the Plotted's content, inside the padding.
//
// A Plotted's outer bounding box is the box around which any border will be
// drawn.
type Plotted struct {
	Aligned
	Bordered
	Bounded
	Displayed
	Padded
	Sized
	// position is the anchoring coordinates (i.e. the (Min.X, Min.Y)) of the
	// absolute or relative position for the Plotted.
	position types.Point
	// absolute is true if the Plotted is using absolute coordinates, false if
	// using relative positioning.
	absolute bool
}

func (p *Plotted) String() string {
	return fmt.Sprintf(
		"position=%s absolute=%t %s %s %s %s",
		p.position, p.absolute,
		p.Bounded.String(), p.Sized.String(),
		p.Displayed.String(), p.Aligned.String(),
	)
}

// Anchor sets the Plotted's anchor point (i.e. its top-left grid coordinates)
// and marks the Plotted as using absolute positioning.
func (p *Plotted) Anchor(pt types.Point) {
	p.position = pt
	p.absolute = true
}

// AbsolutePositioned returns true if the Plotted used absolute positioning.
func (p *Plotted) AbsolutePositioned() bool {
	return p.absolute
}

// InnerBounds returns the inner bounding box for the Plotted, which accounts
// for any border and padding.
func (p *Plotted) InnerBounds() types.Rectangle {
	bounds := p.Bounds()
	border := p.Border()
	if border != nil {
		bounds.Min.X++
		bounds.Min.Y++
		bounds.Max.X--
		bounds.Max.Y--
	}
	return p.Padding().AdjustBounds(bounds)
}

// LeftMargin returns the distance, in cells, from the left edge of the
// outer bounds to left edge of the inner bounds.
func (p *Plotted) LeftMargin() int {
	bounds := p.Bounds()
	inner := p.InnerBounds()
	return inner.Min.X - bounds.Min.X
}

// RightMargin returns the distance, in cells, from the right edge of the
// outer bounds to right edge of the inner bounds.
func (p *Plotted) RightMargin() int {
	bounds := p.Bounds()
	inner := p.InnerBounds()
	return bounds.Max.X - inner.Max.X
}

// TopMargin returns the distance, in cells, from the top edge of the
// outer bounds to top edge of the inner bounds.
func (p *Plotted) TopMargin() int {
	bounds := p.Bounds()
	inner := p.InnerBounds()
	return inner.Min.Y - bounds.Min.Y
}

// BottomMargin returns the distance, in cells, from the bottom edge of the
// outer bounds to bottom edge of the inner bounds.
func (p *Plotted) BottomMargin() int {
	bounds := p.Bounds()
	inner := p.InnerBounds()
	return bounds.Max.Y - inner.Max.Y
}

var _ types.Plotted = (*Plotted)(nil)
