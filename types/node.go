package types

// Node is an interface that represents something in a single-rooted tree of Nodes. It can
// optionally have a parent Node and child Nodes.
type Node interface {
	// NodeID returns a dotted-notation identifier for the node within
	// the tree. Each number in the returned string indicates the child index
	// of this Node's ancestors.
	//
	// So, "0.3" means "the fourth child of the first child of the root node".
	// Returns "root" for the root node.
	NodeID() string
	// ChildIndex returns the Node's index within the Node's parent's
	// collection of children.
	ChildIndex() int
	// SetParent sets the Node's parent and index of the Node within the
	// parent's children.
	SetParent(Node, int)
	// Parent returns the Node that is the parent of this Node, or nil if this
	// is a root Node.
	Parent() Node
	// AppendChild adds a new child Node to the Node at the end of Node's set
	// of children.
	AppendChild(Node)
	// PopChild removes the last child Node from the Node's children and
	// returns it. Returns nil if Node has no children.
	PopChild() Node
	// RemoveAllChildren removes any children from this Node.
	RemoveAllChildren()
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
	// ChildAt returns the child element at the supplied zero-based index, or
	// nil if the index is out of bounds.
	ChildAt(int) Node
	// NextSibling() returns the Node that is the next child of this Node's
	// parent, or nil if there is none.
	NextSibling() Node
	// PreviousSibling returns the Node that is the previous child of the
	// Node's parent, or nil if this Node is the first child of the parent
	// Node.
	PreviousSibling() Node
	// PreviousSiblings returns all Nodes that are children of the Node's
	// parent before this Node, or an empty slice of Nodes if this Node is the
	// first child of the parent Node.
	PreviousSiblings() []Node
}
