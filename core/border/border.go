package border

import (
	"github.com/charmbracelet/x/ansi"
	"github.com/jaypipes/gt/types"
)

// Border decorates the outside of a Box
type Border struct {
	// t is the Cell used for the top edge of the Border
	t types.Cell
	// b is the Cell used for the bottom edge of the Border.
	b types.Cell
	// l is the Cell used for the left edge of the Border.
	l types.Cell
	// r is the Cell used for the right edge of the Border.
	r types.Cell
	// tl is the Cell used for the top left corner of the Border.
	tl types.Cell
	// tr is the Cell used for the top right corner of the Border.
	tr types.Cell
	// bl is the Cell used for the bottom left corner of the Border.
	bl types.Cell
	// br is the Cell used for the bottom right corner of the Border.
	br types.Cell
	// fgColor is the border foreground color (i.e the color of the border
	// cell's underlying grapheme). If all of the Border's edges are the same
	// foreground color, this will be set, otherwise individual edge Cells
	// might have their own Style.
	fgColor types.Color
	// bgColor is the border background color, i.b. the background color of the
	// border cells. If all of the Border's edges are the same background
	// color, this will be set, otherwise individual edge Cells might have
	// their own Style.
	bgColor types.Color
}

// Empty returns true if none of the Border's edges or corners have content.
func (b *Border) Empty() bool {
	return (b.t == nil || b.t.Empty()) &&
		(b.b == nil || b.b.Empty()) &&
		(b.l == nil || b.l.Empty()) &&
		(b.r == nil || b.r.Empty()) &&
		(b.tl == nil || b.tl.Empty()) &&
		(b.tr == nil || b.tr.Empty()) &&
		(b.bl == nil || b.bl.Empty()) &&
		(b.br == nil || b.br.Empty())
}

// T returns the Cell used for the top edge of the Border.
func (b *Border) T() types.Cell {
	return b.t
}

// TSize returns the number of lines the top edge of the Border will consume.
func (b *Border) TSize() types.Dimension {
	if b.t == nil {
		return types.Dimension(0)
	}
	return types.Dimension(ansi.StringWidth(b.t.Content()))
}

// SetT sets the Cell to be used for the top edge of the Border.
func (b *Border) SetT(cell types.Cell) {
	b.t = cell
}

// B returns the Cell used for the bottom edge of the Border.
func (b *Border) B() types.Cell {
	return b.b
}

// BSize returns the number of lines the bottom edge of the Border will
// consume.
func (b *Border) BSize() types.Dimension {
	if b.b == nil {
		return types.Dimension(0)
	}
	return types.Dimension(ansi.StringWidth(b.b.Content()))
}

// SetB sets the Cell to be used for the bottom edge of the Border.
func (b *Border) SetB(cell types.Cell) {
	b.b = cell
}

// L returns the Cell used for the left edge of the Border.
func (b *Border) L() types.Cell {
	return b.l
}

// LSize returns the number of lines the left edge of the Border will consume.
func (b *Border) LSize() types.Dimension {
	if b.l == nil {
		return types.Dimension(0)
	}
	return types.Dimension(ansi.StringWidth(b.l.Content()))
}

// SetT sets the Cell to be used for the left edge of the Border.
func (b *Border) SetL(cell types.Cell) {
	b.l = cell
}

// R returns the Cell used for the right edge of the Border.
func (b *Border) R() types.Cell {
	return b.r
}

// RSize returns the number of lines the right edge of the Border will consume.
func (b *Border) RSize() types.Dimension {
	if b.r == nil {
		return types.Dimension(0)
	}
	return types.Dimension(ansi.StringWidth(b.r.Content()))
}

// SetR sets the Cell to be used for the right edge of the Border.
func (b *Border) SetR(cell types.Cell) {
	b.r = cell
}

// TL returns the Cell used for the top left corner of the Border.
func (b *Border) TL() types.Cell {
	return b.tl
}

