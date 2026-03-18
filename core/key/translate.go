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
		return KeyBacktab
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
		npkc, ok := nonPrintableStringToKeyCode[finalKey]
		if ok {
			return &Key{code: npkc}
		}
		if len(finalKey) != 1 {
			// we only support a single character as the final key if it's a
			// printable key.
			return &Key{}
		}
		mods := types.KeyModifiers(0)
		for _, modKey := range parts[0 : numParts-1] {
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
		npkc, ok := nonPrintableStringToKeyCode[finalKey]
		if ok {
			return &Key{code: npkc}
		}
		if len(finalKey) != 1 {
			// we only support a single character as the final key if it's a
			// printable key.
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
	case subject >= tcell.KeyCtrlA && subject <= tcell.KeyCtrlZ:
		// Unfortunately, tcell.KeyCtrlA through tcell.KeyCtrlZ's integer
		// values map directly to uppercase 'A' through uppercase 'Z' ASCII
		// characters. When a user calls gt.NewKey(tcell.KeyCtrlC), we set
		// set the Ctrl modifier bit manually since we don't have the
		// tcell.EventKey to give us the modifiers. Here, we return the key
		// code for the *lowercased* ASCII character, so Ctrl+C becomes
		// "ctrl+'c'".
		return types.KeyCode(subject + 32) // + 32 converts upper to lower
	case subject == tcell.KeyRune:
		// Special. We handle this in gt.NewKey() only by processing the
		// tcell.EventKey.
		return types.KeyCode(0)
	default:
		return tcellKeyToNonPrintableKeyCode[subject]
	}
}
