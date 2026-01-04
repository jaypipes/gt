package types

import (
	"context"
	"log/slog"
)

// Logged describes something that has a `slog.Logger` and can log records to
// it.
type Logged interface {
	// SetLogger sets the Logged's Logger.
	SetLogger(*slog.Logger)
	// Logger returns the Logged's Logger.
	Logger() *slog.Logger
	// Info outputs an INFO-level log message to the logger configured in the
	// supplied context.
	Info(context.Context, string, ...any)
	// Warn outputs an WARN level log message to the logger configured in the
	// supplied context.
	Warn(context.Context, string, ...any)
	// Debug outputs an DEBUG-level log message to the logger configured in the
	// supplied context.
	Debug(context.Context, string, ...any)
}
