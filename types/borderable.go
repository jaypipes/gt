package types

// Borderable describes something that can have a border.
type Borderable interface {
	// SetBorder sets the Borderable's Border.
	SetBorder(Border)
	// Bounds returns the Border for the Borderable.
	Border() Border
}
