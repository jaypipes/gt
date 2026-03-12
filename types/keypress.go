package types

// KeyPressMap maps key press combination strings to callbacks that will
// execute upon that key press.
type KeyPressMap map[string]EventCallback

// HasKeyPressMap describes something that has a map of key press combinations
// and callbacks associated with it.
type HasKeyPressMap interface {
	// KeyPressMap returns the KeyPressable's map of key press combination
	// strings to callbacks that will execute when that key press combination
	// is entered.
	KeyPressMap() KeyPressMap
}

// KeyPressEvent describes events received when a key press occurs.
type KeyPressEvent interface {
	Event
	KeyModifiable
	// Key returns the virtual key code
	Key() Key
	// SetKey sets the virtual key code
	SetKey(Key)
	// SetStr sets the string representation of the key that was pressed. Only
	// applicable when Key() == tcell.KeyRune.
	SetStr(string)
	// Str returns the string representation of the key that was pressed. Only
	// applicable when Key() == tcell.KeyRune.
	Str() string
	// Printable returns the printable character(s) associated with the key
	// press event.
	Printable() string
	// MatchAny returns true if the KeyPressEvent matches for any of the
	// keypress strings or key codes supplied.
	MatchAny(...any) bool
}

// KeyPressEventWithOption describes an optional varg parameter to
// [core.event.keypress.New] that modifies the returned KeyPressEvent.
type KeyPressEventWithOption func(KeyPressEvent)
