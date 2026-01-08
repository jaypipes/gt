package types

// Styled describes something that can have a Style applied to it
type Styled interface {
	// Style returns the thing's Style
	Style() Style
	// SetStyle applies the supplied Style to the Styled.
	SetStyle(Style)
}
