package types

// Bqrder describes the decoration around the outside of a bounding box.
type Border interface {
	// Empty returns true if none of the Border's edges or corners have
	// content.
	Empty() bool
	// T returns the Cell used for the top edge of the Border.
	T() Cell
	// TSize returns the number of lines the top edge of the Border will
	// consume.
	TSize() Dimension
	// SetT sets the Cell to be used for the top edge of the Border.
	SetT(Cell)
	// B returns the Cell used for the bottom edge of the Border.
	B() Cell
	// BSize returns the number of lines the bottom edge of the Border will
	// consume.
	BSize() Dimension
	// SetB sets the Cell to be used for the bottom edge of the Border.
	SetB(Cell)
	// L returns the Cell used for the left edge of the Border.
	L() Cell
	// LSize returns the number of cells the left edge of the Border will
	// consume.
	LSize() Dimension
	// SetL sets the Cell to be used for the left edge of the Border.
	SetL(Cell)
	// R returns the Cell used for the right edge of the Border.
	R() Cell
	// RSize returns the number of cells the right edge of the Border will
	// consume.
	RSize() Dimension
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
	// HorizontalSpace returns the number of cells the Border consumes.
	HorizontalSpace() Dimension
	// VerticalSpace returns the number of lines the Border consumes.
	VerticalSpace() Dimension
}

// BorderWithOption describes an optional varg parameter to [border.New] that
// modifies the returned Border.
type BorderWithOption func(Border)
