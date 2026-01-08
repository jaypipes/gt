package core

import (
	"fmt"

	"github.com/jaypipes/gt/core/types"
)

// Bounded describes something that has an absolute bounding box. An absolute
// bounding box is essentially a box for a viewport that allows something to be
// "clipped" to some dimensions.
//
// NOTE(jaypipes): The absolute bounding box is NOT the same as the rectangle
// returned by things like [uv.StyledString.Bounds]. That method returns the
// minimum area that the styled string consumes on the screen. In `gt`, the
// `gt.Plotted.Plot` method is the thing that calculates a Plotted's minimum
// area that contents consumes on the screen.
type Bounded struct {
	// bounds is the absolute bounding box of the Bounded.
	bounds types.Rectangle
}

func (b *Bounded) String() string {
	if b.bounds.Empty() {
		return "bounds=none"
	}
	return fmt.Sprintf("bounds=%s", b.bounds)
}

// SetBounds sets the Bounded's bounding box.
func (b *Bounded) SetBounds(r types.Rectangle) {
	b.bounds = r
}

// Bounds returns the absolute bounding box constraint within which the
// Bounded's content must be contained.
func (b *Bounded) Bounds() types.Rectangle {
	return b.bounds
}

// MinX returns the min X coordinate of the Plotted's outer bounding box.
func (b *Bounded) MinX() int {
	return b.bounds.Min.X
}

// MinY returns the min Y coordinate of the Plotted's outer bounding mbr.
func (b *Bounded) MinY() int {
	return b.bounds.Min.Y
}

// MinX returns the max X coordinate of the Plotted's outer bounding mbr.
func (b *Bounded) MaxX() int {
	return b.bounds.Max.X
}

// MaxY returns the max Y coordinate of the Plotted's outer bounding mbr.
func (b *Bounded) MaxY() int {
	return b.bounds.Max.Y
}

// TL returns the top-left (i.e. (min.x, min.y, or "anchoring") coordinates
// for the Plotted.
func (b *Bounded) TL() types.Point {
	return types.Point{X: b.bounds.Min.X, Y: b.bounds.Min.Y}
}

// TR returns the top-right (i.e. (max.x, min.y) coordinates for the
// Plotted.
func (b *Bounded) TR() types.Point {
	return types.Point{X: b.bounds.Max.X, Y: b.bounds.Min.Y}
}

// BL returns the bottom-left (i.e. (min.x, max.y) coordinates for the
// Plotted.
func (b *Bounded) BL() types.Point {
	return types.Point{X: b.bounds.Min.X, Y: b.bounds.Max.Y}
}

// BR returns the bottom-right (i.e. (max.x, max.y) coordinates for the
// Plotted.
func (b *Bounded) BR() types.Point {
	return types.Point{X: b.bounds.Max.X, Y: b.bounds.Max.Y}
}

var _ types.Bounded = (*Bounded)(nil)
