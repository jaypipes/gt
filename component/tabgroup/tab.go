package tabgroup

import (
	"context"

	"github.com/jaypipes/gt/core/key"
	"github.com/jaypipes/gt/core/keyshortcut"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/element/vdiv"
	"github.com/jaypipes/gt/types"
)

// newTab returns a new Tab instance with the supplied group and ID.
func newTab(
	ctx context.Context,
	group *TabGroup,
	id string,
) *Tab {
	d := vdiv.New(ctx, element.WithID(id))
	return &Tab{
		VDiv:  *d,
		group: group,
	}
}

// Tab is a group of Elements displayed when the Tab is active in a TabGroup's
// Bar.
type Tab struct {
	vdiv.VDiv
	// group is the TabGroup the Tab belongs to.
	group *TabGroup
	// title is the text that appears in the Tab's tab.
	title string
	// activeKey is the key combination that should trigger setting
	// this Tab as the active Tab in the TabGroup.
	activeKey types.Key
}

// Group returns a pointer to the TabGroup to which the Tab belongs.
func (t *Tab) Group() *TabGroup {
	return t.group
}

// SetTitle sets the Tab's title, which is the text that appears in the Tab's
// tab.
func (t *Tab) SetTitle(title string) *Tab {
	t.title = title
	return t
}

// Title returns the text that appears in the Tab's tab.
func (t *Tab) Title() string {
	return t.title
}

// SetContent sets the thing that will be rendered in the Tab.
func (t *Tab) SetContent(content types.Node) {
	t.RemoveAllChildren()
	t.AppendChild(content)
}

// WithContent sets the thing that will be rendered in the Tab and returns the
// Tab.
func (t *Tab) WithContent(content types.Node) *Tab {
	t.SetContent(content)
	return t
}

// AppendContent adds a child Element to the Tab's content and returns the
// Tab.
func (t *Tab) AppendContent(content types.Node) *Tab {
	t.AppendChild(content)
	return t
}

// SetActiveKeyPress sets the key combination that should trigger setting
// this Tab as the active Tab in the TabGroup.
//
// The keypress combination can be a string -- e.g. "Ctrl+C", "Esc" -- or a
// [tcell.Key] code -- e.g. tcell.KeyCtrlC, KeyEscape.
func (t *Tab) SetActiveKey(subject any) {
	k := key.New(subject)
	ctx := context.TODO()
	t.activeKey = k
	cb := func(_ context.Context) {
		t.group.SetActiveTab(t.ID())
	}
	ks := keyshortcut.New(ctx, keyshortcut.WithKey(k), keyshortcut.WithCallback(cb))
	t.group.SetKeyShortcut(ks)
}

// SetActiveKeyPress sets the key combination that should trigger setting
// this Tab as the active Tab in the TabGroup and returns the Tab.
func (t *Tab) WithActiveKey(subject any) *Tab {
	t.SetActiveKey(subject)
	return t
}

// ActiveKeyPress returns the key combination that triggers setting this
// Tab as the active Tab in the TabGroup
func (t *Tab) ActiveKey() types.Key {
	return t.activeKey
}
