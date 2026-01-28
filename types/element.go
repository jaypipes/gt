package types

import (
	"image/color"
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
	Plottable
	Drawable

	// WithID sets the Element's unique identifier and returns the Element.
	WithID(string) Element
	// WithClass sets the Element's type/class and returns the Element.
	WithClass(string) Element
	// Class returns the Element's type/class, e.g. "gt.label" or "gt.canvas"
	Class() string
	// Tag returns a string with the Element's type/class and ID
	Tag() string

	// WithParent sets the Element's parent and index of the Element within the
	// parent's children and returns the Element.
	WithParent(Element, int) Element
	// WithBounds sets the Element's outer bounding box and returns the Element.
	WithBounds(Rectangle) Element
	// WithAbsolutePosition sets the Element's outer bounding box's top-left
	// coordinates and marks the Element as using absolute positioning,
	// returning the Element.
	WithAbsolutePosition(Point) Element
	// WithSize constrains the size of the Element and returns the Element.
	WithSize(SizeConstraint) Element
	// SetWidth constrains the width of the Element and returns the Element.
	WithWidth(DimensionConstraint) Element
	// WithMinWidth sets the minimum width of the Element and returns the
	// Element.
	WithMinWidth(Dimension) Element
	// SetHeight constrains the height of the Element and returns the Element.
	WithHeight(DimensionConstraint) Element
	// WithMinHeight sets the minimum height of the Element and returns the
	// Element.
	WithMinHeight(Dimension) Element
	// WithDisplayMode sets the display mode of the Element and returns the
	// Element.
	WithDisplay(Display) Element
	// WithAlignment sets the Element's Alignment and returns the Element
	WithAlignment(Alignment) Element
	// WithWhitespace sets the Element's whitespace mode and returns the
	// Element.
	WithWhitespace(Whitespace) Element
	// WithPadding sets the Padded's padding and returns the Element.
	WithPadding(Padding) Element
	// WithBorder sets the Element's border and returns the Element.
	WithBorder(Border) Element
	// WithBorderForegroundColor sets the Element's border foreground color (i.e
	// the color of the border's cells underlying grapheme) and returns the
	// Element.
	WithBorderForegroundColor(Color) Element
	// WithBorderBackgroundColor sets the Element's border background color (i.e
	// the background color of the border's cells and returns the Element.
	WithBorderBackgroundColor(Color) Element

	// SetTextContent sets the Element's raw, unstyled text contents.
	SetTextContent(string)
	// WithTextContent sets the Element's raw, unstyled text contents and
	// returns the Element.
	WithTextContent(string) Element
	// TextContent returns the Element's raw string contents.
	TextContent() string
	// TextContentWidth returns the width of the Element's raw string contents.
	TextContentWidth() Dimension
	// TextContentHeight returns the height of the Element's raw string
	// contents.
	TextContentHeight() Dimension
	// SetStyle sets the Element's style.
	SetStyle(Style)
	// WithStyle applies the supplied Style to the Element's text content and
	// returns the Element.
	WithStyle(Style) Element
	// Style returns the thing's Style
	Style() Style
	// SetForegroundColor sets the Element's foreground color.
	SetForegroundColor(color.Color)
	// WithForegroundColor sets the thing's foreground color and  returns the
	// Element.
	WithForegroundColor(color.Color) Element
	// ForegroundColor returns the Element's foreground color.
	ForegroundColor() Color
	// SetBackgroundColor sets the Element's background color.
	SetBackgroundColor(color.Color)
	// WithBackgroundColor sets the Element's background color and returns the
	// Element.
	WithBackgroundColor(color.Color) Element
	// BackgroundColor returns the Element's background color.
	BackgroundColor() Color
}
