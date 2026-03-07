package render

import (
	"github.com/jaypipes/gt/types"
)

// Overlapping returns the rectangle representing the overlapping area of the
// two supplied rectangles. If either of the supplied rectangles is empty, the
// non-empty rectangle is returned.
func Overlapping(a, b types.Rectangle) types.Rectangle {
	if a.Empty() {
		return b
	}
	if b.Empty() {
		return a
	}
	return a.Intersect(b)
}
