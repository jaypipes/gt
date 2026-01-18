package types

// Contented describes something that can have some raw string contents.
type Contented interface {
	// SetContent sets the Contented's raw string contents.
	SetContent(string)
	// Content returns the Contented's raw string contents.
	Content() string
	// ContentWidth returns the width of the Contented's raw string contents.
	ContentWidth() Dimension
	// ContentHeight returns the height of the Contented's raw string contents.
	ContentHeight() Dimension
}
