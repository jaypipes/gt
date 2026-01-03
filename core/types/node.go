package types

// Node describes an element in the tree of elements contained in the Canvas.
type Node interface {
	Bounded
	// ChildAt returns the child element at the supplied zero-based index, or nil
	// if the index is out of bounds.
	ChildAt(int) Node

	// PushChild adds a new child Node to the Node at the end of Node's set of
	// children.
	PushChild(child Node)

	// PopChild removes the last child Node from the Node's children and returns
	// it. Returns nil if Node has no children.
	PopChild() Node

	// SplitHorizontal splits the Node's last child element, adding a new
	// child containing the supplied Node. If the Layout was empty,
	// the supplied Renderable is simply added as the child element.
	SplitHorizontal(Node, Constraint)

	// SplitVertical splits the Node's last child element, adding a new child
	// containing the supplied Node. If the Node had no children, the supplied Node
	// is simply added as the child element.
	SplitVertical(Node, Constraint)
}
