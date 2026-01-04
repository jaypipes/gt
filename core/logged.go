package core

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
	"time"
)

// Logged is a based class for anything that has a `slog.Logger` and can write
// log records.
type Logged struct {
	// log is the `log/slog.Logger` for the Logged.
	log *slog.Logger
}

// SetLogger sets the Logged's Logger.
func (l *Logged) SetLogger(log *slog.Logger) {
	l.log = log
}

// Logger returns the Logged's Logger.
func (l *Logged) Logger() *slog.Logger {
	return l.log
}

// Info outputs an INFO-level log message to the logger configured in the
// supplied context.
func (l *Logged) Info(ctx context.Context, format string, args ...any) {
	logger := l.log
	if logger == nil || !logger.Enabled(ctx, slog.LevelInfo) {
		return
	}
	var stack [1]uintptr
	runtime.Callers(2, stack[:]) // skip [Callers, Info]
	r := slog.NewRecord(
		time.Now(),
		slog.LevelInfo,
		strings.TrimSpace(
			fmt.Sprintf(format, args...),
		),
		stack[0],
	)
	_ = logger.Handler().Handle(ctx, r)
}

// Warn outputs an WARN-level log message to the logger configured in the
// supplied context.
func (l *Logged) Warn(ctx context.Context, format string, args ...any) {
	logger := l.log
	if logger == nil || !logger.Enabled(ctx, slog.LevelWarn) {
		return
	}
	var stack [1]uintptr
	runtime.Callers(2, stack[:]) // skip [Callers, Warn]
	r := slog.NewRecord(
		time.Now(),
		slog.LevelWarn,
		strings.TrimSpace(
			fmt.Sprintf(format, args...),
		),
		stack[0],
	)
	_ = logger.Handler().Handle(ctx, r)
}

// Debug outputs an DEBUG-level log message to the logger configured in the
// supplied context.
func (l *Logged) Debug(ctx context.Context, format string, args ...any) {
	logger := l.log
	if logger == nil || !logger.Enabled(ctx, slog.LevelDebug) {
		return
	}
	var stack [1]uintptr
	runtime.Callers(2, stack[:]) // skip [Callers, Debug]
	r := slog.NewRecord(
		time.Now(),
		slog.LevelDebug,
		strings.TrimSpace(
			fmt.Sprintf(format, args...),
		),
		stack[0],
	)
	_ = logger.Handler().Handle(ctx, r)
}
