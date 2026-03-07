package box

import (
	"context"
	"image/color"

	"github.com/jaypipes/gt/core"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/style"
	"github.com/jaypipes/gt/types"
)

// SetBorder sets the Box's border.
func (b *Box) SetBorder(border types.Border) {
	b.border = border
}

// Border returns the Box's border, if any.
func (b *Box) Border() types.Border {
	return b.border
}

// SetBorderForegroundColor sets the Box's border foreground color
// (i.e the color of the border cell's underlying grapheme).
func (b *Box) SetBorderForegroundColor(c types.Color) {
	if b.border != nil {
		b.border.SetForegroundColor(c)
	}
}

// BorderForegroundColor returns the Box's border foreground color.
func (b *Box) BorderForegroundColor() types.Color {
	if b.border != nil {
		return b.border.ForegroundColor()
	}
	return color.Transparent
}

// SetBorderBackgroundColor sets the Box's border background color
// (i.e the background color of the border's cells.
func (b *Box) SetBorderBackgroundColor(c types.Color) {
	if b.border != nil {
		b.border.SetBackgroundColor(c)
	}
}

// BorderBackgroundColor returns the Box's border background color.
func (b *Box) BorderBackgroundColor() types.Color {
	if b.border != nil {
		return b.border.BackgroundColor()
	}
	return color.Transparent
}

// renderBorder draws the border around the outer bounding box's cells.
func (b *Box) renderBorder(
	ctx context.Context,
	screen types.Screen,
) {
	// If we have a border, draw it around the outer bounding box.
	border := b.border
	if border == nil {
		return
	}

	bounds := b.bounds

	gtlog.Debug(ctx, "Box[%s].renderBorder: bounds=%s", core.ID(b), bounds)

	// default style to use if a border edge has no style of its own
	defStyle := style.Empty()
	bfg := border.ForegroundColor()
	bbg := border.BackgroundColor()
	if bfg != nil {
		defStyle.SetForegroundColor(bfg)
	}
	if bbg != nil {
		defStyle.SetBackgroundColor(bbg)
	}

	minY := bounds.Min.Y
	maxY := bounds.Max.Y
	minX := bounds.Min.X
	maxX := bounds.Max.X
	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			switch {
			case y == minY && x == minX:
				corner := border.TL()
				if corner != nil {
					s := corner.Style()
					if s == nil {
						s = defStyle
					}
					screen.PutStrStyled(x, y, corner.Content(), style.TCell(s))
				}
			case y == minY && x == maxX-1:
				corner := border.TR()
				if corner != nil {
					s := corner.Style()
					if s == nil {
						s = defStyle
					}
					screen.PutStrStyled(x, y, corner.Content(), style.TCell(s))
				}
			case y == maxY-1 && x == minX:
				corner := border.BL()
				if corner != nil {
					s := corner.Style()
					if s == nil {
						s = defStyle
					}
					screen.PutStrStyled(x, y, corner.Content(), style.TCell(s))
				}
			case y == maxY-1 && x == maxX-1:
				corner := border.BR()
				if corner != nil {
					s := corner.Style()
					if s == nil {
						s = defStyle
					}
					screen.PutStrStyled(x, y, corner.Content(), style.TCell(s))
				}
			case y == minY:
				edge := border.T()
				if edge != nil {
					s := edge.Style()
					if s == nil {
						s = defStyle
					}
					screen.PutStrStyled(x, y, edge.Content(), style.TCell(s))
				}
			case y == maxY-1:
				edge := border.B()
				if edge != nil {
					s := edge.Style()
					if s == nil {
						s = defStyle
					}
					screen.PutStrStyled(x, y, edge.Content(), style.TCell(s))
				}
			case x == minX:
				edge := border.L()
				if edge != nil {
					s := edge.Style()
					if s == nil {
						s = defStyle
					}
					screen.PutStrStyled(x, y, edge.Content(), style.TCell(s))
				}
			case x == maxX-1:
				edge := border.R()
				if edge != nil {
					s := edge.Style()
					if s == nil {
						s = defStyle
					}
					screen.PutStrStyled(x, y, edge.Content(), style.TCell(s))
				}
			default:
				continue
			}
		}
	}
}
