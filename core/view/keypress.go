package view

import (
	"context"

	"github.com/jaypipes/gt/core/key"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/types"
)

// ActiveKey returns the key press combination that should make this View the
// active (displayed) View.
func (v *View) ActiveKey() types.Key {
	return v.activeKey
}

// SetActiveKey sets the key combination that should trigger setting this View
// as the active View in the Application.
//
// The keypress combination can be a string ("Tab"), a [tcell.Key] code
// (tcell.KeyTab), a [types.KeyCode] value (types.KeyCodeTab) or a [types.Key]
// object (core.key.KeyTab)
func (v *View) SetActiveKey(subject any) {
	v.activeKey = key.New(subject)
}

// WithActiveKey sets the key combination that should trigger setting this View
// as the active View in the Application and returns the View.
//
// The keypress combination can be a string ("Tab"), a [tcell.Key] code
// (tcell.KeyTab), a [types.KeyCode] value (types.KeyCodeTab) or a [types.Key]
// object (core.key.KeyTab)
func (v *View) WithActiveKey(subject any) types.View {
	v.SetActiveKey(subject)
	return v
}

// SetKeyShortcut registers a View-level KeyShortcut that will execute upon a
// key press combination.
func (v *View) SetKeyShortcut(shortcut types.KeyShortcut) {
	k := shortcut.Key()
	for _, ks := range v.keyShortcuts {
		ksk := ks.Key()
		if ksk.Equal(k) {
			gtlog.Warn(
				context.TODO(),
				"key shortcut %q shadows previously-registered "+
					"view-level key shortcut",
				k,
			)
		}
	}
	v.keyShortcuts = append(v.keyShortcuts, shortcut)
}

// KeyPress checks for any KeyShortcuts that are registered withe View and
// executes any matched callback. If no KeyShortcuts are matched, we execute
// the View's internal vdiv Element's KeyPress method.
func (v *View) KeyPress(ctx context.Context, ev types.KeyPressEvent) bool {
	k := ev.Key()
	for _, ks := range v.keyShortcuts {
		ksk := ks.Key()
		if ksk.Equal(k) {
			cb := ks.Callback()
			cb(ctx)
			return true
		}
	}
	return v.VDiv.KeyPress(ctx, ev)
}
