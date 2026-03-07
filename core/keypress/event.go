package keypress

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gdamore/tcell/v3"
	"github.com/samber/lo"

	"github.com/jaypipes/gt/core/event"
	"github.com/jaypipes/gt/types"
)

// Event exposes an easy-to-use interface for handling keypress events.
// Implements [types.KeyPressEvent].
type Event struct {
	event.Event
	// modifiers are the modifier keys that were held down
	modifiers types.KeyModifiers
	// key is the key code that was pressed.
	key tcell.Key
}

// String returns a simple string representation of the event.
func (e *Event) String() string {
	mods := e.modifiers.String()
	key := e.Printable()
	return fmt.Sprintf(
		"keypress:%s%s",
		key, mods,
	)
}

// Printable returns the printable character(s) that were pressed.
func (e *Event) Printable() string {
	return keyCodeToString(e.key)
}

// Shift returns true if the Shift modifier key was held.
func (e *Event) Shift() bool {
	return tcell.ModMask(e.modifiers)&tcell.ModShift != 0
}

// Ctrl returns true if the Ctrl modifier key was held.
func (e *Event) Ctrl() bool {
	return tcell.ModMask(e.modifiers)&tcell.ModCtrl != 0
}

// Alt returns true if the Alt modifier key was held.
func (e *Event) Alt() bool {
	return tcell.ModMask(e.modifiers)&tcell.ModAlt != 0
}

// Key returns the underlying key code that was pressed.
func (e *Event) Key() tcell.Key {
	return e.key
}

// MatchAny returns true if the event matches for any of the supplied
// keypress combination strings or key codes.
func (e *Event) MatchAny(subjects ...any) bool {
	for _, subject := range subjects {
		switch subject := subject.(type) {
		case string:
			parts := strings.Split(strings.ToLower(subject), "+")
			numParts := len(parts)
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
					if tcell.ModMask(e.modifiers)&tcell.ModMask(modIndex+1) == 0 {
						return false
					}
				}
				return matchesKey(finalKey, e.key)
			case numParts == 1:
				key := parts[0]
				if lo.Contains(types.KeyModifierStrings, key) {
					// finalKey cannot be a modifier... ignore and return false.
					return false
				}
				return matchesKey(key, e.key)
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

func matchesKey(key string, code tcell.Key) bool {
	// quick lookup on simple printable runes...
	if len(key) == 1 {
		r := rune(key[0])
		if strconv.IsPrint(r) {
			if int16(r) == int16(code) {
				return true
			}
		}
	}
	return false
}

func keyCodeToString(code tcell.Key) string {
	named, ok := tcell.KeyNames[code]
	if ok {
		return strings.ToLower(named)
	}
	return strconv.QuoteRune(rune(code))
}

// EventFromTCell returns an Event from a [tcell.EventKey]
func EventFromTCell(
	te *tcell.EventKey,
) *Event {
	mods := te.Modifiers()
	e := &Event{
		Event:     event.New(),
		modifiers: types.KeyModifiers(mods),
		key:       te.Key(),
	}
	e.SetWhen(te.When())

	return e
}

var _ tcell.Event = (*Event)(nil)
var _ types.KeyPressEvent = (*Event)(nil)
