package palette

import (
	"github.com/lucasb-eyer/go-colorful"

	"github.com/jaypipes/gt/types"
)

// Nord color palette: https://www.nordtheme.com/docs/colors-and-palettes

var (
	nord0, _   = colorful.Hex("#2e3440")
	nord1, _   = colorful.Hex("#3b4252")
	nord2, _   = colorful.Hex("#434c5e")
	nord3, _   = colorful.Hex("#4c566a")
	nord4, _   = colorful.Hex("#d8dee9")
	nord5, _   = colorful.Hex("#e5e9f0")
	nord6, _   = colorful.Hex("#eceff4")
	nord7, _   = colorful.Hex("#8fbcbb") // blue green
	nord8, _   = colorful.Hex("#88c0d0") // ice blue
	nord9, _   = colorful.Hex("#81a1c1") // grey blue
	nord10, _  = colorful.Hex("#5e81ac") // dark gray blue
	nord11, _  = colorful.Hex("#bf616a") // red
	nord12, _  = colorful.Hex("#d08770") // orange
	nord13, _  = colorful.Hex("#ebcb8b") // yellow
	nord14, _  = colorful.Hex("#a3be8c") // green
	nord15, _  = colorful.Hex("#b48ead") // purple
	nordColors = types.PaletteColors{
		nord0,
		nord1,
		nord2,
		nord3,
		nord4,
		nord5,
		nord6,
		nord7,
		nord8,
		nord9,
		nord10,
		nord11,
		nord12,
		nord13,
		nord14,
		nord15,
	}
)

var (
	Nord = New(WithColors(nordColors))
)
