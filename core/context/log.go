package context

import (
	"context"
	"log/slog"
	"os"
	"strings"

	gtlog "github.com/jaypipes/gt/core/log"
)

const (
	envKeyDebug    = "GT_DEBUG"
	envKeyLogLevel = "GT_LOG_LEVEL"
)

var (
	logLevelKey     = ContextKey("gt.log.level")
	defaultLogLevel = slog.LevelWarn
)

// WithLogLevel allows overriding the default log level of WARN.
func WithLogLevel(level slog.Level) ContextModifier {
	return func(ctx context.Context) context.Context {
		gtlog.SetLevel(level)
		return context.WithValue(ctx, logLevelKey, level)
	}
}

// LogLevel gets a context's logger's log level or the default if none is set.
func LogLevel(ctx context.Context) slog.Level {
	if ctx == nil {
		return defaultLogLevel
	}
	if v := ctx.Value(logLevelKey); v != nil {
		return v.(slog.Level)
	}
	return defaultLogLevel
}

// EnvOrDefaultLogLevel return true if ghw should not output warnings
// based on the GHW_LOG_LEVEL environs variable.
func EnvOrDefaultLogLevel() slog.Level {
	if _, exists := os.LookupEnv(envKeyDebug); exists {
		return slog.LevelDebug
	}
	if ll, exists := os.LookupEnv(envKeyLogLevel); exists {
		switch strings.ToLower(ll) {
		case "debug":
			return slog.LevelDebug
		case "info":
			return slog.LevelInfo
		case "warn", "warning":
			return slog.LevelWarn
		case "err", "error":
			return slog.LevelError
		default:
			return defaultLogLevel
		}
	}
	return defaultLogLevel
}

// WithDebug enables verbose debugging output.
func WithDebug() ContextModifier {
	return WithLogLevel(slog.LevelDebug)
}
