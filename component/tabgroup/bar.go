package tabgroup

import (
	"context"
	"fmt"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/border"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/element/div"
	"github.com/jaypipes/gt/element/span"
	"github.com/jaypipes/gt/types"
)

const (
	DefaultBarLocation = BarLocationTop
)

var (
	// The default bar border is just a single line on the bottom of the bar.
	DefaultBarBorder            = border.New(border.WithB("-"))
	DefaultBarTabPadding        = types.Pad(1)
	DefaultBarTabActiveBorder   = border.New(border.WithT("━"))
	DefaultBarTabInactiveBorder = border.None()
)

func defaultBar(ctx context.Context, group *TabGroup) *Bar {
	barID := fmt.Sprintf("%s-bar", group.ID())
	d := div.New(ctx, element.WithID(barID))
	d.SetDisplay(types.DisplayBlock)
	d.SetHeight(core.Fixed(5))
	d.SetPadding(types.PadHorizontal(2))
	d.SetBorder(DefaultBarBorder)
	return &Bar{
		Div:               *d,
		group:             group,
		location:          DefaultBarLocation,
		tabPadding:        DefaultBarTabPadding,
		tabActiveBorder:   DefaultBarTabActiveBorder,
		tabInactiveBorder: DefaultBarTabInactiveBorder,
	}
}

// BarLocation indicates where the bar of Tabs appears in a TabGroup.
type BarLocation uint8

const (
	BarLocationTop BarLocation = iota
	BarLocationBottom
	BarLocationLeft
	BarLocationRight
)

// Bar represents the bar of Tabs in a TabGroup.
type Bar struct {
	div.Div
	// group is the TabGroup this Bar will display Tab titles for.
	group *TabGroup
	// location is where the Bar will appear.
	location BarLocation
	// tabPadding is the padding around individual Tabs in the Bar.
	tabPadding types.Padding
	// tabInactiveBorder is the border around inactive Tabs in the Bar.
	tabInactiveBorder types.Border
	// tabActiveBorder is the border around the active Tabs in the Bar.
	tabActiveBorder types.Border
	// tabHoverBorder is the border around inactive Tabs in the Bar when the
	// mouse hovers over that Tab.
	tabHoverBorder types.Border
}

// SetLocation sets where the Bar will appear.
func (b *Bar) SetLocation(loc BarLocation) {
	b.location = loc
}

// SetPadding sets the padding around the Bar itself. To set the padding around
// individual titles of Tabs in the Bar, use SetTitlePadding.
func (b *Bar) SetPadding(p types.Padding) {
	b.Div.SetPadding(p)
}

// SetBorder sets the border around the Bar itself. To set the border around
// the titles of the Tabs in the Bar, use SetInactiveTitleBorder and
// SetActiveTitleBorder.
func (b *Bar) SetBorder(border types.Border) {
	b.Div.SetBorder(border)
}

// SetBarTabPadding sets the padding around individual titles of Tabs in the the
// Bar.
func (b *Bar) SetBarTabPadding(p types.Padding) {
	b.tabPadding = p
}

// SetActiveBarTabBorder sets the border around the active Tab in the bar.
func (b *Bar) SetActiveBarTabBorder(border types.Border) {
	b.tabActiveBorder = border
}

// SetInactiveBarTabBorder sets the border around the inactive Tabs in the bar.
func (b *Bar) SetInactiveBarTabBorder(border types.Border) {
	b.tabInactiveBorder = border
}

func (b *Bar) Build(ctx context.Context) {
	// Clear any previously-built children from the Bar's container.
	b.RemoveAllChildren()
	for x, tab := range b.group.tabs {
		tabID := fmt.Sprintf("tab-group-%s-bar-tab-%s", b.group.ID(), tab.ID())
		tabEl := span.New(
			ctx,
			element.WithID(tabID),
			element.WithTextContent(tab.Title()),
			element.WithAlignment(types.AlignmentCenter),
			element.WithPadding(b.tabPadding),
			element.WithDisplay(types.DisplayInlineBlock),
			element.WithWidth(core.Fixed(12)),
		)
		if x == b.group.activeTab {
			tabEl.SetBorder(b.tabActiveBorder)
		} else {
			onClick := func(ctx context.Context, ev types.MouseClickEvent) {
				if ev.Button() == types.MouseButtonPrimary {
					b.group.SetActiveTab(tab.ID())
				}
			}
			tabEl.OnMouseClick(onClick)
			tabEl.SetBorder(b.tabInactiveBorder)
		}
		b.AppendChild(tabEl)
	}
}
