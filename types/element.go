package types

import "context"

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
// Element implements [types.Plottable] and [types.Renderable] which means that
// every Element can draw itself onto a [types.Screen].
type Element interface {
	Borderable
	FocusEventHandler
	Identifiable
	KeyPressEventHandler
	Motifable
	Themeable
	MouseEventHandler
	Node
	Plottable
	Renderable
	ScrollEventHandler
	Style

	// WithID sets the Element's unique identifier and returns the Element.
	WithID(string) Element
	// WithClass sets the Element's type/class and returns the Element.
	WithClass(string) Element
	// Class returns the Element's type/class, e.g. "gt.label" or "gt.canvas"
	Class() string
	// Tag returns a string with the Element's type/class and ID
	Tag() string

	// SetDisabled sets whether the Element is disabled. Disabled Elements cannot
	// receive the focus.
	SetDisabled(bool)
	// Disabled returns true if the Element cannot get the focus.
	Disabled() bool
	// WithDisabled sets whether the Element is disabled and returns the
	// Element.
	WithDisabled(bool) Element

	// WithFocusable sets whether the Element can receive the focus and returns
	// the Element.
	WithFocusable(bool) Element
	// NextFocusable returns the next focusable thing, or nil if there is no
	// next focusable thing. The Element's children will first be inspected and
	// then the next sibling Element.
	NextFocusable(context.Context) FocusEventHandler

	// WithParent sets the Element's parent and index of the Element within the
	// parent's children and returns the Element.
	WithParent(Node, int) Element
	// WithBounds sets the Element's outer bounding box and returns the Element.
	WithBounds(Rectangle) Element
	// WithAbsolutePosition sets the Element's outer bounding box's top-left
	// coordinates and marks the Element as using absolute positioning,
	// returning the Element.
	WithAbsolutePosition(Point) Element
	// WithSize constrains the size of the Element and returns the Element.
	WithSize(SizeConstraint) Element
	// WithWidth constrains the width of the Element and returns the Element.
	WithWidth(DimensionConstraint) Element
	// WithMinWidth sets the minimum width of the Element and returns the
	// Element.
	WithMinWidth(Dimension) Element
	// WithHeight constrains the height of the Element and returns the Element.
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
	// DisabledBorder returns the Border for the Element when the Element is
	// disabled.
	DisabledBorder() Border
	// SetDisabledBorder sets the Border for the Element when the Element is
	// disabled.
	SetDisabledBorder(Border)
	// WithDisabledBorder sets the Border for the Element when the Element is
	// disabled and returns the Element.
	WithDisabledBorder(Border) Element
	// FocusedBorder returns the Border for the Element when the Element has
	// the focus.
	FocusedBorder() Border
	// SetFocusedBorder sets the Border for the Element when the Element has
	// the focus.
	SetFocusedBorder(Border)
	// WithFocusedBorder sets the Border for the Element when the Element has
	// the focus and returns the Element.
	WithFocusedBorder(Border) Element
	// HoveredBorder returns the Border for the Element when the mouse is
	// hovering over the Element.
	HoveredBorder() Border
	// SetHoveredBorder sets the Border for the Element when the mouse is
	// hovering over the Element.
	SetHoveredBorder(Border)
	// WithHoveredBorder sets the Border for the Element when the mouse is
	// hovering over the Element and returns the Element.
	WithHoveredBorder(Border) Element
	// SetBorderForegroundColor sets the Borderable's border foreground color
	// (i.e the color of the border's cells underlying grapheme).
	SetBorderForegroundColor(Color)
	// BorderForegroundColor returns the Borderable's border foreground color.
	BorderForegroundColor() Color
	// SetBorderBackgroundColor sets the Borderable's border background color
	// (i.e the background color of the border's cells.
	SetBorderBackgroundColor(Color)
	// BorderBackgroundColor returns the Borderable's border background color.
	BorderBackgroundColor() Color
	// WithBorderForegroundColor sets the Element's border foreground color (i.e
	// the color of the border cell's underlying grapheme) and returns the
	// Element.
	WithBorderForegroundColor(Color) Element
	// WithBorderBackgroundColor sets the Element's border background color (i.e
	// the background color of the border cell's and returns the Element.
	WithBorderBackgroundColor(Color) Element

	// WithTheme sets the Element's theme class and returns the Element.
	WithThemeClass(ThemeClass) Element
	// WithTheme sets the Element's theme and returns the Element.
	WithTheme(Theme) Element
	// WithMotif sets the Element's motif and returns the Element.
	WithMotif(Motif) Element

	// Style returns the Element's Style.
	Style() Style
	// SetStyle sets the Element's style.
	SetStyle(Style)
	// WithStyle sets the Element's style and returns the Element.
	WithStyle(Style) Element
	// WithForegroundColor sets the thing's foreground color and returns the
	// Element.
	WithForegroundColor(Color) Element
	// WithBackgroundColor sets the Element's background color and returns the
	// Element.
	WithBackgroundColor(Color) Element
	// DisabledStyle returns the Element's Style when the Element is disabled.
	DisabledStyle() Style
	// SetDisabledStyle sets the Element's style when the Element is disabled.
	SetDisabledStyle(Style)
	// WithDisabledStyle sets the Element's style when the Element is disabled
	// and returns the Element.
	WithDisabledStyle(Style) Element
	// FocusedStyle returns the Element's Style when the Element has the focus.
	FocusedStyle() Style
	// SetFocusedStyle sets the Element's style when the Element has the focus.
	// When the Element no longer has the focus, its style will be reset.
	SetFocusedStyle(Style)
	// WithFocusedStyle sets the Element's style when the Element has the focus
	// and returns the Element.
	WithFocusedStyle(Style) Element
	// HoveredStyle returns the Element's Style when the Element has the hover.
	HoveredStyle() Style
	// SetHoveredStyle sets the Element's style when the mouse is currently
	// hovering over the Element and the Element does *NOT* have the focus.
	// When the mouse no longer hovers over the Element, its style will be
	// reset.
	SetHoveredStyle(Style)
	// WithHoveredStyle sets the Element's style when the mouse is currently
	// hovering over the Element and the Element does *NOT* have the focus.
	// WithHoveredStyle returns the Element.
	WithHoveredStyle(Style) Element

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
	// ScrollWidth returns the minimum number of cells (width) that the Element
	// would consume in order to fit all of its content in its containing box
	// without using a horizontal scrollbar.
	ScrollWidth() Dimension
	// ScrollHeight returns the minimum number of lines (height) that the
	// Element would consume in order to fit all of its content in its
	// containing box without using a vertical scrollbar.
	ScrollHeight() Dimension
}

// ElementWithOption describes an optional varg parameter to [element.New] that
// modifies the returned Element.
type ElementWithOption func(Element)
