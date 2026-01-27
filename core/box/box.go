package box

import (
	"context"
	"fmt"
	"sync"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a Box.
func New(ctx context.Context) Box {
	return Box{
		RWMutex: new(sync.RWMutex),
	}
}

// Box is a container for things that can be plotted and rendered to a
// Screen.
//
// Box has an optional border and padding and can be positioned either
// relatively or absolutely on the Screen's grid.
//
// Box has an outer and inner bounding box representing the bounding box
// outside the border and inside the padding of the Box.
type Box struct {
	*sync.RWMutex
	// id is an identifier for the Box.
	id string

	// childIndex is the index of this Box in the parent's children.
	childIndex int
	// parent is the this Node's parent, if any.
	parent types.Plottable
	// children is the collection of Nodes that are the direct children of this
	// Node, if any.
	children []types.Plottable

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

// SetID sets the Box's identifier.
func (b *Box) SetID(id string) {
	b.id = id
}

// ID returns the Box's identifier.
func (b *Box) ID() string {
	return b.id
}

// String returns a short string representation of the Box.
func (b *Box) String() string {
	parentStr := "nil"
	if b.parent != nil {
		parentEl, ok := b.parent.(types.Element)
		if ok {
			parentStr = parentEl.Tag()
		}
	}
	return fmt.Sprintf(
		"child_index=%d parent=%s children=%d "+
			"absolute=%t bounds=%s pad=%s "+
			"display=%s align=%s whitespace=%s",
		b.childIndex, parentStr, len(b.children),
		b.absolute, b.bounds, b.padding,
		b.display, b.alignment, b.whitespace,
	)
}

// Draw implements the uv.Drawable interface
func (b *Box) Draw(screen types.Screen, bounds types.Rectangle) {
	b.DrawBorder(screen)
}

// Render wraps the [uv.Drawablb.Draw] interface method with a context and
// always calls [uv.Drawablb.Draw] with the Rendered's plotted bounds.
func (b *Box) Render(ctx context.Context, screen types.Screen) {
	render.Plot(ctx, b)
	gtlog.Debug(ctx, "box.Box.Render[%s]", b)
	b.Draw(screen, b.Bounds())
	children := b.Children()
	for _, child := range children {
		child.Render(ctx, screen)
	}
}
