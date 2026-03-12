package keypress

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v3"
	"github.com/samber/lo"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// Event exposes an easy-to-use interface for handling keypress events.
// Implements [types.KeyPressEvent].
type Event struct {
	event.Event
	core.KeyModifiable
	// key is the key code that was pressed.
	key tcell.Key
	// str is the string representation of the pressed key.
	str string
}

// String returns a simple string representation of the event.
func (e *Event) String() string {
	mods := e.KeyModifiers().String()
	key := e.Printable()
	return fmt.Sprintf(
		"keypress:%s%s",
		key, mods,
	)
}

// Printable returns the printable character(s) that were pressed.
func (e *Event) Printable() string {
	if e.key == tcell.KeyRune {
		return e.str
	}
	return keyCodeToString(e.key)
}

// SetKey sets the underlying key code that was pressed.
func (e *Event) SetKey(key tcell.Key) {
	e.key = key
}

// Key returns the underlying key code that was pressed.
func (e *Event) Key() tcell.Key {
	return e.key
}

// SetStr sets the string representation of the key that was pressed. Only
// applicable when Key() == tcell.KeyRune.
func (e *Event) SetStr(str string) {
	e.str = str
}

// Str returns the string representation of the key that was pressed. Only
// applicable when Key() == tcell.KeyRune.
func (e *Event) Str() string {
	return e.str
}

// MatchAny returns true if the event matches for any of the supplied
// keypress combination strings or key codes.
func (e *Event) MatchAny(subjects ...any) bool {
	for _, subject := range subjects {
		switch subject := subject.(type) {
		case string:
			parts := strings.Split(strings.ToLower(subject), "+")
			numParts := len(parts)
			mods := tcell.ModMask(e.KeyModifiers())
			switch {
			case numParts > 1:
				// modifiers were passed, e.g. "ctrl+alt+p"
				finalKey := parts[numParts-1]
				if lo.Contains(types.KeyModifierStrings, finalKey) {
					// finalKey cannot be a modifier... ignore and return false.
					return false
				}
				for _, modKey := range parts[0 : numParts-2] {
					modIndex := lo.IndexOf(types.KeyModifierStrings, modKey)
					if modIndex == -1 {
						// modKey must be a modifier... ignore and return false.
						return false
					}
					if mods&tcell.ModMask(modIndex+1) == 0 {
						return false
					}
				}
				return e.matchesKey(finalKey)
			case numParts == 1:
				key := parts[0]
				if lo.Contains(types.KeyModifierStrings, key) {
					// finalKey cannot be a modifier... ignore and return false.
					return false
				}
				return e.matchesKey(key)
			default:
				// empty string...
				return false
			}
		case tcell.Key:

			return false
		}
	}
	return false
}

func (e *Event) matchesKey(key string) bool {
	return strings.EqualFold(key, e.Printable())
}

func keyCodeToString(code tcell.Key) string {
	named, ok := tcell.KeyNames[code]
	if ok {
		return strings.ToLower(named)
	}
	return strconv.QuoteRune(rune(code))
}

var _ tcell.Event = (*Event)(nil)
var _ types.KeyPressEvent = (*Event)(nil)
