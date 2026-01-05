package text

import "github.com/jaypipes/gt/core/types"

// String converts the supplied types.Text to a string
func String[T types.Text](t T) string {
	switch t := any(t).(type) {
	case string:
		return t
	case []byte:
		return string(t)
	}
	return ""
}
