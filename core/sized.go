package core

import (
	"github.com/jaypipes/gt/core/types"
)

// Sized describes something that can have its width and height set and
// calculated/fetched.
type Sized struct {
	// minWidth is the minimum width of the Sized.
	minWidth types.Dimension
	// minHeight is the minimum height of the Sized.
	minHeight types.Dimension
	// widthConstraint is the constraint put on the width dimension
	widthConstraint types.DimensionConstraint
	// heightConstraint is the constraint put on the height dimension
	heightConstraint types.DimensionConstraint
}

// SetSize constrains the size of the Plotted's inner bounding box.
func (s *Sized) SetSize(constraint types.SizeConstraint) {
	wc := constraint.Width()
	if wc != nil {
		s.widthConstraint = wc
	}
	hc := constraint.Height()
	if hc != nil {
		s.heightConstraint = hc
	}
}

// SetWidth constrains the width of the Sized.
func (s *Sized) SetWidth(constraint types.DimensionConstraint) {
	s.widthConstraint = constraint
}

// SetMinWidth sets the minimum width of the Sized.
func (s *Sized) SetMinWidth(w types.Dimension) {
	s.minWidth = w
}

// SetHeight constrains the height of the Sized.
func (s *Sized) SetHeight(constraint types.DimensionConstraint) {
	s.heightConstraint = constraint
}

// SetMinHeight sets the minimum height of the Sized.
func (s *Sized) SetMinHeight(h types.Dimension) {
	s.minHeight = h
}

// Width returns the Sized's width. This should always be overridden by
// Element subclasses.
func (s *Sized) Width() types.Dimension {
	return types.Dimension(0)
}

// FixedWidth returns the Sized's fixed width. If the Sized does not have a
// fixed width constraint, returns 0.
func (s *Sized) FixedWidth() types.Dimension {
	if !s.HasFixedWidth() {
		return types.Dimension(0)
	}
	return types.Dimension(s.widthConstraint.(FixedConstraint))
}

// MinWidth returns the Sized's minimum width.
func (s *Sized) MinWidth() types.Dimension {
	return s.minWidth
}

// Height returns the Sized's height. This should always be overridden by
// Element subclasses.
func (s *Sized) Height() types.Dimension {
	return types.Dimension(0)
}

// FixedHeight returns the Sized's fixed height. If the Sized does not have a
// fixed height constraint, returns 0.
func (s *Sized) FixedHeight() types.Dimension {
	if !s.HasFixedHeight() {
		return types.Dimension(0)
	}
	return types.Dimension(s.heightConstraint.(FixedConstraint))
}

// MinHeight returns the Sized's minimum height.
func (s *Sized) MinHeight() types.Dimension {
	return s.minHeight
}

// HasFixedWidth returns true if the Plotted's inner bounding box has a fixed
// width.
func (s *Sized) HasFixedWidth() bool {
	_, ok := s.widthConstraint.(FixedConstraint)
	return ok
}

// HasFixedHeight returns true if the Plotted's inner bounding box has a fixed
// height.
func (s *Sized) HasFixedHeight() bool {
	_, ok := s.heightConstraint.(FixedConstraint)
	return ok
}

// WidthConstraint returns any optional size constraint for the Sized's width.
// Returns nil when there is no width constraint.
func (s *Sized) WidthConstraint() types.DimensionConstraint {
	return s.widthConstraint
}

// HeightConstraint returns any optional size constraint for the Sized's
// height. Returns nil when there is no height constraint.
func (s *Sized) HeightConstraint() types.DimensionConstraint {
	return s.heightConstraint
}
