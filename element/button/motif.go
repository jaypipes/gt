package button

import (
	"image/color"

	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt/core/border"
	"github.com/jaypipes/gt/core/motif"
	"github.com/jaypipes/gt/core/style"
)

var (
	veryLightGrey, _             = colorful.Hex("#ececec")
	black, _                     = colorful.Hex("#000000")
	DefaultNormalBackgroundColor = color.Transparent
	DefaultNormalForegroundColor = veryLightGrey
	DefaultNormalStyle           = style.New(
		style.WithForegroundColor(
			DefaultNormalForegroundColor,
		),
		style.WithBackgroundColor(
			DefaultNormalBackgroundColor,
		),
	)
	DefaultNormalBorder = border.Rounded().
				WithBackgroundColor(DefaultNormalBackgroundColor).
				WithForegroundColor(DefaultNormalForegroundColor)
	// Default hover is inverse of normal style.
	DefaultHoveredBackgroundColor = veryLightGrey
	DefaultHoveredForegroundColor = black
	DefaultHoveredStyle           = style.New(
		style.WithForegroundColor(
			black,
		),
		style.WithBackgroundColor(
			DefaultHoveredBackgroundColor,
		),
	)
	DefaultHoveredBorder = border.InnerHalfBlock().
				WithBackgroundColor(color.Transparent).
				WithForegroundColor(DefaultHoveredBackgroundColor)
	DefaultMotif = motif.Nord
)
