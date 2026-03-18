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
// on the other hand, has individual values for all printable ASCII characters
// and non-printable keys (e.g. "F1" or "PrintScreen") are assigned a Unicode
// code point in the Basic Multilingual Plane's Private Use Area (U+E000
// through U+F8FF).
//
// The `gt.core.key.New()` function can translate a tcell.Key, a string or a
// *tcell.EventKey object into a properly-formed gt.Key object with a
// gt.KeyCode exposed via the gt.Key.Code() method.
type KeyCode int32

// Printable returns whether the KeyCode is a printable key.
func (c KeyCode) Printable() bool {
	return c < KeyCodeNonPrintableStart
}

// Key describes a key press combination.
type Key interface {
	fmt.Stringer
	KeyModifiable
	// Code returns the Unicode code point for the Key.
	Code() KeyCode
	// Printable returns true if the Key can be directly printed to the Screen.
	Printable() bool
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

const (
	KeyCodeNUL KeyCode = 0
)

// Non-printable keys (e.g. "F1" or "Print") are assigned a Unicode code point
// in the Basic Multilingual Plane's Private Use Area (U+E000 through U+F8FF).
// That way, we can do simple int32 comparisons and use the String() method to
// return a printable representation of the pressed key combination.
const (
	KeyCodeNonPrintableStart KeyCode = iota + 57344 // 57344 == U+E000
	KeyCodeSOH
	KeyCodeSTX
	KeyCodeETX
	KeyCodeEOT
	KeyCodeENQ
	KeyCodeACK
	KeyCodeBEL
	KeyCodeVT
	KeyCodeFF
	KeyCodeSO
	KeyCodeSI
	KeyCodeDLE
	KeyCodeDC1
	KeyCodeDC2
	KeyCodeDC3
	KeyCodeDC4
	KeyCodeNAK
	KeyCodeSYN
	KeyCodeETB
	KeyCodeCAN
	KeyCodeEM
	KeyCodeSUB
	KeyCodeFS
	KeyCodeGS
	KeyCodeRS
	KeyCodeUS
	KeyCodeBackspace // Backward delete
	KeyCodeDelete    // Forward delete
	KeyCodeTab
	KeyCodeEscape
	KeyCodeEnter
	KeyCodeUp
	KeyCodeDown
	KeyCodeRight
	KeyCodeLeft
	KeyCodeUpLeft
	KeyCodeUpRight
	KeyCodeDownLeft
	KeyCodeDownRight
	KeyCodeCenter
	KeyCodePgUp
	KeyCodePgDn
	KeyCodeHome
	KeyCodeEnd
	KeyCodeInsert
	KeyCodeHelp
	KeyCodeExit
	KeyCodeClear
	KeyCodeCancel
	KeyCodePrint
	KeyCodePause
	KeyCodeBacktab // shift+tab
	KeyCodeMenu
	KeyCodeCapsLock
	KeyCodeScrollLock
	KeyCodeNumLock
	KeyCodeF1
	KeyCodeF2
	KeyCodeF3
	KeyCodeF4
	KeyCodeF5
	KeyCodeF6
	KeyCodeF7
	KeyCodeF8
	KeyCodeF9
	KeyCodeF10
	KeyCodeF11
	KeyCodeF12
	KeyCodeF13
	KeyCodeF14
	KeyCodeF15
	KeyCodeF16
	KeyCodeF17
	KeyCodeF18
	KeyCodeF19
	KeyCodeF20
	KeyCodeF21
	KeyCodeF22
	KeyCodeF23
	KeyCodeF24
	KeyCodeF25
	KeyCodeF26
	KeyCodeF27
	KeyCodeF28
	KeyCodeF29
	KeyCodeF30
	KeyCodeF31
	KeyCodeF32
	KeyCodeF33
	KeyCodeF34
	KeyCodeF35
	KeyCodeF36
	KeyCodeF37
	KeyCodeF38
	KeyCodeF39
	KeyCodeF40
	KeyCodeF41
	KeyCodeF42
	KeyCodeF43
	KeyCodeF44
	KeyCodeF45
	KeyCodeF46
	KeyCodeF47
	KeyCodeF48
	KeyCodeF49
	KeyCodeF50
	KeyCodeF51
	KeyCodeF52
	KeyCodeF53
	KeyCodeF54
	KeyCodeF55
	KeyCodeF56
	KeyCodeF57
	KeyCodeF58
	KeyCodeF59
	KeyCodeF60
	KeyCodeF61
	KeyCodeF62
	KeyCodeF63
	KeyCodeF64

	KeyCodeNonPrintableEnd KeyCode = 63743 // 56374 == U+F8FF
)
