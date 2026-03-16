package types

import (
	"fmt"
)

// KeyCode represents the Unicode code point of the pressed key (rune is an
// int32 which is the Go language's Unicode code point type)
//
// Note that KeyCode is *NOT* equivalent to tcell.Key. tcell.Key values are a
// custom int16 code for common non-printable characters, with tcell.KeyRune
// (int16 value of 256) representing all printable ASCII characters. KeyCode,
// on the other hand, has individual values for all printable ASCII
// characters.
//
// The `gt.core.key.New()` function can translate a tcell.Key, a string or a
// *tcell.EventKey object into a properly-formed gt.Key object with a
// gt.KeyCode exposed via the gt.Key.Code() method.
type KeyCode int32

// Key describes a key press combination.
type Key interface {
	fmt.Stringer
	KeyModifiable
	// Code returns the Unicode code point for the key.
	Code() KeyCode
	// Equal returns true if the Key matches the supplied other Key.
	Equal(Key) bool
}

// KeyMap maps key press combination strings to callbacks that will
// execute upon that key press.
type KeyMap map[Key]EventCallback

// HasKeyMap describes something that has a map of key press combinations
// and callbacks associated with it.
type HasKeyMap interface {
	// KeyMap returns the Keyable's map of key press combination
	// strings to callbacks that will execute when that key press combination
	// is entered.
	KeyMap() KeyMap
}
