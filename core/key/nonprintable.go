package key

import (
	"github.com/gdamore/tcell/v3"
	"github.com/jaypipes/gt/types"
)

// Non-printable keys (e.g. "F1" or "Print") are assigned a Unicode code point
// in the Basic Multilingual Plane's Private Use Area (U+E000 through U+F8FF).
// That way, we can do simple int32 comparisons and use the String() method to
// return a printable representation of the pressed key combination.
var (
	KeyUp         = &Key{code: types.KeyCodeUp}
	KeyDown       = &Key{code: types.KeyCodeDown}
	KeyRight      = &Key{code: types.KeyCodeRight}
	KeyLeft       = &Key{code: types.KeyCodeLeft}
	KeyUpLeft     = &Key{code: types.KeyCodeUpLeft}
	KeyUpRight    = &Key{code: types.KeyCodeUpRight}
	KeyDownLeft   = &Key{code: types.KeyCodeDownLeft}
	KeyDownRight  = &Key{code: types.KeyCodeDownRight}
	KeyCenter     = &Key{code: types.KeyCodeCenter}
	KeyPgUp       = &Key{code: types.KeyCodePgUp}
	KeyPgDn       = &Key{code: types.KeyCodePgDn}
	KeyHome       = &Key{code: types.KeyCodeHome}
	KeyEnd        = &Key{code: types.KeyCodeEnd}
	KeyInsert     = &Key{code: types.KeyCodeInsert}
	KeyDelete     = &Key{code: types.KeyCodeDelete}
	KeyHelp       = &Key{code: types.KeyCodeHelp}
	KeyExit       = &Key{code: types.KeyCodeExit}
	KeyClear      = &Key{code: types.KeyCodeClear}
	KeyCancel     = &Key{code: types.KeyCodeCancel}
	KeyPrint      = &Key{code: types.KeyCodePrint}
	KeyPause      = &Key{code: types.KeyCodePause}
	KeyBacktab    = &Key{code: types.KeyCodeBacktab}
	KeyF1         = &Key{code: types.KeyCodeF1}
	KeyF2         = &Key{code: types.KeyCodeF2}
	KeyF3         = &Key{code: types.KeyCodeF3}
	KeyF4         = &Key{code: types.KeyCodeF4}
	KeyF5         = &Key{code: types.KeyCodeF5}
	KeyF6         = &Key{code: types.KeyCodeF6}
	KeyF7         = &Key{code: types.KeyCodeF7}
	KeyF8         = &Key{code: types.KeyCodeF8}
	KeyF9         = &Key{code: types.KeyCodeF9}
	KeyF10        = &Key{code: types.KeyCodeF10}
	KeyF11        = &Key{code: types.KeyCodeF11}
	KeyF12        = &Key{code: types.KeyCodeF12}
	KeyF13        = &Key{code: types.KeyCodeF13}
	KeyF14        = &Key{code: types.KeyCodeF14}
	KeyF15        = &Key{code: types.KeyCodeF15}
	KeyF16        = &Key{code: types.KeyCodeF16}
	KeyF17        = &Key{code: types.KeyCodeF17}
	KeyF18        = &Key{code: types.KeyCodeF18}
	KeyF19        = &Key{code: types.KeyCodeF19}
	KeyF20        = &Key{code: types.KeyCodeF20}
	KeyF21        = &Key{code: types.KeyCodeF21}
	KeyF22        = &Key{code: types.KeyCodeF22}
	KeyF23        = &Key{code: types.KeyCodeF23}
	KeyF24        = &Key{code: types.KeyCodeF24}
	KeyF25        = &Key{code: types.KeyCodeF25}
	KeyF26        = &Key{code: types.KeyCodeF26}
	KeyF27        = &Key{code: types.KeyCodeF27}
	KeyF28        = &Key{code: types.KeyCodeF28}
	KeyF29        = &Key{code: types.KeyCodeF29}
	KeyF30        = &Key{code: types.KeyCodeF30}
	KeyF31        = &Key{code: types.KeyCodeF31}
	KeyF32        = &Key{code: types.KeyCodeF32}
	KeyF33        = &Key{code: types.KeyCodeF33}
	KeyF34        = &Key{code: types.KeyCodeF34}
	KeyF35        = &Key{code: types.KeyCodeF35}
	KeyF36        = &Key{code: types.KeyCodeF36}
	KeyF37        = &Key{code: types.KeyCodeF37}
	KeyF38        = &Key{code: types.KeyCodeF38}
	KeyF39        = &Key{code: types.KeyCodeF39}
	KeyF40        = &Key{code: types.KeyCodeF40}
	KeyF41        = &Key{code: types.KeyCodeF41}
	KeyF42        = &Key{code: types.KeyCodeF42}
	KeyF43        = &Key{code: types.KeyCodeF43}
	KeyF44        = &Key{code: types.KeyCodeF44}
	KeyF45        = &Key{code: types.KeyCodeF45}
	KeyF46        = &Key{code: types.KeyCodeF46}
	KeyF47        = &Key{code: types.KeyCodeF47}
	KeyF48        = &Key{code: types.KeyCodeF48}
	KeyF49        = &Key{code: types.KeyCodeF49}
	KeyF50        = &Key{code: types.KeyCodeF50}
	KeyF51        = &Key{code: types.KeyCodeF51}
	KeyF52        = &Key{code: types.KeyCodeF52}
	KeyF53        = &Key{code: types.KeyCodeF53}
	KeyF54        = &Key{code: types.KeyCodeF54}
	KeyF55        = &Key{code: types.KeyCodeF55}
	KeyF56        = &Key{code: types.KeyCodeF56}
	KeyF57        = &Key{code: types.KeyCodeF57}
	KeyF58        = &Key{code: types.KeyCodeF58}
	KeyF59        = &Key{code: types.KeyCodeF59}
	KeyF60        = &Key{code: types.KeyCodeF60}
	KeyF61        = &Key{code: types.KeyCodeF61}
	KeyF62        = &Key{code: types.KeyCodeF62}
	KeyF63        = &Key{code: types.KeyCodeF63}
	KeyF64        = &Key{code: types.KeyCodeF64}
	KeyMenu       = &Key{code: types.KeyCodeMenu}
	KeyCapsLock   = &Key{code: types.KeyCodeCapsLock}
	KeyScrollLock = &Key{code: types.KeyCodeScrollLock}
	KeyNumLock    = &Key{code: types.KeyCodeNumLock}
)

