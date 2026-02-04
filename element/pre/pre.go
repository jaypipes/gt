package pre

import (
	"context"

	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.pre"
)

// New returns a new Pre instance with the given options.
func New(
	ctx context.Context,
	opts ...types.ElementWithOption,
) *Pre {
	e := element.New(ctx, ElementClass)
	s := &Pre{Element: e}
	s.SetDisplay(types.DisplayBlock)
	s.SetWhitespace(types.WhitespacePreserve)
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Pre is an Element that uses block display mode by default and uses a
// whitespace mode that preserves whitespace characters.
type Pre struct {
	element.Element
}
