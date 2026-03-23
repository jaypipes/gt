package box

import (
	"context"
	"image/color"

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
	h types.ScreenHandler,
) {
	// If we have a border, draw it around the outer bounding box.
	border := b.border
	if border == nil {
		return
	}

	screen := h.Screen()

	bounds := b.bounds

	gtlog.Debug(ctx, "Box.renderBorder: bounds=%s", bounds)

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

	minX := bounds.Min.X
	maxX := bounds.Max.X - 1
	minY := bounds.Min.Y
	maxY := bounds.Max.Y - 1

	// Draw the corners

	tl := border.TL()
	if tl != nil {
		s := tl.Style()
		if s == nil {
			s = defStyle
		}
		screen.PutStrStyled(minX, minY, tl.Content(), style.TCell(s))
	}

	tr := border.TR()
	if tr != nil {
		s := tr.Style()
		if s == nil {
			s = defStyle
		}
		screen.PutStrStyled(maxX, minY, tr.Content(), style.TCell(s))
	}

	bl := border.BL()
	if bl != nil {
		s := bl.Style()
		if s == nil {
			s = defStyle
		}
		screen.PutStrStyled(minX, maxY, bl.Content(), style.TCell(s))
	}

	br := border.BR()
	if br != nil {
		s := br.Style()
		if s == nil {
			s = defStyle
		}
		screen.PutStrStyled(maxX, maxY, br.Content(), style.TCell(s))
	}

	// Draw the edges

	te := border.T()
	if te != nil {
		s := te.Style()
		if s == nil {
			s = defStyle
		}
		ch := te.Content()
		for x := minX + 1; x < maxX; x++ {
			screen.PutStrStyled(x, minY, ch, style.TCell(s))
		}
	}

	be := border.B()
	if be != nil {
		s := be.Style()
		if s == nil {
			s = defStyle
		}
		ch := be.Content()
		for x := minX + 1; x < maxX; x++ {
			screen.PutStrStyled(x, maxY, ch, style.TCell(s))
		}
	}

	le := border.L()
	if le != nil {
		s := le.Style()
		if s == nil {
			s = defStyle
		}
		ch := le.Content()
		for y := minY + 1; y < maxY; y++ {
			screen.PutStrStyled(minX, y, ch, style.TCell(s))
		}
	}

	re := border.R()
	if re != nil {
		s := re.Style()
		if s == nil {
			s = defStyle
		}
		ch := re.Content()
		for y := minY + 1; y < maxY; y++ {
			screen.PutStrStyled(maxX, y, ch, style.TCell(s))
		}
	}
}
