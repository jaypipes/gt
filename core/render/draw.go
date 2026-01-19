package render

import (
	"github.com/jaypipes/gt/types"
)

// Clear clears any rendered cell contents for the supplied bounding box.
func Clear(screen types.Screen, area types.Rectangle) {
	for y := area.Min.Y; y < area.Max.Y; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			screen.SetCell(x, y, nil)
		}
	}
}

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
