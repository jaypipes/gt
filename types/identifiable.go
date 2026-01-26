package types

// Identifiable represents something that can be identified.
type Identifiable interface {
	// SetID sets the Identifiable's unique identifier.
	SetID(string)
	// ID returns a string identifier for the Identifiable.
	ID() string
	// String returns a string representation of the Identifiable.
	String() string
}
