package types

// Document describes a renderable tree of Nodes. It's similar to the Document
// Object Model's (DOM) interface of the same name.
type Document interface {
	// ElementByID returns the Element with the specified ID or nil if the
	// Document contains no Element with that identifier.
	ElementByID(string) Element
	// ElementsByClass returns all Elements having the specified type/class.
	ElementsByClass(string) []Element
}
