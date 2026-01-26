package div

import (
	"context"

	"github.com/jaypipes/gt/element/base"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.div"
)

// New returns a new Div instance containing the supplied raw string content.
func New(
	ctx context.Context,
	content string,
) *Div {
	b := base.New(ctx, ElementClass)
	d := &Div{Base: b}
	d.SetDisplay(types.DisplayBlock)
	d.SetTextContent(content)
	return d
}

// Div is an Element that uses the block display mode by default.
type Div struct {
	base.Base
}
