package key

import (
	"fmt"
	"strconv"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/types"
)

// Key describes a key press combination.
type Key struct {
	fmt.Stringer
	core.KeyModifiable
	// code is the types.KeyCode representing the key pressed.
	code types.KeyCode
}

// String returns a string representation of the Key.
func (k *Key) String() string {
	mods := k.Modifiers().String()
	if len(mods) > 0 {
		mods += "+"
	}
	if !k.code.Printable() {
		return nonPrintableKeyCodeToString[k.code]
	}
	return mods + strconv.QuoteRune(rune(k.code))
}

// Code returns the types.KeyCode for the Key.
func (k *Key) Code() types.KeyCode {
	return k.code
}

// Printable returns whether the Key can be directly printed to the Screen.
func (k *Key) Printable() bool {
	return k.code.Printable()
}

// Equal returns true if the Key matches the other Key.
func (k *Key) Equal(other types.Key) bool {
	return k.Modifiers() == other.Modifiers() &&
		k.code == other.Code()
}
