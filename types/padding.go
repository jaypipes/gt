package types

import "fmt"

// Padding contains padding amounts for a box
type Padding struct {
	// T is the top padding, in lines.
	T Dimension
	// B is the bottom padding, in lines.
	B Dimension
	// L is the left padding, in cells.
	L Dimension
	// R is the right padding, in cells.
	R Dimension
}

// AdjustBounds adjusts the supplied bounding box for the various padding
// sizes.
func (p Padding) AdjustBounds(from Rectangle) Rectangle {
	adjusted := from
	adjusted.Min.X += int(p.L)
	adjusted.Min.Y += int(p.T)
	adjusted.Max.X -= int(p.R)
	adjusted.Max.Y -= int(p.B)
	return adjusted
}

// Horizontal returns the total number of cells of left-right padding
func (p Padding) HorizontalSpace() Dimension {
	return p.L + p.R
}

// VerticalSpace returns the total number of lines of top-bottom padding.
func (p Padding) VerticalSpace() Dimension {
	return p.T + p.B
}

// Empty returns true if there's no padding
func (p Padding) Empty() bool {
	return p.T == 0 && p.B == 0 && p.L == 0 && p.R == 0
}

// String returns a string representation of the Padding.
func (p Padding) String() string {
	return fmt.Sprintf("t:%d,b:%d,l:%d,r:%d", p.T, p.B, p.L, p.R)
}

// Pad is a convenience function that returns a new Padding containing a
// uniform padding of the supplied value.
func Pad(value int) Padding {
	return Padding{
		T: Dimension(value),
		B: Dimension(value),
		L: Dimension(value),
		R: Dimension(value),
	}
}

// PadTBLR is a convenience function that returns a new Padding containing the
// individual padding values for top, bottom, left and right.
func PadTBLR(top, bottom, left, right int) Padding {
	return Padding{
		T: Dimension(top),
		B: Dimension(bottom),
		L: Dimension(left),
		R: Dimension(right),
	}
}

// PadL returns a Padding with the left padding set to the supplied value.
func PadL(left int) Padding {
	return Padding{L: Dimension(left)}
}

// PadR returns a Padding with the right padding set to the supplied value.
func PadR(right int) Padding {
	return Padding{R: Dimension(right)}
}

// PadLR returns a Padding with the left and right padding set to the supplied
// values.
func PadLR(left, right int) Padding {
	return Padding{L: Dimension(left), R: Dimension(right)}
}

// PadT returns a Padding with the top padding set to the supplied value.
func PadT(top int) Padding {
	return Padding{T: Dimension(top)}
}

// PadB returns a Padding with the bottom padding set to the supplied value.
func PadB(bottom int) Padding {
	return Padding{B: Dimension(bottom)}
}

// PadTB returns a Padding with the top and bottom padding set to the supplied
// values.
func PadTB(top, bottom int) Padding {
	return Padding{T: Dimension(top), B: Dimension(bottom)}
}
