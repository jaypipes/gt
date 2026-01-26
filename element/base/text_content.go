package base

import (
	"strings"

	"github.com/charmbracelet/x/ansi"

	"github.com/jaypipes/gt/types"
)

// SetTextContent sets the Element's raw, unstyled text content.
func (b *Base) SetTextContent(textContent string) {
	b.textContent = textContent
}

// WithTextContent sets the Element's raw, unstyled text content and returns
// the Element.
func (b *Base) WithTextContent(textContent string) types.Element {
	b.textContent = textContent
	return b
}

// TextContent returns the Element's raw, unstyled text content.
func (b *Base) TextContent() string {
	return b.textContent
}

// TextContentWidth returns width in cells of the Element's raw, unstyled text
// content.
func (b *Base) TextContentWidth() types.Dimension {
	return types.Dimension(ansi.StringWidth(b.textContent))
}

// TextContentHeight returns the height in lines of the Element's raw, unstyled
// text content.
func (b *Base) TextContentHeight() types.Dimension {
	return types.Dimension(strings.Count(b.textContent, "\n")) + 1
}
