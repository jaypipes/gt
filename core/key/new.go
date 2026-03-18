package key

import (
	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/types"
)

// New returns a new [types.Key] from either a string, [types.KeyCode],
// [tcell.Key] or [tcell.EventKey]
func New(subject any) *Key {
	switch subject := subject.(type) {
	case types.KeyCode:
		return &Key{
			code: subject,
		}
	case string:
		return FromString(subject)
	case tcell.Key:
		kc := keyCodeFromTCellKey(subject)
		// Unfortunately, tcell.KeyCtrlA through tcell.KeyCtrlZ's integer
		// values map directly to uppercase 'A' through uppercase 'Z' ASCII
		// characters. When a user calls gt.NewKey(tcell.KeyCtrlC), we need to
		// set the Ctrl modifier bit manually since we don't have the
		// tcell.EventKey to give us the modifiers.
		mods := types.KeyModifiers(0)
		if subject >= tcell.KeyCtrlA && subject <= tcell.KeyCtrlZ {
			mods |= types.KeyModifierCtrl
		}
		k := &Key{
			code: kc,
		}
		k.SetModifiers(mods)
		return k
	case *tcell.EventKey:
		var code types.KeyCode
		tk := subject.Key()
		if tk == tcell.KeyRune {
			s := subject.Str()
			if len(s) == 1 {
				code = types.KeyCode([]rune(s)[0])
			}
		} else {
			code = keyCodeFromTCellKey(tk)
		}
		k := &Key{
			code: code,
		}
		mods := subject.Modifiers()
		k.SetModifiers(types.KeyModifiers(mods))
		return k
	default:
		return &Key{}
	}
}
