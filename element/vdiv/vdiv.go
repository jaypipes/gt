package vdiv

import (
	"context"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/types"
)

const (
	ElementClass = "gt.vdiv"
)

// New returns a new VDiv instance containing the supplied raw string content.
func New(
	ctx context.Context,
	content string,
) *VDiv {
	e := element.New(ctx, ElementClass)
	v := &VDiv{Element: e}
	v.SetDisplay(types.DisplayBlock)
	v.SetHeight(core.Percent(100))
	v.SetTextContent(content)
	return v
}

// VDiv is an Element that uses the block display mode by default and defaults
// to the full height of its container.
//
// It's essentially a short-cut for creating a Div and calling
// SetHeight(Percent(100)) on it.
type VDiv struct {
	element.Element
}
