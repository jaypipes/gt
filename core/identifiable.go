package core

import (
	"fmt"

	"github.com/jaypipes/gt/types"
)

// Identifiable implements [types.Identifiable].
type Identifiable struct {
	// id is the string identifier of the Identifible
	id string
}

// SetID sets the Identifiable's identifier.
func (i *Identifiable) SetID(id string) {
	i.id = id
}

// ID returns the Identifiable's identifier.
func (i *Identifiable) ID() string {
	return i.id
}

// String returns a string representation of the Identifiable.
func (i *Identifiable) String() string {
	return fmt.Sprintf("id=%s", i.id)
}

// ID returns the string identifier of the supplied thing. If the thing is an
// Identifiable, it returns the result of the Identifiable's ID() method. If
// it's a Node, it returns the results of NodeID().
func ID(subject any) string {
	switch subject := subject.(type) {
	case types.Identifiable:
		return subject.ID()
	case types.Node:
		return subject.NodeID()
	default:
		return fmt.Sprintf("%T:%v", subject, subject)
	}
}

var _ types.Identifiable = (*Identifiable)(nil)
