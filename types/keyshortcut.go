package types

import "context"

// KeyShortcut describes an action taken when a specific Key is pressed by the
// user.
type KeyShortcut interface {
	// Key returns the specific keypress combination that will fire the
	// shortcut action.
	Key() Key
	// SetKey sets the specific keypress combination that will fire the
	// shortcut action.
	SetKey(Key)
	// Callback returns the KeyShortcutCallback that will be executed when the
	// shortcut's keypress combination is pressed by the user.
	Callback() KeyShortcutCallback
	// SetCallback sets the keyShortcutCallback that will be executed when the
	// shortcut's keypress combination is pressed by the user.
	SetCallback(KeyShortcutCallback)
}

// KeyShortcutCallback is the action that will be taken when a keypress
// combination matching the KeyShortcut's Key is pressed by the user.
type KeyShortcutCallback func(context.Context)

// KeyShortcutWithOption describes an optional varg parameter to
// [core.keyshortcut.New] that modifies the returned KeyShortcut.
type KeyShortcutWithOption func(KeyShortcut)

// KeyShortcutHandler describes a thing that can register and execute
// KeyShortcuts.
type KeyShortcutHandler interface {
	// SetKeyShortcut registers a KeyShortcut for the KeyShortcutHandler. If
	// the KeyShortcutHandler already has a KeyShortcut for the registered
	// KeyShortcut's Key, a warning will be sent to the gt log that the new
	// KeyShortcut shadows a previously-registered KeyShortcut.
	SetKeyShortcut(KeyShortcut)
}
