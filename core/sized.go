package core

import (
	"fmt"

	"github.com/jaypipes/gt/core/types"
)

// Sized describes something that has a bounding box.
type Sized struct {
	// size is the size of the Sized.
	size types.Size
	// wconstraint is the optional width size constraint of the Sized.
	wconstraint *types.SizeConstraint
	// wconstraint is the optional height size constraint of the Sized.
	hconstraint *types.SizeConstraint
}

func (s *Sized) String() string {
	if s.size.Empty() {
		return "size=none"
	}
	return fmt.Sprintf("size=%s", s.size)
}

// SetSize sets the Sized's size and marks the Sized as having a fixed width
// and height.
func (s *Sized) SetSize(width, height int) {
	s.size = types.Size{W: width, H: height}
}

// SetWidth sets the Sized's width and marks the Sized as having a fixed width.
func (s *Sized) SetWidth(width int) {
	s.size.H = width
}

// SetHeight sets the Sized's height and marks the Sized as having a fixed
// height.
func (s *Sized) SetHeight(height int) {
	s.size.H = height
}

// SetWidthConstraint sets the Sized's width size constraint.
func (s *Sized) SetWidthConstraint(con types.SizeConstraint) {
	s.wconstraint = &con
}

// SetHeightConstraint sets the Sized's height size constraint.
func (s *Sized) SetHeightConstraint(con types.SizeConstraint) {
	s.hconstraint = &con
}

// Height returns the current height of the Sized.
func (s *Sized) Height() int {
	return s.size.H
}

// Width returns the current width of the Sized.
func (s *Sized) Width() int {
	return s.size.W
}

// Size returns the current width and height for the Sized.
func (s *Sized) Size() types.Size {
	return s.size
}

// FixedWidth returns true if the Sized is using a fixed width.
func (s *Sized) FixedWidth() bool {
	return s.size.W != 0
}

// FixedHeight returns true if the Sized is using a fixed height.
func (s *Sized) FixedHeight() bool {
	return s.size.H != 0
}

// WidthConstraint returns any optional size constraint for the Sized's width.
func (s *Sized) WidthConstraint() *types.SizeConstraint {
	return s.wconstraint
}

// HeightConstraint returns any optional size constraint for the Sized's
// height.
func (s *Sized) HeightConstraint() *types.SizeConstraint {
	return s.hconstraint
}

var _ types.Sized = (*Sized)(nil)
