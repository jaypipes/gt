package view

import (
	"context"

	"github.com/jaypipes/gt/element/vdiv"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a View.
//
// You can pass zero or more ViewWithOptions to optionally set certain
// attributes on the returned View.
func New(
	ctx context.Context,
	opts ...types.ViewWithOption,
) *View {
	d := vdiv.New(ctx)
	v := &View{
		VDiv: *d,
	}
	for _, opt := range opts {
		opt(v)
	}
	return v
}

// WithID sets the types.View's ID to the supplied value.
func WithID(id string) types.ViewWithOption {
	return func(v types.View) {
		v.SetID(id)
	}
}
