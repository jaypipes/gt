package element

import (
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/samber/lo"

	"github.com/jaypipes/gt/types"
)

// WithSize constrains the size of the Element's outer bounding box and returns
// the Element.
func (e *Element) WithSize(constraint types.SizeConstraint) types.Element {
	e.Box.SetSize(constraint)
	return e
}

// WithWidth constrains the width of the Element and returns the Element.
func (e *Element) WithWidth(constraint types.DimensionConstraint) types.Element {
	e.Box.SetWidth(constraint)
	return e
}

// WithMinWidth sets the minimum width of the Element and returns the Element.
func (e *Element) WithMinWidth(w types.Dimension) types.Element {
	e.Box.SetMinWidth(w)
	return e
}

// WithHeight constrains the height of the Element and returns the Element.
func (e *Element) WithHeight(constraint types.DimensionConstraint) types.Element {
	e.Box.SetHeight(constraint)
	return e
}

// WithMinHeight sets the minimum height of the Element and returns the
// Element.
func (e *Element) WithMinHeight(w types.Dimension) types.Element {
	e.Box.SetMinHeight(w)
	return e
}

// ScrollWidth returns the minimum number of cells (width) that the Element
// would consume in order to fit all of its content on the screen without
// using a horizontal scrollbar.
func (e *Element) ScrollWidth() types.Dimension {
	whitespace := e.Whitespace()
	if whitespace&types.WhitespaceWrapNever != 0 {
		return types.Dimension(ansi.StringWidth(e.textContent))
	}
	content := e.TextContent()
	// Determine the widest line of text content.
	lines := strings.Split(content, "\n")
	return lo.Max(lo.Map(lines, func(line string, _ int) types.Dimension {
		return types.Dimension(ansi.StringWidth(line))
	}))
}

// ScrollHeight returns the minimum number of lines (height) that the
// Element would consume in order to fit all of its content on the screen
// without using a vertical scrollbar.
func (e *Element) ScrollHeight() types.Dimension {
	return types.Dimension(strings.Count(e.textContent, "\n")) + 1
}
