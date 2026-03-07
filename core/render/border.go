package render

import (
	"github.com/charmbracelet/x/ansi"

	"github.com/jaypipes/gt/types"
)

// BorderHorizontalSpace returns the number of cells the supplied Border
// consumes.
func BorderHorizontalSpace(b types.Border) types.Dimension {
	if b == nil {
		return types.Dimension(0)
	}
	return types.Dimension(
		max(
			ansi.StringWidth(b.L().Content()),
			ansi.StringWidth(b.TL().Content()),
			ansi.StringWidth(b.BL().Content()),
		) + max(
			ansi.StringWidth(b.R().Content()),
			ansi.StringWidth(b.TR().Content()),
			ansi.StringWidth(b.BR().Content()),
		),
	)
}

// BorderVerticalSpace returns the number lines the supplied Border consumes.
func BorderVerticalSpace(b types.Border) types.Dimension {
	if b == nil {
		return types.Dimension(0)
	}
	return types.Dimension(
		max(
			ansi.StringWidth(b.T().Content()),
			ansi.StringWidth(b.TL().Content()),
			ansi.StringWidth(b.TR().Content()),
		) + max(
			ansi.StringWidth(b.B().Content()),
			ansi.StringWidth(b.BR().Content()),
			ansi.StringWidth(b.BL().Content()),
		),
	)
}
