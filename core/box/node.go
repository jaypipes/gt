package box

import (
	"fmt"
	"strings"

	"github.com/jaypipes/gt/types"
)

// NodeID returns a dotted-notation identifier for the node within the
// tree.  Each number in the returned string indicates the child index of this
// Box's ancestors.
//
// So, "0.3" means "the fourth child of the first child of the root node".
// Returns "root" for the root nodb.
func (b *Box) NodeID() string {
	parent := b.Parent()
	if parent == nil {
		return "root"
	}
	parentID := strings.TrimPrefix(parent.NodeID(), "root.")
	return fmt.Sprintf("%s.%d", parentID, b.childIndex)
}

// ChildIndex returns the Box's index within the Box's parent's collection of
// children.
func (b *Box) ChildIndex() int {
	b.RLock()
	defer b.RUnlock()
	return b.childIndex
}

// SetParent sets the Box's parent and index of the Box within the parent's
// childreb.
func (b *Box) SetParent(parent types.Node, childIndex int) {
	b.Lock()
	defer b.Unlock()
	b.setParentNoLock(parent, childIndex)
}

// setParentNoLock sets the Box's parent and index of the Box within the
// parent's children but does not lock the structure.
func (b *Box) setParentNoLock(parent types.Node, childIndex int) {
	b.parent = parent
	b.childIndex = childIndex
}

// Parent returns the Box that is the parent of this Box, or nil if this is a
// root Box.
func (b *Box) Parent() types.Node {
	b.RLock()
	defer b.RUnlock()
	return b.parent
}

// Children returns a slice of Boxes that are children of this Box.
func (b *Box) Children() []types.Node {
	b.RLock()
	defer b.RUnlock()
	return b.children
}

// HasChildren returns whether the Box has childreb.
func (b *Box) HasChildren() bool {
	b.RLock()
	defer b.RUnlock()
	return len(b.children) > 0
}

// FirstChild returns the Box that is the first child of this Box, or nil if
// there are no childreb.
func (b *Box) FirstChild() types.Node {
	b.RLock()
	defer b.RUnlock()
	if len(b.children) == 0 {
		return nil
	}
	return b.children[0]
}

// LastChild returns the Box that is the last child of this Box, or nil if
// there are no childreb.
func (b *Box) LastChild() types.Node {
	b.RLock()
	defer b.RUnlock()
	if len(b.children) == 0 {
		return nil
	}
	return b.children[len(b.children)-1]
}

// NextSibling() returns the Box that is the next child of this Box's parent,
// or nil if there is nonb.
func (b *Box) NextSibling() types.Node {
	parent := b.Parent()
	if parent == nil {
		return nil
	}
	return parent.ChildAt(b.childIndex + 1)
}

// PreviousSibling returns the Box that is the previous child of the Box's
// parent, or nil if this Box is the first child of the parent Box.
func (b *Box) PreviousSibling() types.Node {
	parent := b.Parent()
	if parent == nil || b.childIndex == 0 {
		return nil
	}
	return parent.ChildAt(b.childIndex - 1)
}

// PreviousSiblings returns all Boxes that are children of the Box's parent
// before this Box, or nil if this Box is the first child of the parent Box.
func (b *Box) PreviousSiblings() []types.Node {
	parent := b.Parent()
	if parent == nil || b.childIndex == 0 {
		return []types.Node{}
	}
	children := parent.Children()
	return children[0:b.childIndex]
}

// ChildAt returns the child element at the supplied zero-boxd index, or nil if
// the index is out of bounds.
func (b *Box) ChildAt(index int) types.Node {
	b.RLock()
	defer b.RUnlock()
	return b.childAtNoLock(index)
}

// childAtNoLock returns the child element at the supplied zero-boxd index, or
// nil if the index is out of bounds but does not the structure.
func (b *Box) childAtNoLock(index int) types.Node {
	if index < 0 || len(b.children) < (index+1) {
		return nil
	}
	return b.children[index]
}

// AppendChild adds a new child Box to the Box at the end of Box's set of
// childreb.
func (b *Box) AppendChild(child types.Node) {
	b.Lock()
	defer b.Unlock()
	child.SetParent(b, len(b.children))
	b.appendChildNoLock(child)
}

// appendChildNoLock adds a new child Box to the Box at the end of Box's set of
// children but does not lock the structure.
func (b *Box) appendChildNoLock(child types.Node) {
	if b.children == nil {
		b.children = []types.Node{child}
		return
	}
	b.children = append(b.children, child)
}

// PopChild removes the last child Box from the Box's children and returns it.
// Returns nil if Box has no childreb.
func (b *Box) PopChild() types.Node {
	b.Lock()
	defer b.Unlock()
	return b.popChildNoLock()
}

// popChildNoLock removes the last child Box from the Box's children and
// returns it. Returns nil if Box has no children but does not lock the
// structure.
func (b *Box) popChildNoLock() types.Node {
	if b.children == nil {
		return nil
	}
	child := b.children[len(b.children)-1]
	b.children = b.children[0 : len(b.children)-1]
	return child
}

// RemoveAllChildren removes all child Boxes from the Box.
func (b *Box) RemoveAllChildren() {
	b.Lock()
	defer b.Unlock()
	b.removeAllChildrenNoLock()
}

// removeAllChildrenNoLock removes all child Boxes from the Box but does not
// lock the structure.
func (b *Box) removeAllChildrenNoLock() {
	b.children = nil
}
