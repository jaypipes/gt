package core

import (
	"sync"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/types"
)

// Node describes an element in the tree of elements contained in the Canvas.
type Node struct {
	sync.RWMutex
	Bounded
	// parent is a pointer to a Node that is parent of this Node, if any.
	parent *Node
	// children is the collection of Nodes that are the direct children of this
	// Node, if any.
	children []types.Node
}

// ChildAt returns the child element at the supplied zero-based index, or nil
// if the index is out of bounds.
func (n *Node) ChildAt(index int) types.Node {
	n.RLock()
	defer n.RUnlock()
	if len(n.children) > (index + 1) {
		return nil
	}
	return n.children[index]
}

// PushChild adds a new child Node to the Node at the end of Node's set of
// children.
func (n *Node) PushChild(child types.Node) {
	n.Lock()
	defer n.Unlock()
	n.pushChildNoLock(child)
}

// pushChildNoLock adds a new child Node to the Node at the end of Node's set
// of children but does not lock the structure.
func (n *Node) pushChildNoLock(child types.Node) {
	if n.children == nil {
		n.children = []types.Node{child}
		return
	}
	n.children = append(n.children, child)
}

// PopChild removes the last child Node from the Node's children and returns
// it. Returns nil if Node has no children.
func (n *Node) PopChild() types.Node {
	n.Lock()
	defer n.Unlock()
	return n.popChildNoLock()
}

// popChildNoLock removes the last child Node from the Node's children and returns
// it. Returns nil if Node has no children but does not lock the structure.
func (n *Node) popChildNoLock() types.Node {
	if n.children == nil {
		return nil
	}
	child := n.children[len(n.children)-1]
	n.children = n.children[0 : len(n.children)-1]
	return child
}

// SplitHorizontal splits the Node's last child element, adding a new
// child containing the supplied Node. If the Layout was empty,
// the supplied Renderable is simply added as the child element.
func (n *Node) SplitHorizontal(child types.Node, con types.Constraint) {
	n.Lock()
	defer n.Unlock()
	lastChild := n.popChildNoLock()
	if lastChild == nil {
		n.pushChildNoLock(child)
		return
	}
	left, right := uv.SplitHorizontal(lastChild.Bounds(), con)
	lastChild.SetBounds(left)
	n.pushChildNoLock(lastChild)
	child.SetBounds(right)
	n.pushChildNoLock(child)
}

// SplitVertical splits the Node's last child element, adding a new child
// containing the supplied Node. If the Node had no children, the supplied Node
// is simply added as the child element.
func (n *Node) SplitVertical(child types.Node, con types.Constraint) {
	n.Lock()
	defer n.Unlock()
	lastChild := n.popChildNoLock()
	if lastChild == nil {
		n.pushChildNoLock(child)
		return
	}
	top, bottom := uv.SplitVertical(lastChild.Bounds(), con)
	lastChild.SetBounds(top)
	n.pushChildNoLock(lastChild)
	child.SetBounds(bottom)
	n.pushChildNoLock(child)
}
