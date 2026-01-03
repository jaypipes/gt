package core

import (
	"github.com/jaypipes/gt/core/types"
)

// Positioned describes something that has a position on a Screen.
type Positioned struct {
	// position is the fixed or relative position for the Positioned.
	position types.Point
	// fixed is true if the Positioned is using fixed coordinates, false if
	// using relative positioning.
	fixed bool
}

// Position returns the position for the Positioned.
func (p *Positioned) FixedPosition() *types.Point {
	if !p.fixed {
		return nil
	}
	return &p.position
}

// FixedX returns the Positioned's fixed X coordinate or -1 if using relative
// positioning.
func (p *Positioned) FixedX() int {
	if !p.fixed {
		return -1
	}
	return p.position.X
}

// FixedY returns the Positioned's fixed Y coordinate or -1 if using relative
// positioning.
func (p *Positioned) FixedY() int {
	if !p.fixed {
		return -1
	}
	return p.position.Y
}

// Fixed returns whether the Positioned is using fixed positioning.
func (p *Positioned) Fixed() bool {
	return p.fixed
}

// SetFixedPosition sets the Positioned's fixed position.
func (p *Positioned) SetFixedPosition(x, y int) {
	p.position = types.Point{X: x, Y: y}
	p.fixed = true
}

// SetFixedX sets the Positioned's fixed X coordinate.
func (p *Positioned) SetFixedX(x int) {
	p.position.X = x
	p.fixed = true
}

// SetFixedY sets the Positioned's fixed Y coordinate.
func (p *Positioned) SetFixedY(y int) {
	p.position.Y = y
	p.fixed = true
}

// SetRelativePosition sets the Positioned's positional offset.
func (p *Positioned) SetRelativePosition(x, y int) {
	p.position = types.Point{X: x, Y: y}
	p.fixed = false
}

// SetRelativeX sets the Positioned's relative X coordinate.
func (p *Positioned) SetRelativeX(x int) {
	p.position.X = x
	p.fixed = false
}

// SetRelativeY sets the Positioned's relative Y coordinate.
func (p *Positioned) SetRelativeY(y int) {
	p.position.Y = y
	p.fixed = false
}

// RelativePosition returns the positional offset for the Positioned.
func (p *Positioned) RelativePosition() *types.Point {
	if p.fixed {
		return nil
	}
	return &p.position
}

// RelativeX returns the Positioned's relative X coordinate or -1 if using
// fixed positioning.
func (p *Positioned) RelativeX() int {
	if p.fixed {
		return -1
	}
	return p.position.X
}

// RelativeY returns the Positioned's relative Y coordinate or -1 if using
// fixed positioning.
func (p *Positioned) RelativeY() int {
	if p.fixed {
		return -1
	}
	return p.position.Y
}

// Relative returns whether the Positioned is using relative positioning.
func (p *Positioned) Relative() bool {
	return !p.fixed
}

var _ types.Positioned = (*Positioned)(nil)
