package core

import (
	"github.com/jaypipes/gt/core/types"
)

// Styled can have a Style.
type Styled struct {
	// style is the style mode of the Styled
	style types.Style
}

// SetStyle sets the Styled's style.
func (s *Styled) SetStyle(style types.Style) {
	s.style = style
}

// Style returns the Styled's Style.
func (s *Styled) Style() types.Style {
	return s.style
}

// SetForegroundColor sets the Styled's foreground color.
func (s *Styled) SetForegroundColor(c types.Color) {
	s.style.Fg = c
}

// SetBackgroundColor sets the Styled's foreground color.
func (s *Styled) SetBackgroundColor(c types.Color) {
	s.style.Bg = c
}

var _ types.Styled = (*Styled)(nil)
