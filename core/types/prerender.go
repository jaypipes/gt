package types

import "context"

// Prerenders
type Prerender interface {
	Prerender(context.Context, Screen, Rectangle)
}
