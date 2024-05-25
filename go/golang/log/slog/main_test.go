package main_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var keys = []string{
	"trace_id",
	"span_id",
}

type Handler struct {
	slog.Handler

	groups []string
}

func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	if value, ok := ctx.Value("trace_id").(string); ok {
		record.AddAttrs(slog.String("trace_id", value))
	}
	return h.Handler.Handle(ctx, record)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	for _, group := range h.groups {
		for i := 0; i < len(attrs); i++ {
			attrs[i] = slog.Group(group, attrs[i])
		}
		h.Handler = h.Handler.WithAttrs(attrs)
	}
	return h
}
func (h *Handler) WithGroup(name string) slog.Handler {
	h.groups = append(h.groups, name)
	return h
}

func TestA(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")
	value := ctx.Value("key")
	assert.NotNil(t, value)
	assert.Equal(t, "value", value)
}

func TestB(t *testing.T) {
	logger := slog.New(&Handler{
		Handler: slog.NewTextHandler(os.Stderr, nil),
		groups:  []string{},
	})
	// logger := slog.New(&Handler{
	// 	Handler: slog.NewTextHandler(os.Stderr, nil),
	// 	nonRoot: false,
	// })
	ctx := context.WithValue(context.Background(), "trace_id", "trace_value")
	// ctx := context.WithValue(context.Background(), "key", "value")
	logger.InfoContext(ctx, "msg0")
	logger = logger.With("k1", "v1")
	logger.InfoContext(ctx, "msg1")
	logger = logger.WithGroup("g2").With("k2", "v2")
	logger.InfoContext(ctx, "msg2")
	logger = logger.With("k3", "v3")
	logger.InfoContext(ctx, "msg3")
	logger = logger.WithGroup("g4").With("k4", "v4")
	logger.InfoContext(ctx, "msg4")
}
