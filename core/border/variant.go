package border

import (
	"github.com/jaypipes/gt/core/cell"
	"github.com/jaypipes/gt/types"
)

// None returns an empty border.
func None() types.Border {
	return &Border{}
}

// Normal returns a standard-type border with a normal weight and 90
// degree corners.
func Normal() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent("─")),
		b:  cell.New(cell.WithContent("─")),
		l:  cell.New(cell.WithContent("│")),
		r:  cell.New(cell.WithContent("│")),
		tl: cell.New(cell.WithContent("┌")),
		tr: cell.New(cell.WithContent("┐")),
		bl: cell.New(cell.WithContent("└")),
		br: cell.New(cell.WithContent("┘")),
	}
}

// Rounded returns a border with rounded corners.
func Rounded() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent("─")),
		b:  cell.New(cell.WithContent("─")),
		l:  cell.New(cell.WithContent("│")),
		r:  cell.New(cell.WithContent("│")),
		tl: cell.New(cell.WithContent("╭")),
		tr: cell.New(cell.WithContent("╮")),
		bl: cell.New(cell.WithContent("╰")),
		br: cell.New(cell.WithContent("╯")),
	}
}

// Block returns a border that takes the whole block.
func Block() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent("█")),
		b:  cell.New(cell.WithContent("█")),
		l:  cell.New(cell.WithContent("█")),
		r:  cell.New(cell.WithContent("█")),
		tl: cell.New(cell.WithContent("█")),
		tr: cell.New(cell.WithContent("█")),
		bl: cell.New(cell.WithContent("█")),
		br: cell.New(cell.WithContent("█")),
	}
}

// OuterHalfBlock returns a half-block border that sits outside the frame.
func OuterHalfBlock() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent("▀")),
		b:  cell.New(cell.WithContent("▄")),
		l:  cell.New(cell.WithContent("▌")),
		r:  cell.New(cell.WithContent("▐")),
		tl: cell.New(cell.WithContent("▛")),
		tr: cell.New(cell.WithContent("▜")),
		bl: cell.New(cell.WithContent("▙")),
		br: cell.New(cell.WithContent("▟")),
	}
}

// InnerHalfBlock returns a half-block border that sits inside the frame.
func InnerHalfBlock() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent("▄")),
		b:  cell.New(cell.WithContent("▀")),
		l:  cell.New(cell.WithContent("▐")),
		r:  cell.New(cell.WithContent("▌")),
		tl: cell.New(cell.WithContent("▗")),
		tr: cell.New(cell.WithContent("▖")),
		bl: cell.New(cell.WithContent("▝")),
		br: cell.New(cell.WithContent("▘")),
	}
}

// Thick returns a border that's thicker than the one returned by
// Normal.
func Thick() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent("━")),
		b:  cell.New(cell.WithContent("━")),
		l:  cell.New(cell.WithContent("┃")),
		r:  cell.New(cell.WithContent("┃")),
		tl: cell.New(cell.WithContent("┏")),
		tr: cell.New(cell.WithContent("┓")),
		bl: cell.New(cell.WithContent("┗")),
		br: cell.New(cell.WithContent("┛")),
	}
}

// Double returns a border comprised of two thin strokes.
func Double() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent("═")),
		b:  cell.New(cell.WithContent("═")),
		l:  cell.New(cell.WithContent("║")),
		r:  cell.New(cell.WithContent("║")),
		tl: cell.New(cell.WithContent("╔")),
		tr: cell.New(cell.WithContent("╗")),
		bl: cell.New(cell.WithContent("╚")),
		br: cell.New(cell.WithContent("╝")),
	}
}

// Hidden returns a border that renders as a series of single-cell spaces. It's
// useful for cases when you want to remove a standard border but maintain
// layout positioning. This said, you can still apply a background color to a
// hidden border.
func Hidden() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent(" ")),
		b:  cell.New(cell.WithContent(" ")),
		l:  cell.New(cell.WithContent(" ")),
		r:  cell.New(cell.WithContent(" ")),
		tl: cell.New(cell.WithContent(" ")),
		tr: cell.New(cell.WithContent(" ")),
		bl: cell.New(cell.WithContent(" ")),
		br: cell.New(cell.WithContent(" ")),
	}
}

// Markdown return a table border in markdown style.
func Markdown() types.Border {
	return &Border{
		l:  cell.New(cell.WithContent("|")),
		r:  cell.New(cell.WithContent("|")),
		tl: cell.New(cell.WithContent("|")),
		tr: cell.New(cell.WithContent("|")),
		bl: cell.New(cell.WithContent("|")),
		br: cell.New(cell.WithContent("|")),
	}
}

// ASCII returns a table border with ASCII characters.
func ASCII() types.Border {
	return &Border{
		t:  cell.New(cell.WithContent("-")),
		b:  cell.New(cell.WithContent("-")),
		l:  cell.New(cell.WithContent("|")),
		r:  cell.New(cell.WithContent("|")),
		tl: cell.New(cell.WithContent("+")),
		tr: cell.New(cell.WithContent("+")),
		bl: cell.New(cell.WithContent("+")),
		br: cell.New(cell.WithContent("+")),
	}
}
