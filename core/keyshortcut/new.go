package keyshortcut

import (
	"context"

	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a KeyShortcut.
//
// You can pass zero or more KeyShortcutWithOptions to optionally set certain
// attributes on the returned KeyShortcut.
func New(
	ctx context.Context,
	opts ...types.KeyShortcutWithOption,
) *KeyShortcut {
	k := &KeyShortcut{}
	for _, opt := range opts {
		opt(k)
	}
	return k
}

// WithKey sets the Key in the KeyShortcut.
func WithKey(key types.Key) types.KeyShortcutWithOption {
	return func(k types.KeyShortcut) {
		k.SetKey(key)
	}
}

// WithCallback sets the KeyShortcutCallback in the KeyShortcut.
func WithCallback(cb types.KeyShortcutCallback) types.KeyShortcutWithOption {
	return func(k types.KeyShortcut) {
		k.SetCallback(cb)
	}
}
