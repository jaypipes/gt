package box

import (
	"context"
	"fmt"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a Box.
func New(ctx context.Context) Box {
	return Box{}
}

// Box can be plotted and rendered to a Screen.
//
// Box has an optional border and padding and can be positioned either
// relatively or absolutely on the Screen's grid.
//
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

	// minWidth is the minimum width of the Element.
	minWidth types.Dimension
	// minHeight is the minimum height of the Element.
	minHeight types.Dimension
	// widthConstraint is the constraint put on the width dimension
	widthConstraint types.DimensionConstraint
	// heightConstraint is the constraint put on the height dimension
	heightConstraint types.DimensionConstraint

	// display is the display mode for the Element.
	display types.Display
	// alignment is the alignment mode of the Element
	alignment types.Alignment
	// whitespace is the whitespace mode of the Element.
	whitespace types.Whitespace
}

// String returns a short string representation of the Box.
func (b *Box) String() string {
	return fmt.Sprintf(
		"absolute=%t bounds=%s pad=%s "+
			"display=%s align=%s whitespace=%s",
		b.absolute, b.bounds, b.padding,
		b.display, b.alignment, b.whitespace,
	)
}

// drawBorder draws the border around the outer bounding box's cells.
func (b *Box) drawBorder(
	ctx context.Context,
	screen types.Screen,
	bounds types.Rectangle,
) {
	// If we have a border, draw it around the outer bounding box.
	border := b.border
	if border == nil {
		return
	}
	gtlog.Debug(ctx, "Box.drawBorder: %s", b)
	style := types.Style{Fg: b.borderFGColor, Bg: b.borderBGColor}
	bb := border.Style(style)
	bb.Draw(screen, b.bounds)
}

// Draw implements the uv.Drawable interface
func (b *Box) Draw(screen types.Screen, bounds types.Rectangle) {
	ctx := context.TODO()
	gtlog.Debug(ctx, "Box.Draw: bounds=%s", bounds)
	b.drawBorder(ctx, screen, bounds)
}

var _ types.Plottable = (*Box)(nil)
var _ types.Drawable = (*Box)(nil)
