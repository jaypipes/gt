package tabgroup

import (
	"context"
	"fmt"
	"strings"

	"github.com/samber/lo"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/box"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/element/div"
)

// New returns a new instance of a TabGroup.
func New(ctx context.Context, id string) *TabGroup {
	b := box.New(ctx)
	b.SetHeight(core.Percent(100))
	b.SetID(id)
	g := &TabGroup{
		Box:  b,
		tabs: []*Tab{},
	}
	g.bar = defaultBar(ctx, g)
	return g
}

// TabGroup is a Component that groups a set of Tab Components.
type TabGroup struct {
	box.Box
	// bar contains styling and layout cues for the bar of Tabs in the
	// TabGroup.
	bar *Bar
	// tabs is the collection of Tabs managed by the TabGroup.
	tabs []*Tab
	// curTab is the ID of the active Tab.
	curTab int
}

// Tab returns the Tab with the supplied ID. If no such Tab exists, a new
// empty Tab with that ID is returned.
func (g *TabGroup) Tab(ctx context.Context, id string) *Tab {
	t, ok := lo.Find(g.tabs, func(t *Tab) bool {
		return strings.EqualFold(t.ID(), id)
	})
	if !ok {
		t = &Tab{group: g, id: id}
		g.tabs = append(g.tabs, t)
		g.curTab = len(g.tabs) - 1
	}
	return t
}

// Tabs returns the collection of the TabGroup's Tabs.
func (g *TabGroup) Tabs() []*Tab {
	return g.tabs
}

// CurrentTab returns the currently active (displaying) Tab.
func (g *TabGroup) CurrentTab() *Tab {
	return g.tabs[g.curTab]
}

// SetCurrentTab sets the currently active (displaying) Tab.
func (g *TabGroup) SetCurrentTab(id string) *TabGroup {
	_, idx, ok := lo.FindIndexOf(g.tabs, func(t *Tab) bool {
		return strings.EqualFold(t.ID(), id)
	})
	if ok {
		g.curTab = idx
	}
	return g
}

// Build constructs the tab bar and tab content elements.
func (g *TabGroup) Build(
	ctx context.Context,
) {
	gtlog.Debug(ctx, "TabGroup.Build[%s]", g.ID())

	g.bar.Build(ctx)
	g.AppendChild(g.bar)

	curTab := g.CurrentTab()
	if curTab != nil {
		tabContentContainer := div.New(ctx, "")
		containerID := fmt.Sprintf("tab-content-container-%s", curTab.ID())
		tabContentContainer.SetID(containerID)
		tabContentContainer.SetHeight(core.Percent(100))
		tabContentContainer.SetWidth(core.Percent(100))
		tabContent := curTab.Content()
		if tabContent != nil {
			tabContentContainer.AppendChild(tabContent)
		}
		g.AppendChild(tabContentContainer)
	}
}
