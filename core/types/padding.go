package types

import "fmt"

// Padding contains padding amounts for a box
type Padding struct {
	// T is the top padding, in lines.
	T int
	// B is the bottom padding, in lines.
	B int
	// L is the left padding, in cells.
	L int
	// R is the right padding, in cells.
	R int
}

// AdjustBounds adjusts the supplied bounding box for the various padding
// sizes.
func (p Padding) AdjustBounds(from Rectangle) Rectangle {
	adjusted := from
	adjusted.Min.X += p.L
	adjusted.Min.Y += p.T
	adjusted.Max.X -= p.R
	adjusted.Max.Y -= p.B
	return adjusted
}

// Horizontal returns the total left-right padding
func (p Padding) Horizontal() int {
	return p.L + p.R
}

// Vertical returns the total top-bottom padding
func (p Padding) Vertical() int {
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

// Pad is a convenience function that returns a new Padding.
func Pad(top, bottom, left, right int) Padding {
	return Padding{
		T: top,
		B: bottom,
		L: left,
		R: right,
	}
}

// PadL returns a Padding with the left padding set to the supplied value.
func PadL(left int) Padding {
	return Padding{L: left}
}

// PadR returns a Padding with the right padding set to the supplied value.
func PadR(right int) Padding {
	return Padding{R: right}
}

// PadLR returns a Padding with the left and right padding set to the supplied
// values.
func PadLR(left, right int) Padding {
	return Padding{L: left, R: right}
}

// PadT returns a Padding with the top padding set to the supplied value.
func PadT(top int) Padding {
	return Padding{T: top}
}

// PadB returns a Padding with the bottom padding set to the supplied value.
func PadB(bottom int) Padding {
	return Padding{B: bottom}
}

// PadTB returns a Padding with the top and bottom padding set to the supplied
// values.
func PadTB(top, bottom int) Padding {
	return Padding{T: top, B: bottom}
}
