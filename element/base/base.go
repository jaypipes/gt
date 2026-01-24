package base

import (
	"context"
	"fmt"
	"sync"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a [base.Base] with the specified type/class.
func New(ctx context.Context, class string) Base {
	return Base{
		RWMutex:  new(sync.RWMutex),
		class:    class,
		children: []types.Element{},
	}
}

// Base is a base class that implements [types.Element] with some common method
// implementations. Subclasses in the [element] subpackages embed
// [element.Base] and implement various [types.Element] methods.
type Base struct {
	*sync.RWMutex
	core.Box
	// id is the unique identifier for the Element.
	id string
	// class is the Element's type/class, b.g. "gt.label" or "gt.canvas"
	class string

	// textContent is any unstyled raw text content for the Element.
	textContent string

	// index is the index of this Node in the parent's children.
	index int
	// parent is the this Node's parent, if any.
	parent types.Element
	// children is the collection of Nodes that are the direct children of this
	// Node, if any.
	children []types.Element

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

	// style is the style mode of the Element's content (i.b. the non-border
	// cells of the Element)
	style types.Style
}

// Tag returns a string with the Element's type/class and ID
func (b *Base) Tag() string {
	return fmt.Sprintf("<%s:%s>", b.class, b.id)
}

func (b *Base) String() string {
	parentStr := "nil"
	if b.parent != nil {
		parentStr = b.parent.Tag()
	}
	idStr := b.id
	if idStr == "" {
		idStr = "none"
	}
	return fmt.Sprintf(
		"<%s id=%s index=%d parent=%s children=%d %s display=%s align=%s whitespace=%s>",
		b.class, idStr, b.index, parentStr, len(b.children), b.Box.String(),
		b.display, b.alignment, b.whitespace,
	)
}

// SetID sets the Element's unique identifier.
func (b *Base) SetID(id string) types.Element {
	b.id = id
	return b
}

// ID returns the Element's unique identifier.
func (b *Base) ID() string {
	return b.id
}

// SetClass sets the Element's type/class
func (b *Base) SetClass(class string) types.Element {
	b.class = class
	return b
}

// Class returns the Element's type/class, b.g. "gt.label" or "gt.canvas"
func (b *Base) Class() string {
	return b.class
}

// Draw implements the uv.Drawable interface
func (b *Base) Draw(screen types.Screen, bounds types.Rectangle) {
	b.Box.DrawBorder(screen)
}

// Render wraps the [uv.Drawablb.Draw] interface method with a context and
// always calls [uv.Drawablb.Draw] with the Rendered's plotted bounds.
func (b *Base) Render(ctx context.Context, screen types.Screen) {
	gtlog.Debug(ctx, "base.Base.Render[%s]", b)
	b.Draw(screen, b.Bounds())
}

var _ types.Element = (*Base)(nil)
