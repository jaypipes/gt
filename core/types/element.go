package types

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"
)

// Element is a specialized type of Node that can be sized and styled.
type Element interface {
	uv.Drawable
	Plotted
	// SetID sets the Element's unique identifier.
	SetID(string)
	// ID returns the Element's unique identifier.
	ID() string
	// SetClass sets the Element's type/class.
	SetClass(string)
	// Class returns the Element's type/class, e.g. "gt.label" or "gt.canvas"
	Class() string
	// Tag returns a string with the Element's type/class and ID
	Tag() string
	// Parent returns the Element that is the parent of this Element, or nil if
	// this is a root Element.
	Parent() Element
	// SetParent sets the Element's parent and index of the Element within the
	// parent's children.
	SetParent(Element, int)
	// Children returns a slice of Elements that are children of this Element.
	Children() []Element
	// HasChildren returns whether the Element has children.
	HasChildren() bool
	// FirstChild returns the Element that is the first child of this Element,
	// or nil if there are no children.
	FirstChild() Element
	// LastChild returns the Element that is the last child of this Element, or
	// nil if there are no children.
	LastChild() Element
	// ChildAt returns the child element at the supplied zero-based index, or
	// nil if the index is out of bounds.
	ChildAt(int) Element
	// NextSibling() returns the Element that is the next child of this
	// Element's parent, or nil if there is none.
	NextSibling() Element
	// PreviousSibling returns the Element that is the previous child of the
	// Element's parent, or nil if this Element is the first child of the
	// parent Element.
	PreviousSibling() Element
	// PushChild adds a new child Element to the Element at the end of
	// Element's set of children.
	PushChild(Element)
	// PopChild removes the last child Element from the Element's children and
	// returns it. Returns nil if Element has no children.
	PopChild() Element
	// RemoveAllChildren removes any children from this Element.
	RemoveAllChildren()
	// VisitChildren executes a callback function against each child Element.
	VisitChildren(context.Context, func(context.Context, Element))
}
