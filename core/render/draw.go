package render

import (
	"context"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/types"
)

// Clear clears any rendered cell contents for the supplied bounding box.
func Clear(buf types.Screen, area types.Rectangle) {
	for y := area.Min.Y; y < area.Max.Y; y++ {
		for x := area.Min.X; x < area.Max.X; x++ {
			buf.SetCell(x, y, nil)
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

// Draw writes the supplied element's contents to the supplied Screen and
// outermost bounding box.
func Draw(
	ctx context.Context,
	el types.Element,
	screen types.Screen,
) {
	gtlog.Debug(ctx, "render.Draw[%s]", el.Tag())
	el.Draw(screen, el.Bounds())
	for _, child := range el.Children() {
		Draw(ctx, child, screen)
	}
}
