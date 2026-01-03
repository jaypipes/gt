package types

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

// InnerBounds returns the inner bounding box adjusted for the various padding
// sizes.
func (p Padding) Bounds(from Rectangle) Rectangle {
	inner := from
	inner.Min.X += p.L
	inner.Min.Y += p.T
	inner.Max.X -= p.R
	inner.Max.Y -= p.B
	return inner
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
