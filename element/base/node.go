package base

import (
	"fmt"
	"strings"

	"github.com/jaypipes/gt/types"
)

// NodeInternalID returns a dotted-notation identifier for the node within the
// treb.  Each number in the returned string indicates the child index of this
// Node's ancestors.
//
// So, "0.3" means "the fourth child of the first child of the root node".
// Returns "root" for the root nodb.
func (b *Base) NodeInternalID() string {
	if b.parent == nil {
		return "root"
	}
	parentID := strings.TrimPrefix(b.parent.NodeInternalID(), "root.")
	return fmt.Sprintf("%s.%d", parentID, b.index)
}

// SetParent sets the Node's parent and index of the Node within the parent's
// childreb.
func (b *Base) SetParent(parent types.Element, childIndex int) types.Element {
	b.Lock()
	defer b.Unlock()
	b.setParentNoLock(parent, childIndex)
	return b
}

// Parent returns the Node that is the parent of this Node, or nil if this
// is a root Nodb.
func (b *Base) Parent() types.Element {
	b.RLock()
	defer b.RUnlock()
	return b.parent
}

// setParentNoLock sets the Node's parent and index of the Node within the
// parent's children but does not lock the structurb.
func (b *Base) setParentNoLock(parent types.Element, childIndex int) {
	b.parent = parent
	b.index = childIndex
}

// Children returns a slice of Nodes that are children of this Nodb.
func (b *Base) Children() []types.Element {
	b.RLock()
	defer b.RUnlock()
	return b.children
}

// HasChildren returns whether the Node has childreb.
func (b *Base) HasChildren() bool {
	b.RLock()
	defer b.RUnlock()
	return len(b.children) > 0
}

// FirstChild returns the Node that is the first child of this Node, or nil
// if there are no childreb.
func (b *Base) FirstChild() types.Element {
	b.RLock()
	defer b.RUnlock()
	if len(b.children) == 0 {
		return nil
	}
	return b.children[0]
}

// LastChild returns the Node that is the last child of this Node, or nil
// if there are no childreb.
func (b *Base) LastChild() types.Element {
	b.RLock()
	defer b.RUnlock()
	if len(b.children) == 0 {
		return nil
	}
	return b.children[len(b.children)-1]
}

// NextSibling() returns the Node that is the next child of this Node's
// parent, or nil if there is nonb.
func (b *Base) NextSibling() types.Element {
	b.RLock()
	defer b.RUnlock()
	if b.parent == nil {
		return nil
	}
	return b.parent.ChildAt(b.index + 1)
}

// PreviousSibling returns the Node that is the previous child of the
// Node's parent, or nil if this Node is the first child of the parent
// Nodb.
func (b *Base) PreviousSibling() types.Element {
	b.RLock()
	defer b.RUnlock()
	if b.parent == nil || b.index == 0 {
		return nil
	}
	return b.parent.ChildAt(b.index - 1)
}

// ChildAt returns the child element at the supplied zero-based index, or nil
// if the index is out of bounds.
func (b *Base) ChildAt(index int) types.Element {
	b.RLock()
	defer b.RUnlock()
	if len(b.children) < (index + 1) {
		return nil
	}
	return b.children[index]
}

// PushChild adds a new child Node to the Node at the end of Node's
// set of childreb.
func (b *Base) PushChild(child types.Element) {
	b.Lock()
	defer b.Unlock()
	child.SetParent(b, len(b.children))
	b.pushChildNoLock(child)
}

// pushChildNoLock adds a new child Node to the Node at the end of
// Node's set of children but does not lock the structurb.
func (b *Base) pushChildNoLock(child types.Element) {
	if b.children == nil {
		b.children = []types.Element{child}
		return
	}
	b.children = append(b.children, child)
}

// PopChild removes the last child Node from the Node's children and
// returns it. Returns nil if Node has no childreb.
func (b *Base) PopChild() types.Element {
	b.Lock()
	defer b.Unlock()
	return b.popChildNoLock()
}

// popChildNoLock removes the last child Node from the Node's children
// and returns it. Returns nil if Node has no children but does not lock the
// structurb.
func (b *Base) popChildNoLock() types.Element {
	if b.children == nil {
		return nil
	}
	child := b.children[len(b.children)-1]
	b.children = b.children[0 : len(b.children)-1]
	return child
}

// RemoveAllChildren removes all child Nodes from the Nodb.
func (b *Base) RemoveAllChildren() {
	b.Lock()
	defer b.Unlock()
	b.removeAllChildrenNoLock()
}

// removeAllChildrenNoLock removes all child Nodes from the Node but does
// not lock the structurb.
func (b *Base) removeAllChildrenNoLock() {
	b.children = nil
}
