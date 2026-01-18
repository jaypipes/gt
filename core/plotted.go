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
	Displayed
	Padded
	Sized
	// bounds is the outer bounding box and positioning coordinates of the
	// Plotted
	bounds types.Rectangle
	// absolute is true if the Plotted is using absolute coordinates, false if
	// using relative positioning.
	absolute bool
}

func (p *Plotted) String() string {
	return fmt.Sprintf(
		"absolute=%t bounds=%s %s %s",
		p.absolute, p.bounds, p.Displayed.String(), p.Aligned.String(),
	)
}

// Bounds returns the Plotted's outer bounding box.
func (p *Plotted) Bounds() types.Rectangle {
	return p.bounds
}

// SetBounds sets the Plotted's outer bounding box.
func (p *Plotted) SetBounds(bounds types.Rectangle) {
	p.bounds = bounds
}

// TL returns the Plotted's outer bounding box's top-left coordinates.
func (p *Plotted) TL() types.Point {
	return p.bounds.Min
}

// TR returns the Plotted's outer bounding box's top-right coordinates.
func (p *Plotted) TR() types.Point {
	return types.Point{
		X: p.bounds.Max.X,
		Y: p.bounds.Min.Y,
	}
}

// MinY returns the Min Y (top) of the Plotted's outer bounding box.
func (p *Plotted) MinY() int {
	return p.bounds.Min.Y
}

// MaxY returns the Max Y (bottom) of the Plotted's outer bounding box.
func (p *Plotted) MaxY() int {
	return p.bounds.Max.Y
}

// SetAbsolutePosition sets the Plotted's outer bounding box's top-left
// coordinates and marks the Plotted as using absolute positioning.
func (p *Plotted) SetAbsolutePosition(pt types.Point) {
	p.bounds.Min = pt
	p.absolute = true
}

// HasAbsolutePosition returns true if the Plotted used absolute positioning.
func (p *Plotted) HasAbsolutePosition() bool {
	return p.absolute
}

// InnerBounds returns the inner bounding box for the Plotted, which is the
// outer bounding box adjusted for any border and padding.
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
