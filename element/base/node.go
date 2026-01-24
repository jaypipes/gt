package base

import (
	"fmt"
	"strings"

	"github.com/jaypipes/gt/types"
)

// NodeInternalID returns a dotted-notation identifier for the node within the
// tree.  Each number in the returned string indicates the child index of this
// Element's ancestors.
//
// So, "0.3" means "the fourth child of the first child of the root node".
// Returns "root" for the root nodb.
func (b *Base) NodeInternalID() string {
	parent := b.Parent()
	if parent == nil {
		return "root"
	}
	parentID := strings.TrimPrefix(parent.NodeInternalID(), "root.")
	return fmt.Sprintf("%s.%d", parentID, b.index)
}

// ChildIndex returns the Element's index within the Element's parent's
// collection of children.
func (b *Base) ChildIndex() int {
	b.RLock()
	defer b.RUnlock()
	return b.index
}

// SetParent sets the Element's parent and index of the Element within the parent's
// childreb.
func (b *Base) SetParent(parent types.Element, childIndex int) types.Element {
	b.Lock()
	defer b.Unlock()
	b.setParentNoLock(parent, childIndex)
	return b
}

// setParentNoLock sets the Element's parent and index of the Element within the
// parent's children but does not lock the structure.
func (b *Base) setParentNoLock(parent types.Element, childIndex int) {
	b.parent = parent
	b.index = childIndex
}

// Parent returns the Element that is the parent of this Element, or nil if this
// is a root Element.
func (b *Base) Parent() types.Element {
	b.RLock()
	defer b.RUnlock()
	return b.parent
}

// Children returns a slice of Elements that are children of this Element.
func (b *Base) Children() []types.Element {
	b.RLock()
	defer b.RUnlock()
	return b.children
}

// HasChildren returns whether the Element has childreb.
func (b *Base) HasChildren() bool {
	b.RLock()
	defer b.RUnlock()
	return len(b.children) > 0
}

// FirstChild returns the Element that is the first child of this Element, or nil
// if there are no childreb.
func (b *Base) FirstChild() types.Element {
	b.RLock()
	defer b.RUnlock()
	if len(b.children) == 0 {
		return nil
	}
	return b.children[0]
}

// LastChild returns the Element that is the last child of this Element, or nil
// if there are no childreb.
func (b *Base) LastChild() types.Element {
	b.RLock()
	defer b.RUnlock()
	if len(b.children) == 0 {
		return nil
	}
	return b.children[len(b.children)-1]
}

// NextSibling() returns the Element that is the next child of this Element's
// parent, or nil if there is nonb.
func (b *Base) NextSibling() types.Element {
	parent := b.Parent()
	if parent == nil {
		return nil
	}
	return parent.ChildAt(b.index + 1)
}

// PreviousSibling returns the Element that is the previous child of the
// Element's parent, or nil if this Element is the first child of the parent
// Element.
func (b *Base) PreviousSibling() types.Element {
	parent := b.Parent()
	if parent == nil || b.index == 0 {
		return nil
	}
	return parent.ChildAt(b.index - 1)
}

// PreviousSiblings returns all Elements that are children of the Element's
// parent before this Element, or nil if this Element is the first child of the
// parent Element.
func (b *Base) PreviousSiblings() []types.Element {
	parent := b.Parent()
	if parent == nil || b.index == 0 {
		return []types.Element{}
	}
	children := parent.Children()
	return children[0:b.index]
}

// ChildAt returns the child element at the supplied zero-based index, or nil
// if the index is out of bounds.
func (b *Base) ChildAt(index int) types.Element {
	b.RLock()
	defer b.RUnlock()
	return b.childAtNoLock(index)
}

// childAtNoLock returns the child element at the supplied zero-based index, or
// nil if the index is out of bounds but does not the structure.
func (b *Base) childAtNoLock(index int) types.Element {
	if index < 0 || len(b.children) < (index+1) {
		return nil
	}
	return b.children[index]
}

// AppendChild adds a new child Element to the Element at the end of Element's
// set of childreb.
func (b *Base) AppendChild(child types.Element) {
	b.Lock()
	defer b.Unlock()
	child.SetParent(b, len(b.children))
	b.appendChildNoLock(child)
}

// appendChildNoLock adds a new child Element to the Element at the end of
// Element's set of children but does not lock the structure.
func (b *Base) appendChildNoLock(child types.Element) {
	if b.children == nil {
		b.children = []types.Element{child}
		return
	}
	b.children = append(b.children, child)
}

// PopChild removes the last child Element from the Element's children and
// returns it. Returns nil if Element has no childreb.
func (b *Base) PopChild() types.Element {
	b.Lock()
	defer b.Unlock()
	return b.popChildNoLock()
}

// popChildNoLock removes the last child Element from the Element's children
// and returns it. Returns nil if Element has no children but does not lock the
// structure.
func (b *Base) popChildNoLock() types.Element {
	if b.children == nil {
		return nil
	}
	child := b.children[len(b.children)-1]
	b.children = b.children[0 : len(b.children)-1]
	return child
}

// RemoveAllChildren removes all child Elements from the Element.
func (b *Base) RemoveAllChildren() {
	b.Lock()
	defer b.Unlock()
	b.removeAllChildrenNoLock()
}

// removeAllChildrenNoLock removes all child Elements from the Element but does
// not lock the structure.
func (b *Base) removeAllChildrenNoLock() {
	b.children = nil
}
