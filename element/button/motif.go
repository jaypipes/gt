package button

import (
	"github.com/jaypipes/gt/core/motif"
	"github.com/jaypipes/gt/core/palette"
)

var (
	DefaultMotif = motif.NordDarkPrimary
)

func init() {
	if !palette.TerminalHasDarkBackground {
		DefaultMotif = motif.NordLightPrimary
	}
}
