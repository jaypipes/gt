package types

// Whitespaced describes something that has a whitespace mode
type Whitespaced interface {
	// SetWhitespace sets the Whitespaced's whitespace mode
	SetWhitespace(Whitespace)
	// Whitespace returns the Whitespaced's whitespace mode
	Whitespace() Whitespace
}
