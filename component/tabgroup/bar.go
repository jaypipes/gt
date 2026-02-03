package tabgroup

import (
	"context"
	"fmt"

	"github.com/jaypipes/gt/core"
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
	DefaultBarBorder = types.Border{
		Top:         types.Side{Content: ""},
		Bottom:      types.Side{Content: "─"},
		Left:        types.Side{Content: ""},
		Right:       types.Side{Content: ""},
		TopLeft:     types.Side{Content: ""},
		TopRight:    types.Side{Content: ""},
		BottomLeft:  types.Side{Content: ""},
		BottomRight: types.Side{Content: ""},
	}
	DefaultTitlePadding      = types.Pad(1)
	DefaultTitleActiveBorder = types.Border{
		Top:         types.Side{Content: "━"},
		Bottom:      types.Side{Content: ""},
		Left:        types.Side{Content: ""},
		Right:       types.Side{Content: ""},
		TopLeft:     types.Side{Content: ""},
		TopRight:    types.Side{Content: ""},
		BottomLeft:  types.Side{Content: ""},
		BottomRight: types.Side{Content: ""},
	}
	DefaultTitleInactiveBorder = types.Border{
		Top:         types.Side{Content: ""},
		Bottom:      types.Side{Content: ""},
		Left:        types.Side{Content: ""},
		Right:       types.Side{Content: ""},
		TopLeft:     types.Side{Content: ""},
		TopRight:    types.Side{Content: ""},
		BottomLeft:  types.Side{Content: ""},
		BottomRight: types.Side{Content: ""},
	}
)

func defaultBar(ctx context.Context, group *TabGroup) *Bar {
	barID := fmt.Sprintf("%s-bar", group.ID())
	d := div.New(ctx, element.WithID(barID))
	d.SetDisplay(types.DisplayBlock)
	d.SetHeight(core.Fixed(5))
	d.SetPadding(types.PadHorizontal(2))
	d.SetBorder(DefaultBarBorder)
	return &Bar{
		Div:                 *d,
		group:               group,
		location:            DefaultBarLocation,
		titlePadding:        DefaultTitlePadding,
		titleActiveBorder:   DefaultTitleActiveBorder,
		titleInactiveBorder: DefaultTitleInactiveBorder,
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
	// titlePadding is the padding around individual Tab titles in the Bar.
	titlePadding types.Padding
	// titleInactiveBorder is the border around inactive Tab titles in the Bar.
	titleInactiveBorder types.Border
	// titleActiveBorder is the border around the active Tab title in the Bar.
	titleActiveBorder types.Border
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

// SetTitlePadding sets the padding around individual titles of Tabs in the the
// Bar.
func (b *Bar) SetTitlePadding(p types.Padding) {
	b.titlePadding = p
}

// SetActiveTitleBorder sets the border around the active Tab in the bar.
func (b *Bar) SetActiveTitleBorder(border types.Border) {
	b.titleActiveBorder = border
}

// SetInactiveTitleBorder sets the border around the inactive Tabs in the bar.
func (b *Bar) SetInactiveTitleBorder(border types.Border) {
	b.titleInactiveBorder = border
}

func (b *Bar) Build(ctx context.Context) {
	// Clear any previously-built children from the TabGroup's container.
	b.RemoveAllChildren()
	for x, tab := range b.group.tabs {
		tabID := fmt.Sprintf("tab-group-%s-bar-tab-%s", b.group.ID(), tab.ID())
		tabEl := span.New(
			ctx,
			element.WithID(tabID),
			element.WithTextContent(tab.Title()),
			element.WithAlignment(types.AlignmentCenter),
			element.WithPadding(b.titlePadding),
			element.WithDisplay(types.DisplayInlineBlock),
			element.WithWidth(core.Fixed(12)),
		)
		if x == b.group.curTab {
			tabEl.SetBorder(b.titleActiveBorder)
		} else {
			tabEl.SetBorder(b.titleInactiveBorder)
		}
		b.AppendChild(tabEl)
	}
}
