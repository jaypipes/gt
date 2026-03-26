package tabgroup

import (
	"context"
	"strings"

	"github.com/samber/lo"

	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/element/vdiv"
	"github.com/jaypipes/gt/types"
)

// New returns a new instance of a TabGroup with the given ID.
func New(ctx context.Context, id string) *TabGroup {
	d := vdiv.New(ctx, element.WithID(id))
	g := &TabGroup{
		VDiv:    *d,
		tabs:    []*Tab{},
		rebuild: true,
	}
	g.bar = defaultBar(ctx, g)
	return g
}

// TabGroup is a Component that groups a set of Tab Components.
type TabGroup struct {
	vdiv.VDiv
	// rebuild will be true when the TabGroup's content needs to be rebuilt.
	rebuild bool
	// bar contains styling and layout cues for the bar of Tabs in the
	// TabGroup.
	bar *Bar
	// tabs is the collection of Tabs managed by the TabGroup.
	tabs []*Tab
	// activeTab is the ID of the active Tab.
	activeTab int

	// keyShortcuts stores the TabGroups's set of key shortcuts.
	keyShortcuts []types.KeyShortcut
}

// Bar returns the Bar object that can be styled separately.
func (g *TabGroup) Bar() *Bar {
	return g.bar
}

// Tab returns the Tab with the supplied ID. If no such Tab exists, a new
// empty Tab with that ID is returned.
func (g *TabGroup) Tab(ctx context.Context, id string) *Tab {
	t, ok := lo.Find(g.tabs, func(t *Tab) bool {
		return strings.EqualFold(t.ID(), id)
	})
	if !ok {
		t = newTab(ctx, g, id)
		g.tabs = append(g.tabs, t)
		g.activeTab = len(g.tabs) - 1
	}
	return t
}

// Tabs returns the collection of the TabGroup's Tabs.
func (g *TabGroup) Tabs() []*Tab {
	return g.tabs
}

// CurrentTab returns the currently active (displaying) Tab.
func (g *TabGroup) CurrentTab() *Tab {
	return g.tabs[g.activeTab]
}

// SetActiveTab sets the currently active (displaying) Tab.
func (g *TabGroup) SetActiveTab(id string) *TabGroup {
	_, idx, ok := lo.FindIndexOf(g.tabs, func(t *Tab) bool {
		return strings.EqualFold(t.ID(), id)
	})
	if ok {
		if g.activeTab != idx {
			g.activeTab = idx
			g.rebuild = true
		}
	}
	return g
}

// KeyPress checks for any KeyShortcuts that are registered with the TabGroup
// and executes any matched callback. If no KeyShortcuts are matched, we
// execute the View's internal vdiv Element's KeyPress method.
func (g *TabGroup) KeyPress(ctx context.Context, ev types.KeyPressEvent) bool {
	k := ev.Key()
	for _, ks := range g.keyShortcuts {
		ksk := ks.Key()
		if ksk.Equal(k) {
			cb := ks.Callback()
			cb(ctx)
			return true
		}
	}
	return g.VDiv.KeyPress(ctx, ev)
}

// SetKeyShortcut registers a TabGroup-level KeyShortcut that will execute upon
// a key press combination.
func (g *TabGroup) SetKeyShortcut(shortcut types.KeyShortcut) {
	k := shortcut.Key()
	for _, ks := range g.keyShortcuts {
		ksk := ks.Key()
		if ksk.Equal(k) {
			gtlog.Warn(
				context.TODO(),
				"key shortcut %q shadows previously-registered "+
					"tabgroup-level key shortcut",
				k,
			)
		}
	}
	g.keyShortcuts = append(g.keyShortcuts, shortcut)
}

// Build constructs the tab bar and tab content elements.
func (g *TabGroup) Build(
	ctx context.Context,
) {
	if !g.rebuild {
		return
	}
	gtlog.Debug(ctx, "TabGroup.Build[%s]", g.ID())

	// Clear any previously-built children from the TabGroup's container.
	g.RemoveAllChildren()

	g.bar.Build(ctx)
	g.AppendChild(g.bar)

	activeTab := g.CurrentTab()
	if activeTab != nil {
		g.AppendChild(&activeTab.VDiv)
	}
	g.rebuild = false
}
