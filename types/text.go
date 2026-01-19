package types

// Text is either a []byte or a string
type Text interface {
	string | []byte
}
