package types

// Display describes the behaviour of a rendering box.
type Display int

const (
	// DisplayInline means the rendered box will not begin on a new line and
	// any width and height will have no effect. This is the default display
	// mode.
	DisplayInline Display = iota
	// DisplayBlock means the rendered box will begin on a new line. The
	// default width of the element will be all available remaining width in
	// the line. Any width and height that has been set on the element will be
	// used.
	DisplayBlock
	// DisplayInlineBlock means the rendered box will NOT begin on a new line,
	// however the default width of the element will be all available remaining
	// width in the line. Any width and height that has been set on the element
	// will be used.
	DisplayInlineBlock
)

var (
	displayStrings = map[Display]string{
		DisplayInline:      "inline",
		DisplayBlock:       "block",
		DisplayInlineBlock: "inline-block",
	}
)

func (d Display) String() string {
	return displayStrings[d]
}
