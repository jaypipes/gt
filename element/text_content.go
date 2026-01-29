package element

import (
	"strings"

	"github.com/charmbracelet/x/ansi"

	"github.com/jaypipes/gt/types"
)

// SetTextContent sets the Element's raw, unstyled text content.
func (e *Element) SetTextContent(textContent string) {
	e.textContent = textContent
}

// WithTextContent sets the Element's raw, unstyled text content and returns
// the Element.
func (e *Element) WithTextContent(textContent string) types.Element {
	e.textContent = textContent
	return e
}

// TextContent returns the Element's raw, unstyled text content.
func (e *Element) TextContent() string {
	return e.textContent
}

// TextContentWidth returns width in cells of the Element's raw, unstyled text
// content.
func (e *Element) TextContentWidth() types.Dimension {
	return types.Dimension(ansi.StringWidth(e.textContent))
}

// TextContentHeight returns the height in lines of the Element's raw, unstyled
// text content.
func (e *Element) TextContentHeight() types.Dimension {
	return types.Dimension(strings.Count(e.textContent, "\n")) + 1
}
