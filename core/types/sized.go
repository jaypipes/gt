package types

type Size struct {
	// W is the width of something in cells.
	W int
	// H is the height of something in lines.
	H int
}

// Empty returns whether there is no width or height in the Size.
func (s Size) Empty() bool {
	return s.W == 0 && s.H == 0
}

// Sized describes something with a width (in cells) and height (in lines).
type Sized interface {
	// SetSize sets the Sized's size.
	SetSize(int, int)
	// SetWidth sets the Sized's width in cells.
	SetWidth(int)
	// SetHeight sets the Sized's height in lines.
	SetHeight(int)
	// Height returns the height of the Sized.
	Height() int
	// Width returns the width of the Sized.
	Width() int
	// Bounds returns the width and height of the Sized.
	Size() Size
}
