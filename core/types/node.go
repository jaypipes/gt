package types

import "context"

// Node describes an element in the tree of elements contained in the Canvas.
type Node interface {
	Logged
	// Parent returns the Node that is the parent of this Node, or nil if this
	// is a root Node.
	Parent() Node
	// SetParent sets the Node's parent and index of the Node within the
	// parent's children.
	SetParent(Node, int)
	// Children returns a slice of Nodes that are children of this Node.
	Children() []Node
	// HasChildren returns whether the Node has children.
	HasChildren() bool
	// FirstChild returns the Node that is the first child of this Node, or nil
	// if there are no children.
	FirstChild() Node
	// LastChild returns the Node that is the last child of this Node, or nil
	// if there are no children.
	LastChild() Node
	// ChildAt returns the child element at the supplied zero-based index, or nil
	// if the index is out of bounds.
	ChildAt(int) Node
	// NextSibling() returns the Node that is the next child of this Node's
	// parent, or nil if there is none.
	NextSibling() Node
	// PreviousSibling returns the Node that is the previous child of the
	// Node's parent, or nil if this Node is the first child of the parent
	// Node.
	PreviousSibling() Node
	// PushChild adds a new child Node to the Node at the end of Node's set of
	// children.
	PushChild(Node)
	// PopChild removes the last child Node from the Node's children and returns
	// it. Returns nil if Node has no children.
	PopChild() Node
	// RemoveAllChildren removes any children from this Node.
	RemoveAllChildren()
	// VisitChildren executes a callback function against each child Node.
	VisitChildren(context.Context, func(context.Context, Node))
}
