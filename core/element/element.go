package element

import (
	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/render"
)

type ElementModifier func(*Element)

// WithBounds sets the rectangle boundaries of the Element.
func WithBounds(x, y, width, height int) ElementModifier {
	return func(e *Element) {
		e.bounds = uv.Rect(x, y, width, height)
	}
}

// WithBorder sets a Border style for the Element.
func WithBorder(border uv.Border) ElementModifier {
	return func(e *Element) {
		e.border = &border
	}
}

// New returns a new Element instance.
func New(mods ...ElementModifier) *Element {
	e := &Element{}
	for _, mod := range mods {
		mod(e)
	}
	return e
}

// Element is a [uv.Drawable] that renders content on the screen. All
// renderable elements and components will inherit from Element.
type Element struct {
	// bounds is the outermost bounding rectangle of the Element.
	bounds uv.Rectangle
	// border is the optional Border information for the Element.
	border *uv.Border
	// children is the collection of Drawable child elements within the Element.
	children []uv.Drawable
}

// SetBounds sets the Element's bounding rectangle
func (e *Element) SetBounds(r uv.Rectangle) {
	e.bounds = r
}

// SetRect sets the Element's bounding rectangle
func (e *Element) SetBorder(border uv.Border) {
	e.border = &border
}

// Height returns the height of the Element
func (e *Element) Height() int {
	return e.bounds.Max.Y - e.bounds.Min.Y
}

// Width returns the width of the Element
func (e *Element) Width() int {
	return e.bounds.Max.X - e.bounds.Min.X
}

// Bounds returns the bounding element for the Element
func (e *Element) Bounds() uv.Rectangle {
	return e.bounds
}

// InnerBounds returns the bounding element for available content inside the Element's
// border and padding.
func (e *Element) InnerBounds() uv.Rectangle {
	bb := e.bounds
	if e.border != nil {
		bb.Min.X++
		bb.Min.Y++
		bb.Max.X--
		bb.Max.Y--
	}
	return bb
}

// Children returns the collection of Drawable child elements within the Element.
func (e *Element) Children() []uv.Drawable {
	return e.children
}

// ChildAt returns the child element at the supplied index, or nil if the index
// is out of bounds.
func (e *Element) ChildAt(index int) uv.Drawable {
	if len(e.children) > (index + 1) {
		return nil
	}
	return e.children[index]
}

// PushChild adds a new Drawable to the end of the Element's collection of
// child elements.
func (e *Element) PushChild(c uv.Drawable) {
	e.children = append(e.children, c)
}

// PopChild removes and returns the last child element from the Element's
// collection of child elements. Returns nil if there are no child elements.
func (e *Element) PopChild() uv.Drawable {
	if e.children == nil {
		return nil
	}
	c := e.children[len(e.children)-1]
	e.children = e.children[0 : len(e.children)-1]
	return c
}

// Border returns the Element's border, if any.
func (e *Element) Border() *uv.Border {
	return e.border
}

// Draw renders the Element to the given buffer at the specified area.
func (e *Element) Draw(buf uv.Screen, area uv.Rectangle) {
	// determine the overlapping bounding element and clear its cells before
	// rendering the element.
	bb := render.Overlapping(area, e.bounds)
	render.Clear(buf, bb)

	// If we have a border, draw it around the Element's outer bounding box.
	if e.border != nil {
		e.border.Draw(buf, bb)
	}

	// Draw any child elements within the inner bounding box.
	inner := e.InnerBounds()
	for _, child := range e.children {
		child.Draw(buf, inner)
	}
}

var _ uv.Drawable = (*Element)(nil)
