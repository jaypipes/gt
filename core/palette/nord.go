package palette

import (
	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt/types"
)

// Nord color palette: https://www.nordtheme.com/docs/colors-and-palettes

var (
	Nord0, _   = colorful.Hex("#2e3440")
	Nord1, _   = colorful.Hex("#3b4252")
	Nord2, _   = colorful.Hex("#434c5e")
	Nord3, _   = colorful.Hex("#4c566a")
	Nord4, _   = colorful.Hex("#d8dee9")
	Nord5, _   = colorful.Hex("#e5e9f0")
	Nord6, _   = colorful.Hex("#eceff4")
	Nord7, _   = colorful.Hex("#8fbcbb") // blue green
	Nord8, _   = colorful.Hex("#88c0d0") // ice blue
	Nord9, _   = colorful.Hex("#81a1c1") // grey blue
	Nord10, _  = colorful.Hex("#5e81ac") // dark gray blue
	Nord11, _  = colorful.Hex("#bf616a") // red
	Nord12, _  = colorful.Hex("#d08770") // orange
	Nord13, _  = colorful.Hex("#ebcb8b") // yellow
	Nord14, _  = colorful.Hex("#a3be8c") // green
	Nord15, _  = colorful.Hex("#b48ead") // purple
	NordColors = types.PaletteColors{
		Nord0,
		Nord1,
		Nord2,
		Nord3,
		Nord4,
		Nord5,
		Nord6,
		Nord7,
		Nord8,
		Nord9,
		Nord10,
		Nord11,
		Nord12,
		Nord13,
		Nord14,
		Nord15,
	}

	NordDarkPrimaryBackground   = Nord10
	NordDarkPrimaryForeground   = Nord6
	NordDarkContrastBackground  = Nord9
	NordDarkContrastForeground  = Nord0
	NordLightPrimaryBackground  = Nord6
	NordLightPrimaryForeground  = Nord3
	NordLightContrastBackground = Nord4
	NordLightContrastForeground = Nord3

	NordError   = Nord11
	NordWarning = Nord12
	NordInfo    = Nord14
	NordSuccess = Nord15
)

var (
	Nord = New(WithColors(NordColors))
)
