package types

// Positioned describes something that has a position on a Screen.
type Positioned interface {
	// SetFixedPosition sets the Positioned's fixed position.
	SetFixedPosition(int, int)
	// SetX sets the Positioned's fixed X coordinate.
	SetFixedX(int)
	// SetX sets the Positioned's fixed X coordinate.
	SetFixedY(int)
	// FixedPosition returns the fixed position for the Positioned or nil if
	// the Positioned is not fixed.
	FixedPosition() *Point
	// FixedX returns the Positioned's fixed X coordinate or -1 if using
	// relative positioning.
	FixedX() int
	// FixedY returns the Positioned's fixed Y coordinate or -1 if using
	// relative positioning.
	FixedY() int
	// Fixed returns true if the Positioned used fixed positioning.
	Fixed() bool

	// SetRelativePosition sets the Positioned's positional offset.
	SetRelativePosition(int, int)
	// SetX sets the Positioned's relative X coordinate.
	SetRelativeX(int)
	// SetX sets the Positioned's relative X coordinate.
	SetRelativeY(int)
	// RelativePosition returns the relative positional offset for the
	// Positioned or nil if the Positioned is fixed.
	RelativePosition() *Point
	// RelativeX returns the Positioned's relative X coordinate or -1 if using
	// fixed positioning.
	RelativeX() int
	// RelativeY returns the Positioned's relative Y coordinate or -1 if using
	// fixed positioning.
	RelativeY() int
	// Relative returns true if the Positioned used relative positioning.
	Relative() bool
}
