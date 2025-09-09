package grid

import (
	"github.com/jaypipes/gt/core/box"
)

type Option func(*Grid)

// New returns a new Grid instance.
func New(opts ...Option) *Grid {
	g := &Grid{}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

// Grid represents an abstraction of multiple Boxes arranged in a set of rows
// and columns.
type Grid struct {
	rows    []box.Box
	columns []box.Box
}
