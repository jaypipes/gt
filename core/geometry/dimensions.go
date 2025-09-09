package geometry

import "golang.org/x/exp/constraints"

// Measure represents whether a dimension measurement is fixed or a
// percentage of a containing boundary.
type Measure int

const (
	// Fixed means the measure is of a fixed size in cells.
	Fixed Measure = iota
	// Flex means the measure is of a percentage size in cells of a containing
	// boundary.
	Flex
)

// Percent is an alias for Flex
const Percent = Flex

type DimensionsOption func(*Dimensions)

// WithFixedWidth sets a Dimensions' fixed width and ensures any flex width is
// zeroed out.
func WithFixedWidth[T constraints.Unsigned](width T) DimensionsOption {
	return func(d *Dimensions) {
		d.fw = uint16(width)
		d.pw = 0.0
	}
}

// WithFlexWidth sets a Dimensions' fixed width and ensures any flex width is
// zeroed out.
func WithFlexWidth[T constraints.Float](width T) DimensionsOption {
	return func(d *Dimensions) {
		d.fw = 0
		d.pw = float32(width)
	}
}

// Dimensions contains information about an elements width and height, which
// can be either fixed or a percentage of the element's containing boundary.
type Dimensions struct {
	fw uint16
	fh uint16
	pw float32
	ph float32
}

// Width returns the Measure of the Dimension's width, its fixed width (will be
// 0 if measure is Flex) and the percentage width (will be 0 if measure is
// Fixed)
func (d Dimensions) Width() (
	Measure, uint16, float32,
) {
	return d.WidthMeasure(), d.fw, d.pw
}

// WidthMeasure returns the Measure of the Dimension's width.
func (d Dimensions) WidthMeasure() Measure {
	if d.fw == 0 {
		return Flex
	}
	return Fixed
}

// Height returns the Measure of the Dimension's height, its fixed height (will be
// 0 if measure is Flex) and the percentage height (will be 0 if measure is
// Fixed)
func (d Dimensions) Height() (
	Measure, uint16, float32,
) {
	return d.HeightMeasure(), d.fh, d.ph
}

// HeightMeasure returns the Measure of the Dimension's height.
func (d Dimensions) HeightMeasure() Measure {
	if d.fw == 0 {
		return Flex
	}
	return Fixed
}
