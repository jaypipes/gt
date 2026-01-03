package core

import (
	"github.com/jaypipes/gt/core/types"
)

// Sized describes something that has a bounding box.
type Sized struct {
	// size is the size of the Sized.
	size types.Size
}

// SetSize sets the Sized's size.
func (s *Sized) SetSize(width, height int) {
	s.size = types.Size{W: width, H: height}
}

// SetWidth sets the Sized's width.
func (s *Sized) SetWidth(width int) {
	s.size.H = width
}

// SetHeight sets the Sized's height.
func (s *Sized) SetHeight(height int) {
	s.size.H = height
}

// Height returns the height of the Sized.
func (s *Sized) Height() int {
	return s.size.H
}

// Width returns the width of the Sized.
func (s *Sized) Width() int {
	return s.size.W
}

// Size returns the bounding box for the Sized
func (s *Sized) Size() types.Size {
	return s.size
}

var _ types.Sized = (*Sized)(nil)
