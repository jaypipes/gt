package types

import "fmt"

// Identifiable represents something that can be identified.
type Identifiable interface {
	fmt.Stringer
	// SetID sets the Identifiable's unique identifier.
	SetID(string)
	// ID returns a string identifier for the Identifiable.
	ID() string
}
