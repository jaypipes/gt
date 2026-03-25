package application

import (
	"context"

	"github.com/samber/lo"

	"github.com/jaypipes/gt/core/view"
	"github.com/jaypipes/gt/types"
)

// View returns the View with the supplied ID. If no such View exists, a new
// empty View with that ID is returned.
func (a *Application) View(ctx context.Context, id string) types.View {
	v, ok := a.views[id]
	if !ok {
		v = view.New(ctx, view.WithID(id))
		a.views[id] = v
		a.activeView = id
	}
	return v
}

// Views returns the collection of the Application's Views.
func (a *Application) Views() []types.View {
	return lo.Values(a.views)
}

// ActiveView returns the currently active (displaying) View.
func (a *Application) ActiveView() types.View {
	return a.views[a.activeView]
}

// SetActiveView sets the currently active (displaying) View.
func (a *Application) SetActiveView(id string) {
	a.activeView = id
}
