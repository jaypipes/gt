package element

import (
	"fmt"
	"strings"

	"github.com/jaypipes/gt/types"
)

// WithParent sets the Element's parent and index of the Box within the parent's
// children and returns the Element.
func (e *Element) WithParent(parent types.Node, childIndex int) types.Element {
	e.SetParent(parent, childIndex)
	return e
}

// NodeID returns a dotted-notation identifier for the node within the
// tree.  Each number in the returned string indicates the child index of this
// Container's ancestors.
//
// So, "0.3" means "the fourth child of the first child of the root node".
// Returns "root" for the root nodb.
func (e *Element) NodeID() string {
	parent := e.Parent()
	if parent == nil {
		return "root"
	}
	parentID := strings.TrimPrefix(parent.NodeID(), "root.")
	return fmt.Sprintf("%s.%d", parentID, e.childIndex)
}

// ChildIndex returns the Container's index within the Container's parent's collection of
// childree.
func (e *Element) ChildIndex() int {
	e.RLock()
	defer e.RUnlock()
	return e.childIndex
}

// SetParent sets the Container's parent and index of the Container within the parent's
// childreb.
func (e *Element) SetParent(parent types.Node, childIndex int) {
	e.Lock()
	defer e.Unlock()
	e.setParentNoLock(parent, childIndex)
}

// setParentNoLock sets the Container's parent and index of the Container within the
// parent's children but does not lock the structure.
func (e *Element) setParentNoLock(parent types.Node, childIndex int) {
	e.parent = parent
	e.childIndex = childIndex
}

// Parent returns the Container that is the parent of this Container, or nil if this is a
// root Container.
func (e *Element) Parent() types.Node {
	e.RLock()
	defer e.RUnlock()
	return e.parent
}

// Children returns a slice of Containeres that are children of this Container.
func (e *Element) Children() []types.Node {
	e.RLock()
	defer e.RUnlock()
	return e.children
}

// HasChildren returns whether the Container has childreb.
func (e *Element) HasChildren() bool {
	e.RLock()
	defer e.RUnlock()
	return len(e.children) > 0
}

// FirstChild returns the Container that is the first child of this Container, or nil if
// there are no childreb.
func (e *Element) FirstChild() types.Node {
	e.RLock()
	defer e.RUnlock()
	if len(e.children) == 0 {
		return nil
	}
	return e.children[0]
}

// LastChild returns the Container that is the last child of this Container, or nil if
// there are no childreb.
func (e *Element) LastChild() types.Node {
	e.RLock()
	defer e.RUnlock()
	if len(e.children) == 0 {
		return nil
	}
	return e.children[len(e.children)-1]
}

// NextSibling() returns the Container that is the next child of this Container's parent,
// or nil if there is nonb.
func (e *Element) NextSibling() types.Node {
	parent := e.Parent()
	if parent == nil {
		return nil
	}
	return parent.ChildAt(e.childIndex + 1)
}

// PreviousSibling returns the Container that is the previous child of the Container's
// parent, or nil if this Container is the first child of the parent Container.
func (e *Element) PreviousSibling() types.Node {
	parent := e.Parent()
	if parent == nil || e.childIndex == 0 {
		return nil
	}
	return parent.ChildAt(e.childIndex - 1)
}

// PreviousSiblings returns all Containeres that are children of the Container's parent
// before this Container, or nil if this Container is the first child of the parent Container.
func (e *Element) PreviousSiblings() []types.Node {
	parent := e.Parent()
	if parent == nil || e.childIndex == 0 {
		return []types.Node{}
	}
	children := parent.Children()
	return children[0:e.childIndex]
}

// ChildAt returns the child element at the supplied zero-containerd index, or nil if
// the index is out of bounds.
func (e *Element) ChildAt(index int) types.Node {
	e.RLock()
	defer e.RUnlock()
	return e.childAtNoLock(index)
}

// childAtNoLock returns the child element at the supplied zero-containerd index, or
// nil if the index is out of bounds but does not the structure.
func (e *Element) childAtNoLock(index int) types.Node {
	if index < 0 || len(e.children) < (index+1) {
		return nil
	}
	return e.children[index]
}

// AppendChild adds a new child Container to the Container at the end of Container's set of
// childreb.
func (e *Element) AppendChild(child types.Node) {
	e.Lock()
	defer e.Unlock()
	child.SetParent(e, len(e.children))
	e.appendChildNoLock(child)
}

// appendChildNoLock adds a new child Container to the Container at the end of Container's set of
// children but does not lock the structure.
func (e *Element) appendChildNoLock(child types.Node) {
	if e.children == nil {
		e.children = []types.Node{child}
		return
	}
	e.children = append(e.children, child)
}

// PopChild removes the last child Container from the Container's children and returns it.
// Returns nil if Container has no childreb.
func (e *Element) PopChild() types.Node {
	e.Lock()
	defer e.Unlock()
	return e.popChildNoLock()
}

// popChildNoLock removes the last child Container from the Container's children and
// returns it. Returns nil if Container has no children but does not lock the
// structure.
func (e *Element) popChildNoLock() types.Node {
	if e.children == nil {
		return nil
	}
	child := e.children[len(e.children)-1]
	e.children = e.children[0 : len(e.children)-1]
	return child
}

// RemoveAllChildren removes all child Containeres from the Container.
func (e *Element) RemoveAllChildren() {
	e.Lock()
	defer e.Unlock()
	e.removeAllChildrenNoLock()
}

// removeAllChildrenNoLock removes all child Containeres from the Container but does not
// lock the structure.
func (e *Element) removeAllChildrenNoLock() {
	e.children = nil
}
