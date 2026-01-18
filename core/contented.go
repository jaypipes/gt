package core

import (
	"strings"

	"github.com/charmbracelet/x/ansi"

	"github.com/jaypipes/gt/core/types"
)

// Contented can have some raw string content.
type Contented struct {
	// content is the unstyled raw string content.
	content string
}

// SetContent sets the Contented's raw string content.
func (c *Contented) SetContent(content string) {
	c.content = content
}

// Content returns the Contented's raw string content.
func (c *Contented) Content() string {
	return c.content
}

// ContentWidth returns width in cells of the Contented's raw string content.
func (c *Contented) ContentWidth() types.Dimension {
	return types.Dimension(ansi.StringWidth(c.content))
}

// ContentHeight returns the height in lines of the Contented's raw string
// content.
func (c *Contented) ContentHeight() types.Dimension {
	return types.Dimension(strings.Count(c.content, "\n")) + 1
}

var _ types.Contented = (*Contented)(nil)
