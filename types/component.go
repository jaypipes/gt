package types

// Component represents a collection of Elements that, together, act as a
// single unit.
type Component interface {
	Buildable
	Plottable
}
