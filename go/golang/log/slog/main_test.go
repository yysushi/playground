package main_test

import (
	"bytes"
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeTime(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	return a
}

var keys = []string{
	"trace_id",
	"span_id",
}

type Handler struct {
	slog.Handler

	group string
}

func (h *Handler) Handle(ctx context.Context, record slog.Record) error {
	if value, ok := ctx.Value("trace_id").(string); ok && h.group == "" {
		record.AddAttrs(slog.String("trace_id", value))
	}
	return h.Handler.Handle(ctx, record)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if h.group == "" {
		return &Handler{Handler: h.Handler.WithAttrs(attrs)}
	}
	var groupedAttrs []slog.Attr
	for _, attr := range attrs {
		groupedAttrs = append(groupedAttrs, slog.Group(h.group, attr))
	}
	return &Handler{
		Handler: h.Handler.WithAttrs(groupedAttrs),
		group:   h.group,
	}
}
func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{
		Handler: h,
		group:   name,
	}
}

func TestA(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")
	value := ctx.Value("key")
	assert.NotNil(t, value)
	assert.Equal(t, "value", value)
}

func TestB(t *testing.T) {
	var buf bytes.Buffer

	logger := slog.New(&Handler{
		Handler: slog.NewTextHandler(&buf, &slog.HandlerOptions{ReplaceAttr: removeTime}),
		group:   "",
	})

	ctx := context.WithValue(context.Background(), "trace_id", "trace_value")
	logger.InfoContext(ctx, "msg0")
	assert.Equal(t, "level=INFO msg=msg0 trace_id=trace_value\n", buf.String())
	buf.Reset()

	logger = logger.With("k1", "v1")
	logger.InfoContext(ctx, "msg1")
	assert.Equal(t, "level=INFO msg=msg1 k1=v1 trace_id=trace_value\n", buf.String())
	buf.Reset()

	logger = logger.WithGroup("g2").With("k2", "v2")
	logger.InfoContext(ctx, "msg2")
	assert.Equal(t, "level=INFO msg=msg2 k1=v1 g2.k2=v2 trace_id=trace_value\n", buf.String())
	buf.Reset()

	logger = logger.With("k3", "v3")
	logger.InfoContext(ctx, "msg3")
	assert.Equal(t, "level=INFO msg=msg3 k1=v1 g2.k2=v2 g2.k3=v3 trace_id=trace_value\n", buf.String())
	buf.Reset()

	logger = logger.WithGroup("g4").With("k4", "v4")
	logger.InfoContext(ctx, "msg4")
	assert.Equal(t, "level=INFO msg=msg4 k1=v1 g2.k2=v2 g2.k3=v3 g2.g4.k4=v4 trace_id=trace_value\n", buf.String())
}
