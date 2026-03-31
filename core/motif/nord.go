package motif

import (
	"image/color"

	"github.com/jaypipes/gt/core/border"
	"github.com/jaypipes/gt/core/palette"
	"github.com/jaypipes/gt/core/style"
)

var (
	NordNormalBackgroundColor = color.Transparent
	NordNormalForegroundColor = palette.Nord.Grayscale(6)
	NordNormalStyle           = style.New(
		style.WithForegroundColor(
			NordNormalForegroundColor,
		),
		style.WithBackgroundColor(
			NordNormalBackgroundColor,
		),
	)
	NordNormalBorder = border.Rounded().
				WithBackgroundColor(NordNormalBackgroundColor).
				WithForegroundColor(NordNormalForegroundColor)

	NordHoveredBackgroundColor = palette.Nord.Grayscale(6)
	NordHoveredForegroundColor = palette.Nord.Grayscale(0)
	NordHoveredStyle           = style.New(
		style.WithForegroundColor(
			NordHoveredForegroundColor,
		),
		style.WithBackgroundColor(
			NordHoveredBackgroundColor,
		),
	)
	NordHoveredBorder = border.InnerHalfBlock().
				WithBackgroundColor(color.Transparent).
				WithForegroundColor(NordHoveredBackgroundColor)

	NordPrimary = New(
		WithNormalStyle(NordNormalStyle),
		WithNormalBorder(NordNormalBorder),
		WithHoveredStyle(NordHoveredStyle),
		WithHoveredBorder(NordHoveredBorder),
	)
	Nord = NordPrimary
)
