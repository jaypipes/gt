package key

import (
	"strings"

	"github.com/gdamore/tcell/v3"
	"github.com/samber/lo"

	"github.com/jaypipes/gt/types"
)

var (
	tcellKeyNameToCode map[string]tcell.Key
)

func init() {
	// populate a reverse lookup of lowercased key string to key code
	tcellKeyNameToCode = lo.Invert(
		lo.MapValues(
			tcell.KeyNames,
			func(v string, _ tcell.Key) string {
				return strings.ToLower(v)
			},
		),
	)
}

// FromString returns a [types.Key] from a string.
func FromString(subject string) *Key {
	subject = strings.ToLower(subject)
	// handle some special exceptions first...
	if subject == "shift+tab" {
		return &Key{
			code: types.KeyCode(tcell.KeyBacktab),
		}
	}
	parts := strings.Split(subject, "+")
	numParts := len(parts)
	switch {
	case numParts > 1:
		// modifiers were passed, e.g. "ctrl+alt+p"
		finalKey := parts[numParts-1]
		if lo.Contains(keyModifierStrings, finalKey) {
			// finalKey cannot be a modifier... ignore and return NUL.
			return &Key{}
		}
		if len(finalKey) != 1 {
			// we only support a single character as the final key.
			return &Key{}
		}
		mods := types.KeyModifiers(0)
		for _, modKey := range parts[0 : numParts-2] {
			mod, ok := stringToKeyModifier[modKey]
			if !ok {
				// modKey must be a modifier... ignore and return NUL.
				return &Key{}
			}
			mods |= mod
		}
		k := &Key{
			code: types.KeyCode([]rune(finalKey)[0]),
		}
		k.SetModifiers(mods)
		return k
	case numParts == 1:
		finalKey := parts[0]
		if lo.Contains(keyModifierStrings, finalKey) {
			// finalKey cannot be a modifier... ignore and return NUL.
			return &Key{}
		}
		if len(finalKey) != 1 {
			// we only support a single character as the final key.
			return &Key{}
		}
		return &Key{
			code: types.KeyCode([]rune(finalKey)[0]),
		}
	default:
		// empty string...
		return &Key{}
	}
}

var (
	keyModifierStrings = []string{
		"shift", "ctrl", "alt",
	}
	stringToKeyModifier = map[string]types.KeyModifiers{
		"shift": types.KeyModifierShift,
		"ctrl":  types.KeyModifierCtrl,
		"alt":   types.KeyModifierAlt,
	}
)

// tcellCodeFromString returns the [tcell.Key] representing a single keystroke.
func tcellCodeFromString(subject string) tcell.Key {
	named, ok := tcellKeyNameToCode[subject]
	if ok {
		return named
	}
	return tcell.KeyRune
}

// keyCodeFromTCellKey returns the [types.KeyCode] from a [tcell.Key]. Note
// that this only works for non-printable or control-key combination
// [tcell.Key] values since tcell.KeyRune is used for all printable characters.
func keyCodeFromTCellKey(subject tcell.Key) types.KeyCode {
	switch {
	case subject < tcell.KeyRune:
		return types.KeyCode(subject)
	case subject > tcell.KeyRune && subject <= tcell.KeyNumLock:
		return tcellKeyToNonPrintableKeyCode[subject]
	}
	return types.KeyCode(0)
}
