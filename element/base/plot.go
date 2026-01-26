package base

import (
	"context"
)

// Plotting describes the process of determining an outer bounding box and
// positioning coordinates for something on a Screen.
//
// # Determining size and bounds.
//
// An Element has an outer bounding box and an inner bounding box. The outer
// bounding box represents the outer edge of the Element. The outer bounding
// box's cells contain the Element's border, if any. The inner bounding box
// represents the outer edge of the Element's *content* after accounting for
// any padding. The Element's content may be some text or it may be the content
// of child elements of the Element.
//
// ## No border, no padding, no content
//
// An empty Element has no size, no border, and no padding. It can be
// considered simply a Point at (0,0) on the Screen.
//
// ## No border, no padding, some content
//
// The inner bounding box will be equal to the outer bounding box when the
// Element has no border or padding.
//
// Consider an Element that has no border or padding and its content is the
// string "Hello". You can envision the Element laid out on the Screen like so:
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
// Here's what the Element's methods would return:
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
// Consider an Element that has a border but no padding and its content is the
// string "Hello". You can envision the Element laid out on the Screen like so:
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
// Here's what the Element's methods would return:
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
// Consider an Element has a border, a padding of 1 on all sides and its
// content is the string "Hello".  You can envision the Element's two bounding
// boxes like so:
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
// Here's what the Element's methods would return:
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
// Element, the user is saying that they want the Element's *inner bounding
// box* to be a specific fixed width and/or height.
//
// In other words, if the Element with a fixed width and/or height has a border
// or padding, that border and padding will cause the Element's OuterHeight()
// and OuterWidth() to be more than the specified fixed width and height.

// Plot calculates the anchoring positioning coordinates of the element.
//
// It traverses the tree of elements rooted at this element and calculates the
// top left coordinates for the element.
//
// To calculate the top left (anchor point) coordinates of the element's
// bounding box, we use the following algorithm:
//
// If the element is using absolute positioning, its bounding box is anchored
// at the absolute coordinates. If the element is using relative positioning,
// the anchor point is calculated based on the element's Display property and
// is relative to the previous sibling or, if no previous sibling, the parent.
func (b *Base) Plot(ctx context.Context) {
	b.Box.Plot(ctx)
}
