package box

import (
	"github.com/jaypipes/gt/types"
)

// SetBounds sets the Box's outer bounding box.
func (b *Box) SetBounds(bounds types.Rectangle) {
	b.bounds = bounds
}

// Bounds returns the Box's outer bounding box.
func (b *Box) Bounds() types.Rectangle {
	return b.bounds
}

// TL returns the Box's outer bounding box's top-left coordinates.
func (b *Box) TL() types.Point {
	return b.bounds.Min
}

// TR returns the Box's outer bounding box's top-right coordinates.
func (b *Box) TR() types.Point {
	return types.Point{
		X: b.bounds.Max.X,
		Y: b.bounds.Min.Y,
	}
}

// MinY returns the Min Y (top) of the Box's outer bounding box.
func (b *Box) MinY() int {
	return b.bounds.Min.Y
}

// MaxY returns the Max Y (bottom) of the Box's outer bounding box.
func (b *Box) MaxY() int {
	return b.bounds.Max.Y
}

// InnerBounds returns the inner bounding box for the Box, which is the
// outer bounding box adjusted for any border and padding.
func (b *Box) InnerBounds() types.Rectangle {
	bounds := b.Bounds()
	border := b.Border()
	if border != nil {
		bounds.Min.X++
		bounds.Min.Y++
		bounds.Max.X--
		bounds.Max.Y--
	}
	return b.Padding().AdjustBounds(bounds)
}

// SetAbsolutePosition sets the Box's outer bounding box's top-left
// coordinates and marks the Box as using absolute positioning.
func (b *Box) SetAbsolutePosition(pt types.Point) {
	b.bounds.Min = pt
	b.absolute = true
}

// HasAbsolutePosition returns true if the Box used absolute positioning.
func (b *Box) HasAbsolutePosition() bool {
	return b.absolute
}
