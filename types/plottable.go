package types

import (
	"context"

	uv "github.com/charmbracelet/ultraviolet"
)

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
	uv.Drawable
	Identifiable

	// Plot calculates the anchoring positioning coordinates of a Plottable.
	// It traverses the tree of Plottables and calculates the top left
	// coordinates for each node in the tree.
	Plot(context.Context)

	// Render wraps the [uv.Drawable.Draw] interface method with a context and
	// always calls [uv.Drawable.Draw] with the Plottable's plotted bounds.
	Render(context.Context, Screen)

	// SetBounds sets the Plottable's outer bounding box.
	SetBounds(Rectangle)
	// Bounds returns the Plottable's outer bounding box.
	Bounds() Rectangle
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

	// SetSize constrains the size of the Plottable.
	SetSize(SizeConstraint)
	// Size returns the width and height of the Plottable.
	Size() Size
	// SetWidth constrains the width of the Plottable.
	SetWidth(DimensionConstraint)
	// OuterWidth returns the width of the Plottable's outer bounding box.
	OuterWidth() Dimension
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
	// OuterHeight returns the height of the Plottable's outer bounding box.
	OuterHeight() Dimension
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

	// NodeInternalID returns a dotted-notation identifier for the node within
	// the tree. Each number in the returned string indicates the child index
	// of this Node's ancestors.
	//
	// So, "0.3" means "the fourth child of the first child of the root node".
	// Returns "root" for the root node.
	NodeInternalID() string
	// ChildIndex returns the Plottable's index within the Plottable's parent's
	// collection of children.
	ChildIndex() int
	// SetParent sets the Plottable's parent and index of the Plottable within the
	// parent's children.
	SetParent(Plottable, int)
	// Parent returns the Plottable that is the parent of this Plottable, or nil if this
	// is a root Plottable.
	Parent() Plottable
	// AppendChild adds a new child Plottable to the Plottable at the end of Plottable's set of
	// children.
	AppendChild(Plottable)
	// PopChild removes the last child Plottable from the Plottable's children and returns
	// it. Returns nil if Plottable has no children.
	PopChild() Plottable
	// RemoveAllChildren removes any children from this Node.
	RemoveAllChildren()
	// Children returns a slice of Plottables that are children of this Plottable.
	Children() []Plottable
	// HasChildren returns whether the Plottable has children.
	HasChildren() bool
	// FirstChild returns the Plottable that is the first child of this Plottable, or nil
	// if there are no children.
	FirstChild() Plottable
	// LastChild returns the Plottable that is the last child of this Plottable, or nil
	// if there are no children.
	LastChild() Plottable
	// ChildAt returns the child element at the supplied zero-based index, or nil
	// if the index is out of bounds.
	ChildAt(int) Plottable
	// NextSibling() returns the Plottable that is the next child of this Plottable's
	// parent, or nil if there is none.
	NextSibling() Plottable
	// PreviousSibling returns the Plottable that is the previous child of the
	// Plottable's parent, or nil if this Plottable is the first child of the parent
	// Plottable.
	PreviousSibling() Plottable
	// PreviousSiblings returns all Plottables that are children of the Plottable's
	// parent before this Plottable, or an empty slice of Plottables if this
	// Plottable is the first child of the parent Plottable.
	PreviousSiblings() []Plottable
}
