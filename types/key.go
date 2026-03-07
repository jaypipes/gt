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
