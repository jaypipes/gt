package types

// Whitespace is a mode for handling wrapping and whitespaces.
type Whitespace uint8

const (
	// WhitespaceNormal indicates sequences of whitespace will collapse into a
	// single whitespace. Text will wrap when necessary. This is the default.
	WhitespaceNormal Whitespace = 0
	// WhitespacePreserve indicates whitespace is preserved when rendered.
	WhitespacePreserve = 1
	// WhitespaceWrapNever indicates text will never wrap to the next line.
	// Text continues on the same line until the container's right margin and
	// will then be clipped.
	WhitespaceWrapNever = 1 << 1
	// WhitespaceWrapLine indicates text will only wrap on line breaks (i.e. \n
	// or \r\n)
	WhitespaceWrapLine = 1 << 2
)
