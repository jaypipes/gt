package types

// Plotted describes something that can be plotted on a Screen.
//
// # Determining size and bounds.
//
// A Plotted has an outer bounding box and an inner bounding box. The outer
// bounding box represents the outer edge of the Plotted. The outer bounding
// box's cells contain the Plotted's border, if any. The inner bounding box
// represents the outer edge of the Plotted's *content* after accounting for
// any padding. The Plotted's content may be some text or it may be the content
// of child elements of the Plotted.
//
// ## No border, no padding, no content
//
// An empty Plotted has no size, no border, and no padding. It can be
// considered simply a Point at (0,0) on the Screen.
//
// ## No border, no padding, some content
//
// The inner bounding box will be equal to the outer bounding box when the
// Plotted has no border or padding.
//
// Consider a Plotted that has no border or padding and its content is the
// string "Hello". You can envision the Plotted laid out on the Screen like so:
//
// |     0   1   2   3   4   5   6   7   8
// |
// | 0   H   e   l   l   o
// |
// | 1
// |
// | 2
// |
// | 3
// |
// | 4
//
// Here's what the Plotted's methods would return:
//
// * OuterBounds(): (0,0)-(0,4)
// * OuterWidth(): 5
// * OuterHeight(): 1
// * InnerBounds(): (0,0)-(0,4)
// * InnerWidth(): 5
// * InnerHeight(): 1
//
// ## A border, no padding, some content
//
// Consider a Plotted that has a border but no padding and its content is the
// string "Hello". You can envision the Plotted laid out on the Screen like so:
//
// |     0   1   2   3   4   5   6   7   8
// |
// | 0   B   B   B   B   B   B   B
// |
// | 1   B   H   e   l   l   o   B
// |
// | 2   B   B   B   B   B   B   B
// |
// | 3
// |
// | 4
//
// In the above diagram, the letter "B" has been placed in the cells where the
// border will be drawn. The letters for the content "Hello" have been placed
// in their appropriate cells.
//
// Here's what the Plotted's methods would return:
//
// * OuterBounds(): (0,0)-(3,6)
// * OuterWidth(): 7
// * OuterHeight(): 3
// * InnerBounds(): (1,1)-(1,5)
// * InnerWidth(): 5
// * InnerHeight(): 1
//
// ## A border, some padding, some content
//
// Consider a Plotted has a border, a padding of 1 on all sides and its content
// is the string "Hello".  You can envision the Plotted's two bounding boxes
// like so:
//
// |     0   1   2   3   4   5   6   7   8
// |
// | 0   B   B   B   B   B   B   B   B   B
// |
// | 1   B   P   P   P   P   P   P   P   B
// |
// | 2   B   P   H   e   l   l   o   P   B
// |
// | 3   B   P   P   P   P   P   P   P   B
// |
// | 4   B   B   B   B   B   B   B   B   B
//
// In the above diagram, the letter "B" has been placed in the cells where the
// border will be drawn. The letter "P" has been placed in the cells where the
// padding takes up some width and height. And the letters for the content
// "Hello" have been placed in their appropriate cells.
//
// Here's what the Plotted's methods would return:
//
// * OuterBounds(): (0,0)-(4,8)
// * OuterWidth(): 9
// * OuterHeight(): 5
// * InnerBounds(): (2,2)-(2,6)
// * InnerWidth(): 5
// * InnerHeight(): 1
//
// # Impact of fixed width or height
//
// When the SetSize(), SetWidth() and SetHeight() methods are called on a
// Plotted, the user is saying that they want the Plotted's *inner bounding
// box* to be a specific fixed width and/or height.
//
// In other words, if the Plotted with a fixed width and/or height has a border
// or padding, that border and padding will cause the Plotted's OuterHeight()
// and OuterWidth() to be more than the specified fixed width and height.
type Plotted interface {
	Aligned
	Bordered
	Displayed
	Padded
	Sized
	// Bounds returns the Plotted's outer bounding box.
	Bounds() Rectangle

	// SetBounds sets the Plotted's outer bounding box.
	SetBounds(Rectangle)

	// TL returns the Plotted's outer bounding box's top-left coordinates.
	TL() Point

	// TR returns the Plotted's outer bounding box's top-right coordinates.
	TR() Point

	// MinY returns the Min Y (top) of the Plotted's outer bounding box.
	MinY() int

	// MaxY returns the Max Y (bottom) of the Plotted's outer bounding box.
	MaxY() int

	// SetAbsolutePosition sets the Plotted's outer bounding box's top-left
	// coordinates and marks the Plotted as using absolute positioning.
	SetAbsolutePosition(Point)

	// HasAbsolutePosition returns true if the Plotted used absolute positioning.
	HasAbsolutePosition() bool
	// InnerBounds returns the inner bounding box for the Plotted, which is the
	// outer bounding box adjusted for any border and padding.
	InnerBounds() Rectangle
}
