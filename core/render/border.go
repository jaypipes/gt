package render

import (
	"github.com/charmbracelet/x/ansi"

	"github.com/jaypipes/gt/types"
)

// BorderHorizontalSpace returns the number of cells the supplied Border
// consumes.
func BorderHorizontalSpace(b types.Border) types.Dimension {
	return types.Dimension(
		max(
			ansi.StringWidth(b.Left.Content),
			ansi.StringWidth(b.TopLeft.Content),
			ansi.StringWidth(b.BottomLeft.Content),
		) + max(
			ansi.StringWidth(b.Right.Content),
			ansi.StringWidth(b.TopRight.Content),
			ansi.StringWidth(b.BottomRight.Content),
		),
	)
}

// BorderVerticalSpace returns the number lines the supplied Border consumes.
func BorderVerticalSpace(b types.Border) types.Dimension {
	return types.Dimension(
		max(
			ansi.StringWidth(b.Top.Content),
			ansi.StringWidth(b.TopLeft.Content),
			ansi.StringWidth(b.TopRight.Content),
		) + max(
			ansi.StringWidth(b.Bottom.Content),
			ansi.StringWidth(b.BottomRight.Content),
			ansi.StringWidth(b.BottomLeft.Content),
		),
	)
}
