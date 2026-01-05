package simple

import (
	"context"
	"io"
	"log"
	"log/slog"
)

// New returns a new Handler that outputs simple log records formatted as
// `{prefix}{level}: {message}` to the supplied io.Writer.
func New(prefix string, w io.Writer, opts *slog.HandlerOptions) *Handler {
	return &Handler{
		Handler: slog.NewTextHandler(
			w, opts,
		),
		prefix: prefix,
		l:      log.New(w, "", 0),
	}
}

// Handler is a custom log handler that outputs simple log records formatted as `{prefix}{level}: {message}`.
type Handler struct {
	slog.Handler
	prefix string
	l      *log.Logger
}

func (h *Handler) Handle(
	ctx context.Context,
	r slog.Record,
) error {
	level := r.Level.String() + ":"

	h.l.Printf("%s%-6s %s", h.prefix, level, r.Message)

	return nil
}
