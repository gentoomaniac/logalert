package safelog

import (
	"context"
	"log/slog"
	"sync"
)

const (
	replacement = "*****"
)

type SafeLogHandler struct {
	mu      *sync.Mutex
	handler slog.Handler
}

func NewSafeLogHandler(handler slog.Handler) *SafeLogHandler {
	h := &SafeLogHandler{mu: &sync.Mutex{}}
	h.handler = handler

	return h
}

func (h *SafeLogHandler) Enabled(ctx context.Context, level slog.Level) bool {
	// This is passed on to the nested handler
	return true
}

func (h *SafeLogHandler) Handle(ctx context.Context, r slog.Record) error {
	filteredRecord := slog.Record{Time: r.Time, Message: r.Message, PC: r.PC}
	var attrs []slog.Attr
	r.Attrs(func(attr slog.Attr) bool {
		attrs = append(attrs, attr)
		return true
	})
	filteredRecord.AddAttrs(filterAttrs(attrs)...)
	return h.handler.Handle(ctx, filteredRecord)
}

func (h *SafeLogHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h.handler = h.handler.WithAttrs(filterAttrs(attrs))
	return h
}

func (h *SafeLogHandler) WithGroup(name string) slog.Handler {
	h.handler = h.handler.WithGroup(name)
	return h
}

func filterAttrs(attrs []slog.Attr) []slog.Attr {
	fillteredAttrs := []slog.Attr{}

	for _, attr := range attrs {
		matched := false

		for _, re := range keyRegexp {
			if re.MatchString(attr.Key) {
				matched = true
				break
			}
		}

		if !matched {
			for _, re := range valueRegexp {
				if re.MatchString(attr.Value.String()) {
					matched = true
					break
				}
			}
		}

		if matched {
			attr.Value = slog.StringValue(replacement)
		}
		fillteredAttrs = append(fillteredAttrs, attr)
	}

	return fillteredAttrs
}
