package node

import (
	"context"
	"log/slog"
	"sync"

	"github.com/samber/lo"

	"github.com/jaypipes/gt/core"
	gtctx "github.com/jaypipes/gt/core/context"
	"github.com/jaypipes/gt/core/types"
)

// New returns a new instance of a Node
func New(ctx context.Context) *Node {
	log := gtctx.Logger(ctx)
	n := &Node{mu: new(sync.RWMutex)}
	if log != nil {
		n.SetLogger(log)
	}
	return n
}

// Node describes an element in the tree of elements contained in the Canvas.
type Node struct {
	core.Logged
	mu *sync.RWMutex
	// index is the index of this Node in the parent's children.
	index  int
	parent types.Node
	// children is the collection of Nodes that are the direct children of this
	// Node, if any.
	children []types.Node
}

// Logger returns the Node's Logger. If the Node's Logger is nil, fetches the
// parent's Logger if present.
func (n *Node) Logger() *slog.Logger {
	log := n.Logged.Logger()
	if log != nil {
		return log
	}
	if n.parent != nil {
		return n.parent.Logger()
	}
	return nil
}

// Parent returns the Node that is the parent of this Node, or nil if this
// is a root Node.
func (n *Node) Parent() types.Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.parent
}

// SetParent sets the Node's parent and index of the Node within the parent's
// children.
func (n *Node) SetParent(parent types.Node, childIndex int) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.setParentNoLock(parent, childIndex)
}

// setParentNoLock sets the Node's parent and index of the Node within the
// parent's children but does not lock the structure.
func (n *Node) setParentNoLock(parent types.Node, childIndex int) {
	n.parent = parent
	n.index = childIndex
}

// Children returns a slice of Nodes that are children of this Node.
func (n *Node) Children() []types.Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return n.children
}

// HasChildren returns whether the Node has children.
func (n *Node) HasChildren() bool {
	n.mu.RLock()
	defer n.mu.RUnlock()
	return len(n.children) > 0
}

// FirstChild returns the Node that is the first child of this Node, or nil
// if there are no children.
func (n *Node) FirstChild() types.Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	if len(n.children) == 0 {
		return nil
	}
	return n.children[0]
}

// LastChild returns the Node that is the last child of this Node, or nil
// if there are no children.
func (n *Node) LastChild() types.Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	if len(n.children) == 0 {
		return nil
	}
	return n.children[len(n.children)-1]
}

// NextSibling() returns the Node that is the next child of this Node's
// parent, or nil if there is none.
func (n *Node) NextSibling() types.Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	if n.parent == nil {
		return nil
	}
	return n.parent.ChildAt(n.index + 1)
}

// PreviousSibling returns the Node that is the previous child of the
// Node's parent, or nil if this Node is the first child of the parent
// Node.
func (n *Node) PreviousSibling() types.Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	if n.parent == nil {
		return nil
	}
	return n.parent.ChildAt(n.index - 1)
}

// ChildAt returns the child element at the supplied zero-based index, or nil
// if the index is out of bounds.
func (n *Node) ChildAt(index int) types.Node {
	n.mu.RLock()
	defer n.mu.RUnlock()
	if len(n.children) > (index + 1) {
		return nil
	}
	return n.children[index]
}

// PushChild adds a new child Node to the Node at the end of Node's set of
// children.
func (n *Node) PushChild(child types.Node) {
	n.mu.Lock()
	defer n.mu.Unlock()
	child.SetParent(n, len(n.children))
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
	n.mu.Lock()
	defer n.mu.Unlock()
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

// RemoveAllChildren removes all child Nodes from the Node.
func (n *Node) RemoveAllChildren() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.removeAllChildrenNoLock()
}

// removeAllChildrenNoLock removes all child Nodes from the Node but does not lock
// the structure.
func (n *Node) removeAllChildrenNoLock() {
	n.children = nil
}

// VisitChildren executes a callback function against each child Node.
func (n *Node) VisitChildren(
	ctx context.Context,
	fn func(context.Context, types.Node),
) {
	lo.ForEach(n.children, func(child types.Node, _ int) {
		fn(ctx, child)
		child.VisitChildren(ctx, fn)
	})
}

var _ types.Node = (*Node)(nil)
