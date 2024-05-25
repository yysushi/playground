package main_test

import (
	"bytes"
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	logsv1 "go.opentelemetry.io/proto/otlp/logs/v1"
)

var traceKey = logsv1.File_opentelemetry_proto_logs_v1_logs_proto.Messages().ByName("LogRecord").Fields().ByNumber(9).Name()
var spanKey = logsv1.File_opentelemetry_proto_logs_v1_logs_proto.Messages().ByName("LogRecord").Fields().ByNumber(10).Name()

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

type OTLPHandler struct {
	slog.Handler

	group string
}

func (h *OTLPHandler) Handle(ctx context.Context, record slog.Record) error {
	if value, ok := ctx.Value("trace_id").(string); ok && h.group == "" {
		record.AddAttrs(slog.String("trace_id", value))
	}
	return h.Handler.Handle(ctx, record)
}

func (h *OTLPHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if h.group == "" {
		return &OTLPHandler{Handler: h.Handler.WithAttrs(attrs)}
	}
	var groupedAttrs []slog.Attr
	for _, attr := range attrs {
		groupedAttrs = append(groupedAttrs, slog.Group(h.group, attr))
	}
	return &OTLPHandler{
		Handler: h.Handler.WithAttrs(groupedAttrs),
		group:   h.group,
	}
}
func (h *OTLPHandler) WithGroup(name string) slog.Handler {
	return &OTLPHandler{
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

	logger := slog.New(&OTLPHandler{
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

func TestC(t *testing.T) {
	var spanA_ID [8]byte
	var spanB_ID [8]byte
	var spanA_TraceID [16]byte
	var spanB_TraceID [16]byte
	func() {
		// writer
		f, err := os.Create("traces.txt")
		require.NoError(t, err)
		defer f.Close()
		// exporter
		exporter, err := stdouttrace.New(
			stdouttrace.WithWriter(f),
		)
		require.NoError(t, err)
		// trace provider
		tp := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exporter))
		otel.SetTracerProvider(tp)

		tracer1 := otel.Tracer("tracer")
		ctxA, spanA := tracer1.Start(context.Background(), "spanA", trace.WithNewRoot())
		defer spanA.End()
		ctxB, spanB := tracer1.Start(ctxA, "spanA")
		defer spanB.End()
		spanA_ID = trace.SpanContextFromContext(ctxA).SpanID()
		spanB_ID = trace.SpanContextFromContext(ctxB).SpanID()
		spanA_TraceID = trace.SpanContextFromContext(ctxA).TraceID()
		spanB_TraceID = trace.SpanContextFromContext(ctxB).TraceID()
	}()
	assert.Equal(t, spanA_TraceID, spanB_TraceID)
	assert.NotEqual(t, spanA_ID, spanB_ID)
	// jq .Parent < traces.txt
}
