package motif

import (
	"image/color"

	"github.com/jaypipes/gt/core/border"
	"github.com/jaypipes/gt/core/palette"
	"github.com/jaypipes/gt/core/style"
)

var (
	// Motifs using Nord palette with dark terminal screen backgrounds.
	NordDarkPrimaryNormalBackgroundColor = palette.NordDarkPrimaryBackground
	NordDarkPrimaryNormalForegroundColor = palette.NordDarkPrimaryForeground
	NordDarkPrimaryNormalStyle           = style.New(
		style.WithForegroundColor(
			NordDarkPrimaryNormalForegroundColor,
		),
		style.WithBackgroundColor(
			NordDarkPrimaryNormalBackgroundColor,
		),
	)
	NordDarkPrimaryNormalBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordDarkPrimaryNormalBackgroundColor)

	NordDarkPrimaryHoveredBackgroundColor = palette.NordDarkPrimaryForeground
	NordDarkPrimaryHoveredForegroundColor = palette.NordDarkPrimaryBackground
	NordDarkPrimaryHoveredStyle           = style.New(
		style.WithForegroundColor(
			NordDarkPrimaryHoveredForegroundColor,
		),
		style.WithBackgroundColor(
			NordDarkPrimaryHoveredBackgroundColor,
		),
	)
	NordDarkPrimaryHoveredBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordDarkPrimaryHoveredBackgroundColor)

	NordDarkPrimaryFocusedBackgroundColor = palette.NordDarkContrastForeground
	NordDarkPrimaryFocusedForegroundColor = palette.NordDarkContrastBackground
	NordDarkPrimaryFocusedStyle           = style.New(
		style.WithForegroundColor(
			NordDarkPrimaryFocusedForegroundColor,
		),
		style.WithBackgroundColor(
			NordDarkPrimaryFocusedBackgroundColor,
		),
	)
	NordDarkPrimaryFocusedBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordDarkPrimaryFocusedBackgroundColor)

	NordDarkPrimary = New(
		WithNormalStyle(NordDarkPrimaryNormalStyle),
		WithNormalBorder(NordDarkPrimaryNormalBorder),
		WithFocusedStyle(NordDarkPrimaryFocusedStyle),
		WithFocusedBorder(NordDarkPrimaryFocusedBorder),
		WithHoveredStyle(NordDarkPrimaryHoveredStyle),
		WithHoveredBorder(NordDarkPrimaryHoveredBorder),
	)

	NordDarkContrastNormalBackgroundColor = palette.NordDarkContrastBackground
	NordDarkContrastNormalForegroundColor = palette.NordDarkContrastForeground
	NordDarkContrastNormalStyle           = style.New(
		style.WithForegroundColor(
			NordDarkContrastNormalForegroundColor,
		),
		style.WithBackgroundColor(
			NordDarkContrastNormalBackgroundColor,
		),
	)
	NordDarkContrastNormalBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordDarkContrastNormalBackgroundColor)

	NordDarkContrastFocusedBackgroundColor = palette.NordDarkPrimaryForeground
	NordDarkContrastFocusedForegroundColor = palette.NordDarkPrimaryBackground
	NordDarkContrastFocusedStyle           = style.New(
		style.WithForegroundColor(
			NordDarkContrastFocusedForegroundColor,
		),
		style.WithBackgroundColor(
			NordDarkContrastFocusedBackgroundColor,
		),
	)
	NordDarkContrastFocusedBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordDarkContrastFocusedBackgroundColor)

	NordDarkContrastHoveredBackgroundColor = palette.NordDarkContrastForeground
	NordDarkContrastHoveredForegroundColor = palette.NordDarkContrastBackground
	NordDarkContrastHoveredStyle           = style.New(
		style.WithForegroundColor(
			NordDarkContrastHoveredForegroundColor,
		),
		style.WithBackgroundColor(
			NordDarkContrastHoveredBackgroundColor,
		),
	)
	NordDarkContrastHoveredBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordDarkContrastHoveredBackgroundColor)

	NordDarkContrast = New(
		WithNormalStyle(NordDarkContrastNormalStyle),
		WithNormalBorder(NordDarkContrastNormalBorder),
		WithFocusedStyle(NordDarkContrastFocusedStyle),
		WithFocusedBorder(NordDarkContrastFocusedBorder),
		WithHoveredStyle(NordDarkContrastHoveredStyle),
		WithHoveredBorder(NordDarkContrastHoveredBorder),
	)

	// Motifs using Nord palette with light terminal screen backgrounds.
	NordLightPrimaryNormalBackgroundColor = palette.NordLightPrimaryBackground
	NordLightPrimaryNormalForegroundColor = palette.NordLightPrimaryForeground
	NordLightPrimaryNormalStyle           = style.New(
		style.WithForegroundColor(
			NordLightPrimaryNormalForegroundColor,
		),
		style.WithBackgroundColor(
			NordLightPrimaryNormalBackgroundColor,
		),
	)
	NordLightPrimaryNormalBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordLightPrimaryNormalBackgroundColor)

	NordLightPrimaryHoveredBackgroundColor = palette.NordLightPrimaryForeground
	NordLightPrimaryHoveredForegroundColor = palette.NordLightPrimaryBackground
	NordLightPrimaryHoveredStyle           = style.New(
		style.WithForegroundColor(
			NordLightPrimaryHoveredForegroundColor,
		),
		style.WithBackgroundColor(
			NordLightPrimaryHoveredBackgroundColor,
		),
	)
	NordLightPrimaryHoveredBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordLightPrimaryHoveredBackgroundColor)

	NordLightPrimary = New(
		WithNormalStyle(NordLightPrimaryNormalStyle),
		WithNormalBorder(NordLightPrimaryNormalBorder),
		WithHoveredStyle(NordLightPrimaryHoveredStyle),
		WithHoveredBorder(NordLightPrimaryHoveredBorder),
	)

	NordLightContrastNormalBackgroundColor = palette.NordLightContrastBackground
	NordLightContrastNormalForegroundColor = palette.NordLightContrastForeground
	NordLightContrastNormalStyle           = style.New(
		style.WithForegroundColor(
			NordLightContrastNormalForegroundColor,
		),
		style.WithBackgroundColor(
			NordLightContrastNormalBackgroundColor,
		),
	)
	NordLightContrastNormalBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordLightContrastNormalBackgroundColor)

	NordLightContrastHoveredBackgroundColor = palette.NordLightContrastForeground
	NordLightContrastHoveredForegroundColor = palette.NordLightContrastBackground
	NordLightContrastHoveredStyle           = style.New(
		style.WithForegroundColor(
			NordLightContrastHoveredForegroundColor,
		),
		style.WithBackgroundColor(
			NordLightContrastHoveredBackgroundColor,
		),
	)
	NordLightContrastHoveredBorder = border.InnerHalfBlock().
					WithBackgroundColor(color.Transparent).
					WithForegroundColor(NordLightContrastHoveredBackgroundColor)

	NordLightContrast = New(
		WithNormalStyle(NordLightContrastNormalStyle),
		WithNormalBorder(NordLightContrastNormalBorder),
		WithHoveredStyle(NordLightContrastHoveredStyle),
		WithHoveredBorder(NordLightContrastHoveredBorder),
	)
)
