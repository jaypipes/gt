package types

import "fmt"

// Plottable represents a single node in gt's Document Object Model (DOM).
//
// The DOM is a single-rooted tree structure, and as such each Plottable in the
// tree can have zero or more child Plottables and only the root node in the
// tree has no parent Plottable.
//
// Every Plottable knows how to calculate its width and height, how to
// calculate its inner and outer bounding boxes, and how to style itself with a
// border, padding, foreground and background color, etc.
//
// Plottable implements Plot which means that every Plottable can draw itself
// onto a [uv.Screen]. That said, structures that implement Plottable generally
// will embed [core.box.Box] which has basic implementations of most of the
// Plottable interface's methods.
type Plottable interface {
	fmt.Stringer
	Bounded

	// SetBounds sets the Plottable's outer bounding box.
	SetBounds(Rectangle)
	// SetAbsolutePosition sets the Plottable's outer bounding box's top-left
	// coordinates and marks the Plottable as using absolute positioning.
	SetAbsolutePosition(Point)
	// HasAbsolutePosition returns true if the Plottable used absolute positioning.
	HasAbsolutePosition() bool
	// TL returns the Plottable's outer bounding box's top-left coordinates.
	TL() Point
	// TR returns the Plottable's outer bounding box's top-right coordinates.
	TR() Point
	// MinY returns the Min Y (top) of the Plottable's outer bounding box.
	MinY() int
	// MaxY returns the Max Y (bottom) of the Plottable's outer bounding box.
	MaxY() int
	// InnerBounds returns the inner bounding box for the Plottable, which is the
	// outer bounding box adjusted for any border and padding.
	InnerBounds() Rectangle
	// ContainsPoint returns true if the supplied Point lies inside the
	// Plottable's outer bounding box.
	ContainsPoint(Point) bool

	// SetSize constrains the size of the Plottable.
	SetSize(SizeConstraint)
	// Size returns the width and height of the Plottable.
	Size() Size
	// SetWidth constrains the width of the Plottable.
	SetWidth(DimensionConstraint)
	// Width returns the width of the Plottable's outer bounding box.
	Width() Dimension
	// HasFixedWidth returns true if the Plottable has a fixed width.
	HasFixedWidth() bool
	// HasPercentWidth returns true if the Plottable has a percent width
	// cosntraint.
	HasPercentWidth() bool
	// FixedWidth returns the Plottable's fixed width. If the Plottable does not
	// have a fixed width constraint, returns 0.
	FixedWidth() Dimension
	// PercentWidth returns the Plottable's percent width. If the Plottable does
	// not have a percent width constraint, returns 0.
	PercentWidth() Dimension
	// WidthConstraint returns any DimensionConstraint set for the Plottable's
	// width.
	WidthConstraint() DimensionConstraint
	// SetMinWidth sets the minimum width of the Plottable.
	SetMinWidth(Dimension)
	// MinWidth returns the Plottable's minimum width.
	MinWidth() Dimension
	// SetHeight constrains the height of the Plottable.
	SetHeight(DimensionConstraint)
	// HasFixedHeight returns true if the Plottable has a fixed height.
	HasFixedHeight() bool
	// HasPercentHeight returns true if the Plottable has a percent height
	// constraint.
	HasPercentHeight() bool
	// Height returns the height of the Plottable's outer bounding box.
	Height() Dimension
	// FixedHeight returns the Plottable's fixed height. If the Plottable does not
	// have a fixed height constraint, returns 0.
	FixedHeight() Dimension
	// PercentHeight returns the Plottable's percent height. If the Plottable does
	// not have a percent height constraint, returns 0.
	PercentHeight() Dimension
	// HeightConstraint returns any DimensionConstraint set for the Plottable's
	// height.
	HeightConstraint() DimensionConstraint
	// SetMinHeight sets the minimum height of the Plottable.
	SetMinHeight(Dimension)
	// MinHeight returns the Plottable's minimum height.
	MinHeight() Dimension

	// SetDisplayMode sets the display mode of the Displayed
	SetDisplay(Display)
	// DisplayMode returns the display mode of the Displayed
	Display() Display

	// SetAlignment sets the Aligneds' Alignment
	SetAlignment(Alignment)
	// Alignment returns the Aligned's Alignment
	Alignment() Alignment

	// SetWhitespace sets the Plottable's whitespace mode
	SetWhitespace(Whitespace)
	// Whitespace returns the Plottable's whitespace mode
	Whitespace() Whitespace

	// SetPadding sets the Padded's padding.
	SetPadding(Padding)
	// Padding returns the padding for the Padded.
	Padding() Padding

	// SetBounds sets the Plottable's bounding box.
	SetBorder(Border)
	// Bounds returns the bounding box for the Plottable.
	Border() *Border
	// SetBorderForegroundColor sets the Plottable's border foreground color
	// (i.e the color of the border's cells underlying grapheme).
	SetBorderForegroundColor(Color)
	// BorderForegroundColor returns the Plottable's border foreground color.
	BorderForegroundColor() Color
	// SetBorderBackgroundColor sets the Plottable's border background color
	// (i.e the background color of the border's cells.
	SetBorderBackgroundColor(Color)
	// BorderBackgroundColor returns the Plottable's border background color.
	BorderBackgroundColor() Color

	// HorizontalSpace returns the number of cells consumed by the element's
	// left-right padding and border.
	HorizontalSpace() Dimension
	// VerticalSpace returns the number of lines consumed by the element's
	// top-bottom padding and border
	VerticalSpace() Dimension
}
