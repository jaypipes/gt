package render

import (
	"github.com/jaypipes/gt/types"
)

// BorderHorizontalSpace returns the number of cells the supplied Border
// consumes.
func BorderHorizontalSpace(b types.Border) types.Dimension {
	if b == nil {
		return types.Dimension(0)
	}
	return b.HorizontalSpace()
}

// BorderVerticalSpace returns the number lines the supplied Border consumes.
func BorderVerticalSpace(b types.Border) types.Dimension {
	if b == nil {
		return types.Dimension(0)
	}
	return b.VerticalSpace()
}
