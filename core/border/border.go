package border

import (
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

// T returns the Cell used for the top edge of the Border.
func (b *Border) T() types.Cell {
	return b.t
}

// SetT sets the Cell to be used for the top edge of the Border.
func (b *Border) SetT(cell types.Cell) {
	b.t = cell
}

// B returns the Cell used for the bottom edge of the Border.
func (b *Border) B() types.Cell {
	return b.b
}

// SetB sets the Cell to be used for the bottom edge of the Border.
func (b *Border) SetB(cell types.Cell) {
	b.b = cell
}

// L returns the Cell used for the left edge of the Border.
func (b *Border) L() types.Cell {
	return b.l
}

// SetT sets the Cell to be used for the left edge of the Border.
func (b *Border) SetL(cell types.Cell) {
	b.l = cell
}

// R returns the Cell used for the right edge of the Border.
func (b *Border) R() types.Cell {
	return b.r
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

// BackgroundColor returns the background color used for all the Border's
// Cells. If this is empty, individual border Cells may have their own Style.
func (b *Border) BackgroundColor() types.Color {
	return b.bgColor
}

// SetBackgroundColor sets the background color for all the Border's Cells.
func (b *Border) SetBackgroundColor(color types.Color) {
	b.bgColor = color
}

var _ types.Border = (*Border)(nil)
