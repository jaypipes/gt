package core

import "github.com/jaypipes/gt/core/types"

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

var _ types.Contented = (*Contented)(nil)