// SetTL sets the Cell to be used for the top left corner of the Border.
func (b *Border) SetTL(cell types.Cell) {
	b.tl = cell
}

// TR returns the Cell used for the top right corner of the Border.
func (b *Border) TR() types.Cell {
	return b.tr
}

// SetTL sets the Cell to be used for the top right corner of the Border.
func (b *Border) SetTR(cell types.Cell) {
	b.tr = cell
}

// BL returns the Cell used for the bottom left corner of the Border.
func (b *Border) BL() types.Cell {
	return b.bl
}

// SetBL sets the Cell to be used for the bottom left corner of the Border.
func (b *Border) SetBL(cell types.Cell) {
	b.bl = cell
}

// BR returns the Cell used for the bottom right corner of the Border.
func (b *Border) BR() types.Cell {
	return b.br
}

// SetBR sets the Cell to be used for the bottom right corner of the Border.
func (b *Border) SetBR(cell types.Cell) {
	b.br = cell
}

// ForegroundColor returns the foreground color used for all the Border's
// Cells. If this is empty, individual border Cells may have their own Style.
func (b *Border) ForegroundColor() types.Color {
	return b.fgColor
}

// SetForegroundColor sets the foreground color for all the Border's Cells.
func (b *Border) SetForegroundColor(color types.Color) {
	b.fgColor = color
}

// WithForegroundColor sets the foreground color for all the Border's Cells and
// returns the Border.
func (b *Border) WithForegroundColor(color types.Color) types.Border {
	b.fgColor = color
	return b
}

// BackgroundColor returns the background color used for all the Border's
// Cells. If this is empty, individual border Cells may have their own Style.
func (b *Border) BackgroundColor() types.Color {
	return b.bgColor
}

// SetBackgroundColor sets the background color for all the Border's Cells.
func (b *Border) SetBackgroundColor(color types.Color) {
	b.bgColor = color
}

// WithBackgroundColor sets the background color for all the Border's Cells and
// returns the Border.
func (b *Border) WithBackgroundColor(color types.Color) types.Border {
	b.bgColor = color
	return b
}

// HorizontalSpace returns the number of cells the supplied Border
// consumes.
func (b *Border) HorizontalSpace() types.Dimension {
	if b == nil || b.Empty() {
		return types.Dimension(0)
	}
	ls := 0
	if b.l != nil {
		ls = ansi.StringWidth(b.L().Content())
	}
	tls := 0
	if b.tl != nil {
		tls = ansi.StringWidth(b.TL().Content())
	}
	bls := 0
	if b.bl != nil {
		bls = ansi.StringWidth(b.BL().Content())
	}
	rs := 0
	if b.r != nil {
		rs = ansi.StringWidth(b.L().Content())
	}
	trs := 0
	if b.tr != nil {
		trs = ansi.StringWidth(b.TL().Content())
	}
	brs := 0
	if b.br != nil {
		brs = ansi.StringWidth(b.BL().Content())
	}
	return types.Dimension(max(ls, tls, bls) + max(rs, trs, brs))
}

// VerticalSpace returns the number of lines the supplied Border
// consumes.
func (b *Border) VerticalSpace() types.Dimension {
	if b == nil || b.Empty() {
		return types.Dimension(0)
	}
	ts := 0
	if b.t != nil {
		ts = ansi.StringWidth(b.T().Content())
	}
	tls := 0
	if b.tl != nil {
		tls = ansi.StringWidth(b.TL().Content())
	}
	trs := 0
	if b.tr != nil {
		trs = ansi.StringWidth(b.TL().Content())
	}
	bs := 0
	if b.b != nil {
		bs = ansi.StringWidth(b.B().Content())
	}
	bls := 0
	if b.bl != nil {
		bls = ansi.StringWidth(b.BL().Content())
	}
	brs := 0
	if b.br != nil {
		brs = ansi.StringWidth(b.BL().Content())
	}
	return types.Dimension(max(ts, tls, trs) + max(bs, bls, brs))
}

var _ types.Border = (*Border)(nil)
