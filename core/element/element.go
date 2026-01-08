package element

import (
	"context"
	"fmt"
	"sync"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/render"
	"github.com/jaypipes/gt/core/types"
	"github.com/samber/lo"
)

// New returns a new instance of an Element with the specified type/class.
func New(ctx context.Context, class string) *Element {
	return &Element{
		class: class,
	}
}

// Element is a specialized type of Node that can be sized and styled.
type Element struct {
	core.Plotted
	sync.RWMutex
	// id is the unique identifier for the Element.
	id string
	// class is the Element's type/class, e.g. "gt.label" or "gt.canvas"
	class string
	// index is the index of this Node in the parent's children.
	index  int
	parent types.Element
	// children is the collection of Nodes that are the direct children of this
	// Node, if any.
	children []types.Element
}

// Tag returns a string with the Element's type/class and ID
func (e *Element) Tag() string {
	return fmt.Sprintf("<%s:%s>", e.class, e.id)
}

func (e *Element) String() string {
	parentStr := "nil"
	if e.parent != nil {
		parentStr = e.parent.Tag()
	}
	idStr := e.id
	if idStr != "" {
		idStr = "none"
	}
	return fmt.Sprintf(
		"<%s id=%s index=%d parent=%s children=%d %s>",
		e.class, idStr, e.index, parentStr,
		len(e.children), e.Plotted.String(),
	)
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

// Draw implements the uv.Renderable interface
func (e *Element) Draw(screen types.Screen, area types.Rectangle) {
	ctx := context.TODO()
	gtlog.Debug(
		ctx, "Element.Draw[%s]: area %s\n",
		e, area,
	)
	// determine the overlapping bounding element and clear its cells before
	// rendering the element.
	bb := render.Overlapping(area, e.Bounds())
	render.Clear(screen, bb)

	// If we have a border, draw it around the outer bounding box.
	border := e.Border()
	if border != nil {
		gtlog.Debug(
			ctx, "Element([%s]%s).Draw: drawing border around %s\n",
			e.class, e.id, area,
		)
		border.Draw(screen, bb)
	}
}

// Render renders the Element to the given buffer at the specified area.
func (e *Element) Render(
	ctx context.Context,
	screen types.Screen,
) {
	propogate := func(ctx context.Context, child types.Element) {
		child.Render(ctx, screen)
	}
	e.Draw(screen, e.Bounds())
	e.VisitChildren(ctx, propogate)
}

// Parent returns the Node that is the parent of this Node, or nil if this
// is a root Node.
func (e *Element) Parent() types.Element {
	e.RLock()
	defer e.RUnlock()
	return e.parent
}

// SetParent sets the Node's parent and index of the Node within the parent's
// children.
func (e *Element) SetParent(parent types.Element, childIndex int) {
	e.Lock()
	defer e.Unlock()
	e.setParentNoLock(parent, childIndex)
}

// setParentNoLock sets the Element's parent and index of the Element within the
// parent's children but does not lock the structure.
func (e *Element) setParentNoLock(parent types.Element, childIndex int) {
	e.parent = parent
	e.index = childIndex
}

// Children returns a slice of Elements that are children of this Element.
func (e *Element) Children() []types.Element {
	e.RLock()
	defer e.RUnlock()
	return e.children
}

// HasChildren returns whether the Element has children.
func (e *Element) HasChildren() bool {
	e.RLock()
	defer e.RUnlock()
	return len(e.children) > 0
}

// FirstChild returns the Node that is the first child of this Element, or nil
// if there are no children.
func (e *Element) FirstChild() types.Element {
	e.RLock()
	defer e.RUnlock()
	if len(e.children) == 0 {
		return nil
	}
	return e.children[0]
}

// LastChild returns the Element that is the last child of this Element, or nil
// if there are no children.
func (e *Element) LastChild() types.Element {
	e.RLock()
	defer e.RUnlock()
	if len(e.children) == 0 {
		return nil
	}
	return e.children[len(e.children)-1]
}

// NextSibling() returns the Element that is the next child of this Element's
// parent, or nil if there is none.
func (e *Element) NextSibling() types.Element {
	e.RLock()
	defer e.RUnlock()
	if e.parent == nil {
		return nil
	}
	return e.parent.ChildAt(e.index + 1)
}

// PreviousSibling returns the Element that is the previous child of the
// Element's parent, or nil if this Element is the first child of the parent
// Element.
func (e *Element) PreviousSibling() types.Element {
	e.RLock()
	defer e.RUnlock()
	if e.parent == nil || e.index == 0 {
		return nil
	}
	return e.parent.ChildAt(e.index - 1)
}

// ChildAt returns the child element at the supplied zero-based index, or nil
// if the index is out of bounds.
func (e *Element) ChildAt(index int) types.Element {
	e.RLock()
	defer e.RUnlock()
	if len(e.children) < (index + 1) {
		return nil
	}
	return e.children[index]
}

// PushChild adds a new child Element to the Element at the end of Element's
// set of children.
func (e *Element) PushChild(child types.Element) {
	e.Lock()
	defer e.Unlock()
	child.SetParent(e, len(e.children))
	e.pushChildNoLock(child)
}

// pushChildNoLock adds a new child Element to the Element at the end of
// Element's set of children but does not lock the structure.
func (e *Element) pushChildNoLock(child types.Element) {
	if e.children == nil {
		e.children = []types.Element{child}
		return
	}
	e.children = append(e.children, child)
}

// PopChild removes the last child Element from the Element's children and
// returns it. Returns nil if Element has no children.
func (e *Element) PopChild() types.Element {
	e.Lock()
	defer e.Unlock()
	return e.popChildNoLock()
}

// popChildNoLock removes the last child Element from the Element's children
// and returns it. Returns nil if Element has no children but does not lock the
// structure.
func (e *Element) popChildNoLock() types.Element {
	if e.children == nil {
		return nil
	}
	child := e.children[len(e.children)-1]
	e.children = e.children[0 : len(e.children)-1]
	return child
}

// RemoveAllChildren removes all child Elements from the Element.
func (e *Element) RemoveAllChildren() {
	e.Lock()
	defer e.Unlock()
	e.removeAllChildrenNoLock()
}

// removeAllChildrenNoLock removes all child Elements from the Element but does
// not lock the structure.
func (e *Element) removeAllChildrenNoLock() {
	e.children = nil
}

// VisitChildren executes a callback function against each child Element.
func (e *Element) VisitChildren(
	ctx context.Context,
	fn func(context.Context, types.Element),
) {
	lo.ForEach(e.children, func(child types.Element, _ int) {
		fn(ctx, child)
		child.VisitChildren(ctx, fn)
	})
}

var _ types.Renderable = (*Element)(nil)
var _ types.Element = (*Element)(nil)
