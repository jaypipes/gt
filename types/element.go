package types

import (
	"context"
	"image/color"

	uv "github.com/charmbracelet/ultraviolet"
)

// Element represents a single node in gt's Document Object Model (DOM).
//
// The DOM is a single-rooted tree structure, and as such each Element in the
// tree can have zero or more child Elements and only the root node in the tree
// has no parent Element.
//
// Every Element knows how to calculate its width and height, how to calculate
// its inner and outer bounding boxes, and how to style itself with a border,
// padding, foreground and background color, etc.
//
// Element implements [uv.Drawable] which means that every Element can draw
// itself onto a [uv.Screen]. That said, structures that implement Element
// generally will embed [element.base.Base] which has basic implementations of
// most of the Element interface's methods, one of which is Render, which wraps
// [uv.Drawable.Draw] for the user.
type Element interface {
	uv.Drawable

	// String returns a string representation of the Element.
	String() string

	// SetID sets the Element's unique identifier.
	SetID(string) Element
	// ID returns the Element's unique identifier.
	ID() string
	// SetClass sets the Element's type/class.
	SetClass(string) Element
	// Class returns the Element's type/class, e.g. "gt.label" or "gt.canvas"
	Class() string
	// Tag returns a string with the Element's type/class and ID
	Tag() string

	// SetTextContent sets the Element's raw string contents.
	SetTextContent(string) Element
	// TextContent returns the Element's raw string contents.
	TextContent() string
	// TextContentWidth returns the width of the Element's raw string contents.
	TextContentWidth() Dimension
	// TextContentHeight returns the height of the Element's raw string
	// contents.
	TextContentHeight() Dimension

	// SetBounds sets the Element's outer bounding box.
	SetBounds(Rectangle) Element
	// Bounds returns the Element's outer bounding box.
	Bounds() Rectangle
	// SetAbsolutePosition sets the Element's outer bounding box's top-left
	// coordinates and marks the Element as using absolute positioning.
	SetAbsolutePosition(Point) Element
	// HasAbsolutePosition returns true if the Element used absolute positioning.
	HasAbsolutePosition() bool
	// TL returns the Element's outer bounding box's top-left coordinates.
	TL() Point
	// TR returns the Element's outer bounding box's top-right coordinates.
	TR() Point
	// MinY returns the Min Y (top) of the Element's outer bounding box.
	MinY() int
	// MaxY returns the Max Y (bottom) of the Element's outer bounding box.
	MaxY() int
	// InnerBounds returns the inner bounding box for the Element, which is the
	// outer bounding box adjusted for any border and padding.
	InnerBounds() Rectangle
	// Plot calculates the anchoring positioning coordinates of a supplied element.
	// It traverses the tree of elements rooted at the supplied element and
	// calculates the top left coordinates for the element.
	Plot(context.Context)

	// Render wraps the [uv.Drawable.Draw] interface method with a context and
	// always calls [uv.Drawable.Draw] with the Rendered's plotted bounds.
	Render(context.Context, Screen)

	// SetSize constrains the size of the Element.
	SetSize(SizeConstraint) Element
	// Size returns the width and height of the Element.
	Size() Size
	// SetWidth constrains the width of the Element.
	SetWidth(DimensionConstraint) Element
	// Width returns the Element's width.
	Width() Dimension
	// HasFixedWidth returns true if the Element has a fixed width.
	HasFixedWidth() bool
	// HasPercentWidth returns true if the Element has a percent width
	// cosntraint.
	HasPercentWidth() bool
	// FixedWidth returns the Element's fixed width. If the Element does not
	// have a fixed width constraint, returns 0.
	FixedWidth() Dimension
	// PercentWidth returns the Element's percent width. If the Element does
	// not have a percent width constraint, returns 0.
	PercentWidth() Dimension
	// WidthConstraint returns any DimensionConstraint set for the Element's
	// width.
	WidthConstraint() DimensionConstraint
	// SetMinWidth sets the minimum width of the Element.
	SetMinWidth(Dimension) Element
	// MinWidth returns the Element's minimum width.
	MinWidth() Dimension
	// SetHeight constrains the height of the Element.
	SetHeight(DimensionConstraint) Element
	// HasFixedHeight returns true if the Element has a fixed height.
	HasFixedHeight() bool
	// HasPercentHeight returns true if the Element has a percent height
	// constraint.
	HasPercentHeight() bool
	// Height returns the Element's height.
	Height() Dimension
	// FixedHeight returns the Element's fixed height. If the Element does not
	// have a fixed height constraint, returns 0.
	FixedHeight() Dimension
	// PercentHeight returns the Element's percent height. If the Element does
	// not have a percent height constraint, returns 0.
	PercentHeight() Dimension
	// HeightConstraint returns any DimensionConstraint set for the Element's
	// height.
	HeightConstraint() DimensionConstraint
	// SetMinHeight sets the minimum height of the Element.
	SetMinHeight(Dimension) Element
	// MinHeight returns the Element's minimum height.
	MinHeight() Dimension

	// SetDisplayMode sets the display mode of the Displayed
	SetDisplay(Display) Element
	// DisplayMode returns the display mode of the Displayed
	Display() Display

	// SetAlignment sets the Aligneds' Alignment
	SetAlignment(Alignment) Element
	// Alignment returns the Aligned's Alignment
	Alignment() Alignment

	// SetWhitespace sets the Element's whitespace mode
	SetWhitespace(Whitespace) Element
	// Whitespace returns the Element's whitespace mode
	Whitespace() Whitespace

	// SetPadding sets the Padded's padding.
	SetPadding(Padding) Element
	// Padding returns the padding for the Padded.
	Padding() Padding

	// SetBounds sets the Element's bounding box.
	SetBorder(Border) Element
	// Bounds returns the bounding box for the Element.
	Border() *Border
	// SetBorderForegroundColor sets the Element's border foreground color
	// (i.e the color of the border's cells underlying grapheme).
	SetBorderForegroundColor(Color) Element
	// BorderForegroundColor returns the Element's border foreground color.
	BorderForegroundColor() Color
	// SetBorderBackgroundColor sets the Element's border background color
	// (i.e the background color of the border's cells.
	SetBorderBackgroundColor(Color) Element
	// BorderBackgroundColor returns the Element's border background color.
	BorderBackgroundColor() Color

	// HorizontalSpace returns the number of cells consumed by the element's
	// left-right padding and border.
	HorizontalSpace() Dimension
	// VerticalSpace returns the number of lines consumed by the element's
	// top-bottom padding and border
	VerticalSpace() Dimension

	// SetStyle applies the supplied Style to the Styled.
	SetStyle(Style) Element
	// Style returns the thing's Style
	Style() Style
	// SetForegroundColor sets the thing's foreground color
	SetForegroundColor(color.Color) Element
	// ForegroundColor returns the Element's foreground color.
	ForegroundColor() Color
	// SetBackgroundColor sets the thing's background color
	SetBackgroundColor(color.Color) Element
	// BackgroundColor returns the Element's background color.
	BackgroundColor() Color

	// NodeInternalID returns a dotted-notation identifier for the node within
	// the tree. Each number in the returned string indicates the child index
	// of this Node's ancestors.
	//
	// So, "0.3" means "the fourth child of the first child of the root node".
	// Returns "root" for the root node.
	NodeInternalID() string
	// ChildIndex returns the Element's index within the Element's parent's
	// collection of children.
	ChildIndex() int
	// SetParent sets the Element's parent and index of the Element within the
	// parent's children.
	SetParent(Element, int) Element
	// Parent returns the Element that is the parent of this Element, or nil if this
	// is a root Element.
	Parent() Element
	// AppendChild adds a new child Element to the Element at the end of Element's set of
	// children.
	AppendChild(Element)
	// PopChild removes the last child Element from the Element's children and returns
	// it. Returns nil if Element has no children.
	PopChild() Element
	// RemoveAllChildren removes any children from this Node.
	RemoveAllChildren()
	// Children returns a slice of Elements that are children of this Element.
	Children() []Element
	// HasChildren returns whether the Element has children.
	HasChildren() bool
	// FirstChild returns the Element that is the first child of this Element, or nil
	// if there are no children.
	FirstChild() Element
	// LastChild returns the Element that is the last child of this Element, or nil
	// if there are no children.
	LastChild() Element
	// ChildAt returns the child element at the supplied zero-based index, or nil
	// if the index is out of bounds.
	ChildAt(int) Element
	// NextSibling() returns the Element that is the next child of this Element's
	// parent, or nil if there is none.
	NextSibling() Element
	// PreviousSibling returns the Element that is the previous child of the
	// Element's parent, or nil if this Element is the first child of the parent
	// Element.
	PreviousSibling() Element
	// PreviousSiblings returns all Elements that are children of the Element's
	// parent before this Element, or an empty slice of Elements if this
	// Element is the first child of the parent Element.
	PreviousSiblings() []Element
}
