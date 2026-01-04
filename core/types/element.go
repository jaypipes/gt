package types

// Element is a specialized type of Node that can be sized and styled.
type Element interface {
	Renderable
	Node
	// SetID sets the Element's unique identifier.
	SetID(string)
	// ID returns the Element's unique identifier.
	ID() string
	// SetClass sets the Element's type/class.
	SetClass(string)
	// Class returns the Element's type/class, e.g. "gt.label" or "gt.canvas"
	Class() string
}
