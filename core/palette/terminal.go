package palette

import (
	"os"

	"github.com/muesli/termenv"
)

var (
	// TerminalHasDarkBackground will be true if we detect the default terminal
	// has a dark background.
	TerminalHasDarkBackground = true
)

func init() {
	o := termenv.NewOutput(os.Stdout)
	TerminalHasDarkBackground = o.HasDarkBackground()
}
