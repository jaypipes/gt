package element

import (
	"context"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/node"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/types"
)

// New returns a new instance of an Element with the specified type/class.
func New(ctx context.Context, class string) *Element {
	n := node.New(ctx)
	return &Element{
		Node:  *n,
		class: class,
	}
}

// Element is a specialized type of Node that can be sized and styled.
type Element struct {
	core.Bordered
	core.Bounded
	node.Node
	core.Padded
	core.Positioned
	core.Sized
	// id is the unique identifier for the Element.
	id string
	// class is the Element's type/class, e.g. "gt.label" or "gt.canvas"
	class string
}

// SetID sets the Element's unique identifier.
func (e *Element) SetID(id string) {
	e.id = id
}

// ID returns the Element's unique identifier.
func (e *Element) ID() string {
	return e.id
}

// SetClass sets the Element's type/class
func (e *Element) SetClass(class string) {
	e.class = class
}

// Class returns the Element's type/class, e.g. "gt.label" or "gt.canvas"
func (e *Element) Class() string {
	return e.class
}

// SetBounds propogates a bounding box constraint down to all child elements.
func (e *Element) SetBounds(bounds types.Rectangle) {
	ctx := context.TODO()
	e.Debug(ctx, "Element([%s]%s).SetBounds: bounding box %s\n", e.class, e.id, bounds)
	e.Bounded.SetBounds(bounds)
	propogate := func(ctx context.Context, child types.Node) {
		el, ok := child.(types.Element)
		if !ok {
			// it might be a Bounded (which doesn't propogate bounds to
			// children...
			bounded, ok := child.(types.Bounded)
			if !ok {
				// Not a bounded or element, just ignore
				return
			}
			bounded.SetBounds(bounds)
		}
		el.SetBounds(bounds)
	}
	e.VisitChildren(ctx, propogate)
}

// Bounds returns the minimum bounding box within which the Element's content is
// contained.
func (e *Element) Bounds() types.Rectangle {
	bounds := e.Bounded.Bounds()
	size := e.Sized.Size()
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

	if e.Fixed() {
		fixed := e.FixedPosition()
		bounds.Min.X = fixed.X
		bounds.Min.Y = fixed.Y
	} else {
		offset := e.RelativePosition()
		bounds.Min.X += offset.X
		bounds.Min.Y += offset.Y
		bounds.Max.X += offset.X
		bounds.Max.Y += offset.Y
	}
	return bounds
}

// InnerBounds returns the inner bounding box for the Element, which accounts for
// any border and padding.
func (e *Element) InnerBounds() types.Rectangle {
	outer := e.Bounds()
	border := e.Border()
	if border != nil {
		outer.Min.X++
		outer.Min.Y++
		outer.Max.X--
		outer.Max.Y--
	}

	return e.Padding().Bounds(outer)
}

// Draw implements the uv.Renderable interface
func (e *Element) Draw(screen types.Screen, area types.Rectangle) {
	ctx := context.TODO()
	e.Debug(ctx, "Element([%s]%s).Draw: bounding box %s\n", e.class, e.id, area)
	// determine the overlapping bounding element and clear its cells before
	// rendering the element.
	bb := render.Overlapping(area, e.Bounds())
	render.Clear(screen, bb)

	// If we have a border, draw it around the outer bounding box.
	border := e.Border()
	if border != nil {
		e.Debug(ctx, "Element([%s]%s).Draw: drawing border around %s\n", e.class, e.id, area)
		border.Draw(screen, bb)
	}
}

// Render renders the Element to the given buffer at the specified area.
func (e *Element) Render(
	ctx context.Context,
	screen types.Screen,
) {
	bounds := e.Bounds()

	// Position the root element within the inner bounding box. The root
	// element is responsible for propogating this positioning change to any
	// child elements.
	inner := e.InnerBounds()
	e.Debug(ctx, "Element([%s]%s).Render: outer bounds: %s, inner bounds: %s, children: %d\n", e.class, e.id, bounds, inner, len(e.Children()))
	render := func(ctx context.Context, child types.Node) {
		el, ok := child.(types.Element)
		if !ok {
			// it's not an Element, so do nothing...
			return
		}
		el.Render(ctx, screen)
	}
	e.Draw(screen, inner)
	e.VisitChildren(ctx, render)
}

var _ types.Renderable = (*Element)(nil)
var _ types.Element = (*Element)(nil)