var nonPrintableKeyCodeToString = map[types.KeyCode]string{
	types.KeyCodeUp:         "Up",
	types.KeyCodeDown:       "Down",
	types.KeyCodeRight:      "Right",
	types.KeyCodeLeft:       "Left",
	types.KeyCodeUpLeft:     "UpLeft",
	types.KeyCodeUpRight:    "UpRight",
	types.KeyCodeDownLeft:   "DownLeft",
	types.KeyCodeDownRight:  "DownRight",
	types.KeyCodeCenter:     "Center",
	types.KeyCodePgUp:       "PgUp",
	types.KeyCodePgDn:       "PgDn",
	types.KeyCodeHome:       "Home",
	types.KeyCodeEnd:        "End",
	types.KeyCodeInsert:     "Insert",
	types.KeyCodeDelete:     "Delete",
	types.KeyCodeHelp:       "Help",
	types.KeyCodeExit:       "Exit",
	types.KeyCodeClear:      "Clear",
	types.KeyCodeCancel:     "Cancel",
	types.KeyCodePrint:      "Print",
	types.KeyCodePause:      "Pause",
	types.KeyCodeBacktab:    "Backtab",
	types.KeyCodeF1:         "F1",
	types.KeyCodeF2:         "F2",
	types.KeyCodeF3:         "F3",
	types.KeyCodeF4:         "F4",
	types.KeyCodeF5:         "F5",
	types.KeyCodeF6:         "F6",
	types.KeyCodeF7:         "F7",
	types.KeyCodeF8:         "F8",
	types.KeyCodeF9:         "F9",
	types.KeyCodeF10:        "F10",
	types.KeyCodeF11:        "F11",
	types.KeyCodeF12:        "F12",
	types.KeyCodeF13:        "F13",
	types.KeyCodeF14:        "F14",
	types.KeyCodeF15:        "F15",
	types.KeyCodeF16:        "F16",
	types.KeyCodeF17:        "F17",
	types.KeyCodeF18:        "F18",
	types.KeyCodeF19:        "F19",
	types.KeyCodeF20:        "F20",
	types.KeyCodeF21:        "F21",
	types.KeyCodeF22:        "F22",
	types.KeyCodeF23:        "F23",
	types.KeyCodeF24:        "F24",
	types.KeyCodeF25:        "F25",
	types.KeyCodeF26:        "F26",
	types.KeyCodeF27:        "F27",
	types.KeyCodeF28:        "F28",
	types.KeyCodeF29:        "F29",
	types.KeyCodeF30:        "F30",
	types.KeyCodeF31:        "F31",
	types.KeyCodeF32:        "F32",
	types.KeyCodeF33:        "F33",
	types.KeyCodeF34:        "F34",
	types.KeyCodeF35:        "F35",
	types.KeyCodeF36:        "F36",
	types.KeyCodeF37:        "F37",
	types.KeyCodeF38:        "F38",
	types.KeyCodeF39:        "F39",
	types.KeyCodeF40:        "F40",
	types.KeyCodeF41:        "F41",
	types.KeyCodeF42:        "F42",
	types.KeyCodeF43:        "F43",
	types.KeyCodeF44:        "F44",
	types.KeyCodeF45:        "F45",
	types.KeyCodeF46:        "F46",
	types.KeyCodeF47:        "F47",
	types.KeyCodeF48:        "F48",
	types.KeyCodeF49:        "F49",
	types.KeyCodeF50:        "F50",
	types.KeyCodeF51:        "F51",
	types.KeyCodeF52:        "F52",
	types.KeyCodeF53:        "F53",
	types.KeyCodeF54:        "F54",
	types.KeyCodeF55:        "F55",
	types.KeyCodeF56:        "F56",
	types.KeyCodeF57:        "F57",
	types.KeyCodeF58:        "F58",
	types.KeyCodeF59:        "F59",
	types.KeyCodeF60:        "F60",
	types.KeyCodeF61:        "F61",
	types.KeyCodeF62:        "F62",
	types.KeyCodeF63:        "F63",
	types.KeyCodeF64:        "F64",
	types.KeyCodeMenu:       "Menu",
	types.KeyCodeCapsLock:   "CapsLock",
	types.KeyCodeScrollLock: "ScrollLock",
	types.KeyCodeNumLock:    "NumLock",
}

