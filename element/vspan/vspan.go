package vspan

import (
	"context"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.vspan"
)

// New returns a new VSpan instance containing the supplied raw string content.
func New(
	ctx context.Context,
	content string,
) *VSpan {
	e := element.New(ctx, ElementClass)
	v := &VSpan{Element: e}
	v.SetDisplay(types.DisplayInlineBlock)
	v.SetHeight(core.Percent(100))
	v.SetTextContent(content)
	return v
}

// VSpan is an Element that uses the inline-block display mode by default and
// defaults to the full height of any parent container.
//
// It's essentially a short-cut for creating a Span and calling
// SetHeight(Percent(100)) on it.
type VSpan struct {
	element.Element
}
