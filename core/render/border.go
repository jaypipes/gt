package render

import (
	"github.com/charmbracelet/x/ansi"
	"github.com/jaypipes/gt/core/types"
)

// BorderHorizontalSpace returns the number of cells the supplied Border
// consumes.
func BorderHorizontalSpace(b types.Border) types.Dimension {
	return types.Dimension(
		ansi.StringWidth(b.Left.Content) +
			ansi.StringWidth(b.Right.Content),
	)
}

// BorderVerticalSpace returns the number lines the supplied Border consumes.
func BorderVerticalSpace(b types.Border) types.Dimension {
	return types.Dimension(
		ansi.StringWidth(b.Top.Content) +
			ansi.StringWidth(b.Bottom.Content),
	)
}
