package types

import "context"

// EventCallback describes a function that will execute when some Event fires.
type EventCallback func(context.Context)
