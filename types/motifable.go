package types

// Motifable describes something that can have a Motif.
type Motifable interface {
	// Motif returns the Motif for the Motifable
	Motif() Motif
	// SetMotif sets the Motif for the Motifable
	SetMotif(Motif)
}
