package tabgroup

import (
	"context"

	"github.com/jaypipes/gt/element/vdiv"
	"github.com/jaypipes/gt/types"
)

// newTab returns a new Tab instance.
func newTab(
	ctx context.Context,
	group *TabGroup,
	id string,
) *Tab {
	d := vdiv.New(ctx, "")
	d.SetID(id)
	return &Tab{
		VDiv:  *d,
		group: group,
	}
}

// Tab is a group of Elements displayed when the Tab is active in a TabGroup.
type Tab struct {
	vdiv.VDiv
	// group is the TabGroup the Tab belongs to.
	group *TabGroup
	// title is the text that appears in the Tab's tab.
	title string
	// currentTabKeyPress is the key combination that should trigger setting
	// this Tab as the current Tab in the TabGroup.
	currentTabKeyPress string
	// keyPressMap contains key press combination callbacks registered for the
	// Tab.
	keyPressMap types.KeyPressMap
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

// SetCurrentTabKeyPress sets the key combination that should trigger setting
// this Tab as the current Tab in the TabGroup.
func (t *Tab) SetCurrentTabKeyPress(key string) {
	t.currentTabKeyPress = key
}

// SetCurrentTabKeyPress sets the key combination that should trigger setting
// this Tab as the current Tab in the TabGroup and returns the Tab.
func (t *Tab) WithCurrentTabKeyPress(key string) *Tab {
	t.SetCurrentTabKeyPress(key)
	return t
}

// CurrentTabKeyPress returns the key combination that triggers setting this
// Tab as the current Tab in the TabGroup
func (t *Tab) CurrentTabKeyPress() string {
	return t.currentTabKeyPress
}

// KeyPressMap returns a map, keyed by key press string combination, of
// callbacks to execute upon that key press.
func (t *Tab) KeyPressMap() types.KeyPressMap {
	return t.keyPressMap
}
