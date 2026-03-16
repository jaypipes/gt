package key

import (
	"github.com/gdamore/tcell/v3"

	"github.com/jaypipes/gt/types"
)

// New returns a new [types.Key] from either a string, [tcell.Key] or
// [tcell.EventKey]
func New(subject any) *Key {
	switch subject := subject.(type) {
	case string:
		return FromString(subject)
	case tcell.Key:
		return &Key{
			code: keyCodeFromTCellKey(subject),
		}
	case *tcell.EventKey:
		var code types.KeyCode
		tk := subject.Key()
		if tk < tcell.KeyRune {
			code = types.KeyCode(tk)
		} else {
			s := subject.Str()
			if len(s) == 1 {
				code = types.KeyCode([]rune(s)[0])
			}
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
