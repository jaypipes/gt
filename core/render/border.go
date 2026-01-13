package render

import (
	"github.com/charmbracelet/x/ansi"
	"github.com/jaypipes/gt/core/types"
)

// BorderHorizontalWidth returns the width, in cells, that the supplied Border
// consumes.
func BorderHorizontalWidth(b types.Border) int {
	return ansi.StringWidth(b.Left.Content) + ansi.StringWidth(b.Right.Content)
}

// BorderVerticalHeight returns the height, in lines, that the supplied Border
// consumes.
func BorderVerticalHeight(b types.Border) int {
	return ansi.StringWidth(b.Top.Content) + ansi.StringWidth(b.Bottom.Content)
}
