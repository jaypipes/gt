package keyshortcut

import "github.com/jaypipes/gt/types"

// KeyShortcut describes an action taken when a specific Key is pressed by the
// user.
type KeyShortcut struct {
	// key is the specific keypress combination that will fire the shortcut
	// action.
	key types.Key
	// callback isthe KeyShortcutCallback that will be executed when the
	// shortcut's keypress combination is pressed by the user.
	callback types.KeyShortcutCallback
}

// Key returns the specific keypress combination that will fire the
// shortcut action.
func (k *KeyShortcut) Key() types.Key {
	return k.key
}

// SetKey sets the specific keypress combination that will fire the
// shortcut action.
func (k *KeyShortcut) SetKey(key types.Key) {
	k.key = key
}

// Callback returns the KeyShortcutCallback that will be executed when the
// shortcut's keypress combination is pressed by the user.
func (k *KeyShortcut) Callback() types.KeyShortcutCallback {
	return k.callback
}

// SetCallback sets the KeyShortcutCallback that will be executed when the
// shortcut's keypress combination is pressed by the user.
func (k *KeyShortcut) SetCallback(cb types.KeyShortcutCallback) {
	k.callback = cb
}
