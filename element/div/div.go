package div

import (
	"context"

	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.div"
)

// New returns a new Div instance with the given options.
func New(
	ctx context.Context,
	opts ...types.ElementWithOption,
) *Div {
	e := element.New(ctx, ElementClass)
	d := &Div{Element: e}
	d.SetDisplay(types.DisplayBlock)
	for _, opt := range opts {
		opt(d)
	}
	return d
}

// Div is an Element that uses the block display mode by default.
type Div struct {
	element.Element
}
