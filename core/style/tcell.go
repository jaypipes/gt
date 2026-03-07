package style

import (
	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/types"
)

// TCell returns a tcell.Style given a gt Style
func TCell(s types.Style) tcell.Style {
	out := tcell.StyleDefault
	if s.Bold() {
		out = out.Bold(true)
	}
	if s.Italic() {
		out = out.Italic(true)
	}
	if s.Dim() {
		out = out.Dim(true)
	}
	if s.Strikethrough() {
		out = out.StrikeThrough(true)
	}
	if s.Blink() {
		out = out.Blink(true)
	}
	if s.Underline() {
		params := []any{
			tcell.UnderlineStyle(s.UnderlineStyle()),
			s.UnderlineColor(),
		}
		out = out.Underline(params...)
	}
	return out
}
