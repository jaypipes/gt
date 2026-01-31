package tabgroup

import (
	"github.com/jaypipes/gt/types"
)

// Tab is a group of Elements displayed when the Tab is active in a TabGroup.
type Tab struct {
	// group is the TabGroup the Tab belongs to.
	group *TabGroup
	// id is the unique identifier of the Tab.
	id string
	// title is the text that appears in the Tab's tab.
	title string
	// content is the Element that represents the root of the Tab's content.
	content types.Element
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

// ID returns the ID of the Tab
func (t *Tab) ID() string {
	return t.id
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

// SetContent sets the Element that is the root of the Tab's content.
func (t *Tab) SetContent(el types.Element) *Tab {
	t.content = el
	return t
}

// Content returns the Element that is the root of the Tab's content.
func (t *Tab) Content() types.Element {
	return t.content
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
