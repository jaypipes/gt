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
