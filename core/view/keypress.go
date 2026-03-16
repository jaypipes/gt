package view

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// handleKeyPressEvent passes a KeyPressEvent to any handlers that are
// listening for KeyPressEvents.
func (v *View) handleKeyPressEvent(
	ctx context.Context,
	ev types.KeyPressEvent,
) bool {
	return v.KeyPress(ctx, ev)
}
