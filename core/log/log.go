package log

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"strings"
	"time"

	"github.com/jaypipes/gt/core/log/handler/simple"
)

const (
	DefaultLogLevel        = slog.LevelWarn
	defaultLogBufferSize   = 10 * 1024 // 10KiB
	defaultLogRecordPrefix = "(gt) "
)

var (
	// logBuf stores log records before a `gt.Application` is initialized,
	// allowing use of the `core/log.Info|Warn|Debug` functions before the
	// `gt.Application` has had a chance to route log records to its own
	// buffer. We do this to prevent outputting to STDERR and polluting
	// terminal output.
	//
	// When core/application.New()` is called, stored log records in logBuf are
	// copied into the `gt.Application`'s log buffer.
	logBuf *bytes.Buffer
)

var (
	logLevelVar   = new(slog.LevelVar)
	defaultLogger *slog.Logger
	// Logger is the package-level `slog.Logger` that the gt library uses.
	Logger *slog.Logger
)

func init() {
	logBuf = &bytes.Buffer{}
	logBuf.Grow(defaultLogBufferSize)
	logLevelVar.Set(DefaultLogLevel)
	/*
		defaultLogger = slog.New(slog.DiscardHandler)
	*/
	defaultLogger = slog.New(
		simple.New(
			defaultLogRecordPrefix,
			logBuf,
			&slog.HandlerOptions{
				Level: logLevelVar,
			},
		),
	)
	Logger = defaultLogger
}

// Level returns the active logging level for the gt package-level logger.
func Level() slog.Level {
	return logLevelVar.Level()
}

// SetLevel changes the log level for the package-level `slog.Logger` that the
// gt library uses.
func SetLevel(level slog.Level) {
	logLevelVar.Set(level)
}

// Records returns the buffered log records as a string.
func Records() string {
	return logBuf.String()
}

// Info outputs an INFO-level log message to the logger configured in the
// supplied context.
func Info(ctx context.Context, format string, args ...any) {
	if !Logger.Enabled(ctx, slog.LevelInfo) {
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
	_ = Logger.Handler().Handle(ctx, r)
}

// Warn outputs an WARN-level log message to the logger configured in the
// supplied context.
func Warn(ctx context.Context, format string, args ...any) {
	if !Logger.Enabled(ctx, slog.LevelWarn) {
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
	_ = Logger.Handler().Handle(ctx, r)
}

// Debug outputs an DEBUG-level log message to the logger configured in the
// supplied context.
func Debug(ctx context.Context, format string, args ...any) {
	if !Logger.Enabled(ctx, slog.LevelDebug) {
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
	_ = Logger.Handler().Handle(ctx, r)
}
