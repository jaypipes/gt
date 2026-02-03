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

// New returns a new VDiv instance with the given options.
func New(
	ctx context.Context,
	opts ...types.ElementWithOption,
) *VDiv {
	e := element.New(ctx, ElementClass)
	v := &VDiv{Element: e}
	v.SetDisplay(types.DisplayBlock)
	v.SetHeight(core.Percent(100))
	for _, opt := range opts {
		opt(v)
	}
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
