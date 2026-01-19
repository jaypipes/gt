package types

import "fmt"

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

// String returns a simple string representation of the Size.
func (s Size) String() string {
	if s.Empty() {
		return "empty"
	}
	return fmt.Sprintf("%dx%d", s.W, s.H)
}
