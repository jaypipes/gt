package core

import (
	"fmt"

	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// Box has an outer and inner bounding box representing the bounding box
// outside the border and inside the padding of the Box.
type Box struct {
	// bounds is the outer bounding box and positioning coordinates of the
	// Box
	bounds types.Rectangle
	// absolute is true if the Box is using absolute coordinates, false if
	// using relative positioning.
	absolute bool
	// padding is any padding applied to the Box.
	padding types.Padding
	// border is the optional Border information for the Box.
	border *types.Border
	// borderFGColor is the border foreground color (i.e the color of the
	// border cell's underlying grapheme).
	borderFGColor types.Color
	// borderBGColor is the border background color, i.b. the background color
	// of the border cells.
	borderBGColor types.Color
}

func (b *Box) String() string {
	return fmt.Sprintf(
		"absolute=%t bounds=%s pad=%s",
		b.absolute, b.bounds, b.padding,
	)
}

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

// SetBorder sets the Box's border.
func (b *Box) SetBorder(border types.Border) {
	b.border = &border
}

// Border returns the Box's border, if any.
func (b *Box) Border() *types.Border {
	return b.border
}

// SetBorderForegroundColor sets the Box's border foreground color
// (i.e the color of the border cell's underlying grapheme).
func (b *Box) SetBorderForegroundColor(c types.Color) {
	b.borderFGColor = c
}

// BorderForegroundColor returns the Box's border foreground color.
func (b *Box) BorderForegroundColor() types.Color {
	return b.borderFGColor
}

// SetBorderBackgroundColor sets the Box's border background color
// (i.e the background color of the border's cells.
func (b *Box) SetBorderBackgroundColor(c types.Color) {
	b.borderBGColor = c
}

// BorderBackgroundColor returns the Box's border background color.
func (b *Box) BorderBackgroundColor() types.Color {
	return b.borderBGColor
}

// DrawBorder draws the border around the outer bounding box's cells.
func (b *Box) DrawBorder(screen types.Screen) {
	// If we have a border, draw it around the outer bounding box.
	border := b.border
	if border == nil {
		return
	}
	style := types.Style{Fg: b.borderFGColor, Bg: b.borderBGColor}
	bb := border.Style(style)
	bb.Draw(screen, b.bounds)
}

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
		space += render.BorderHorizontalSpace(*b.border)
	}
	return space
}

// VerticalSpace returns the number of lines consumed by the Box's
// top-bottom padding and border
func (b *Box) VerticalSpace() types.Dimension {
	space := b.padding.VerticalSpace()
	if b.border != nil {
		space += render.BorderVerticalSpace(*b.border)
	}
	return space
}
