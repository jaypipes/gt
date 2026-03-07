package types

// Cell represents a single cell on a [Screen].
type Cell interface {
	Style
	// Style returns the Cell's Style
	Style() Style
	// SetStyle sets the Cell's style.
	SetStyle(Style)
	// WithStyle sets the Cell's style and returns the Cell.
	WithStyle(Style) Cell
	// Empty returns true if the Cell has no content.
	Empty() bool
	// Content returns the Cell's string content.
	Content() string
	// SetContent sets the Cell's string content.
	SetContent(string)
}

// CellWithOption describes an optional varg parameter to [style.New] that
// modifies the returned Cell.
type CellWithOption func(Cell)
