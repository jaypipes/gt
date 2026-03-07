package types

// Bqrder describes the decoration around the outside of a bounding box.
type Border interface {
	// T returns the Cell used for the top edge of the Border.
	T() Cell
	// SetT sets the Cell to be used for the top edge of the Border.
	SetT(Cell)
	// B returns the Cell used for the bottom edge of the Border.
	B() Cell
	// SetB sets the Cell to be used for the bottom edge of the Border.
	SetB(Cell)
	// L returns the Cell used for the left edge of the Border.
	L() Cell
	// SetL sets the Cell to be used for the left edge of the Border.
	SetL(Cell)
	// R returns the Cell used for the right edge of the Border.
	R() Cell
	// SetR sets the Cell to be used for the right edge of the Border.
	SetR(Cell)
	// TL returns the Cell used for the top left corner of the Border.
	TL() Cell
	// SetTL sets the Cell to be used for the top left corner of the Border.
	SetTL(Cell)
	// TR returns the Cell used for the top right corner of the Border.
	TR() Cell
	// SetTR sets the Cell to be used for the top right corner of the Border.
	SetTR(Cell)
	// BL returns the Cell used for the bottom left corner of the Border.
	BL() Cell
	// SetBL sets the Cell to be used for the bottom left corner of the Border.
	SetBL(Cell)
	// BR returns the Cell used for the bottom right corner of the Border.
	BR() Cell
	// SetBR sets the Cell to be used for the bottom right corner of the
	// Border.
	SetBR(Cell)
	// ForegroundColor returns the foreground color used for all the Border's
	// Cells. If this is empty, individual border Cells may have their own
	// Style.
	ForegroundColor() Color
	// SetForegroundColor sets the foreground color for all the Border's Cells.
	SetForegroundColor(Color)
	// BackgroundColor returns the background color used for all the Border's
	// Cells. If this is empty, individual border Cells may have their own
	// Style.
	BackgroundColor() Color
	// SetBackgroundColor sets the background color for all the Border's Cells.
	SetBackgroundColor(Color)
}
