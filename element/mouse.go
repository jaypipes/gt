package element

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// Click executes any OnClick callbacks that were registered for the
// Clickable.
func (e *Element) Click(ctx context.Context, ev types.MouseClickEvent) {
	for _, cb := range e.onClick {
		cb(ctx, ev)
	}
}

// OnClick registers a callback that will be executed when the Clickable is
// clicked.
func (e *Element) OnClick(cb types.ClickCallback) {
	e.onClick = append(e.onClick, cb)
}