var tcellKeyToNonPrintableKeyCode = map[tcell.Key]types.KeyCode{
	tcell.KeyUp:         types.KeyCodeUp,
	tcell.KeyDown:       types.KeyCodeDown,
	tcell.KeyRight:      types.KeyCodeRight,
	tcell.KeyLeft:       types.KeyCodeLeft,
	tcell.KeyUpLeft:     types.KeyCodeUpLeft,
	tcell.KeyUpRight:    types.KeyCodeUpRight,
	tcell.KeyDownLeft:   types.KeyCodeDownLeft,
	tcell.KeyDownRight:  types.KeyCodeDownRight,
	tcell.KeyCenter:     types.KeyCodeCenter,
	tcell.KeyPgUp:       types.KeyCodePgUp,
	tcell.KeyPgDn:       types.KeyCodePgDn,
	tcell.KeyHome:       types.KeyCodeHome,
	tcell.KeyEnd:        types.KeyCodeEnd,
	tcell.KeyInsert:     types.KeyCodeInsert,
	tcell.KeyDelete:     types.KeyCodeDelete,
	tcell.KeyHelp:       types.KeyCodeHelp,
	tcell.KeyExit:       types.KeyCodeExit,
	tcell.KeyClear:      types.KeyCodeClear,
	tcell.KeyCancel:     types.KeyCodeCancel,
	tcell.KeyPrint:      types.KeyCodePrint,
	tcell.KeyPause:      types.KeyCodePause,
	tcell.KeyBacktab:    types.KeyCodeBacktab,
	tcell.KeyF1:         types.KeyCodeF1,
	tcell.KeyF2:         types.KeyCodeF2,
	tcell.KeyF3:         types.KeyCodeF3,
	tcell.KeyF4:         types.KeyCodeF4,
	tcell.KeyF5:         types.KeyCodeF5,
	tcell.KeyF6:         types.KeyCodeF6,
	tcell.KeyF7:         types.KeyCodeF7,
	tcell.KeyF8:         types.KeyCodeF8,
	tcell.KeyF9:         types.KeyCodeF9,
	tcell.KeyF10:        types.KeyCodeF10,
	tcell.KeyF11:        types.KeyCodeF11,
	tcell.KeyF12:        types.KeyCodeF12,
	tcell.KeyF13:        types.KeyCodeF13,
	tcell.KeyF14:        types.KeyCodeF14,
	tcell.KeyF15:        types.KeyCodeF15,
	tcell.KeyF16:        types.KeyCodeF16,
	tcell.KeyF17:        types.KeyCodeF17,
	tcell.KeyF18:        types.KeyCodeF18,
	tcell.KeyF19:        types.KeyCodeF19,
	tcell.KeyF20:        types.KeyCodeF20,
	tcell.KeyF21:        types.KeyCodeF21,
	tcell.KeyF22:        types.KeyCodeF22,
	tcell.KeyF23:        types.KeyCodeF23,
	tcell.KeyF24:        types.KeyCodeF24,
	tcell.KeyF25:        types.KeyCodeF25,
	tcell.KeyF26:        types.KeyCodeF26,
	tcell.KeyF27:        types.KeyCodeF27,
	tcell.KeyF28:        types.KeyCodeF28,
	tcell.KeyF29:        types.KeyCodeF29,
	tcell.KeyF30:        types.KeyCodeF30,
	tcell.KeyF31:        types.KeyCodeF31,
	tcell.KeyF32:        types.KeyCodeF32,
	tcell.KeyF33:        types.KeyCodeF33,
	tcell.KeyF34:        types.KeyCodeF34,
	tcell.KeyF35:        types.KeyCodeF35,
	tcell.KeyF36:        types.KeyCodeF36,
	tcell.KeyF37:        types.KeyCodeF37,
	tcell.KeyF38:        types.KeyCodeF38,
	tcell.KeyF39:        types.KeyCodeF39,
	tcell.KeyF40:        types.KeyCodeF40,
	tcell.KeyF41:        types.KeyCodeF41,
	tcell.KeyF42:        types.KeyCodeF42,
	tcell.KeyF43:        types.KeyCodeF43,
	tcell.KeyF44:        types.KeyCodeF44,
	tcell.KeyF45:        types.KeyCodeF45,
	tcell.KeyF46:        types.KeyCodeF46,
	tcell.KeyF47:        types.KeyCodeF47,
	tcell.KeyF48:        types.KeyCodeF48,
	tcell.KeyF49:        types.KeyCodeF49,
	tcell.KeyF50:        types.KeyCodeF50,
	tcell.KeyF51:        types.KeyCodeF51,
	tcell.KeyF52:        types.KeyCodeF52,
	tcell.KeyF53:        types.KeyCodeF53,
	tcell.KeyF54:        types.KeyCodeF54,
	tcell.KeyF55:        types.KeyCodeF55,
	tcell.KeyF56:        types.KeyCodeF56,
	tcell.KeyF57:        types.KeyCodeF57,
	tcell.KeyF58:        types.KeyCodeF58,
	tcell.KeyF59:        types.KeyCodeF59,
	tcell.KeyF60:        types.KeyCodeF60,
	tcell.KeyF61:        types.KeyCodeF61,
	tcell.KeyF62:        types.KeyCodeF62,
	tcell.KeyF63:        types.KeyCodeF63,
	tcell.KeyF64:        types.KeyCodeF64,
	tcell.KeyMenu:       types.KeyCodeMenu,
	tcell.KeyCapsLock:   types.KeyCodeCapsLock,
	tcell.KeyScrollLock: types.KeyCodeScrollLock,
	tcell.KeyNumLock:    types.KeyCodeNumLock,
}
