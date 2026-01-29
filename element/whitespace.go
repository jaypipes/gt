package element

import "github.com/jaypipes/gt/types"

// WithWhitespace sets the Element's whitespace mode and returns the Element.
func (e *Element) WithWhitespace(whitespace types.Whitespace) types.Element {
	e.SetWhitespace(whitespace)
	return e
}
