package otelhandler_test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/vikstrous/slogexamples/otelhandler"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	trace "go.opentelemetry.io/otel/trace"
	"golang.org/x/exp/slog"
)

var (
	testTraceID = trace.TraceID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	testSpanID  = trace.SpanID{1, 2, 3, 4, 5, 6, 7, 8}
)

type TestIDGenerator struct{}

func (TestIDGenerator) NewIDs(ctx context.Context) (trace.TraceID, trace.SpanID) {
	return testTraceID, testSpanID
}

func (TestIDGenerator) NewSpanID(ctx context.Context, traceID trace.TraceID) trace.SpanID {
	return testSpanID
}

func TestOtelHandler(t *testing.T) {
	tracer := sdktrace.NewTracerProvider(sdktrace.WithIDGenerator(TestIDGenerator{})).Tracer("")

	// Pre-allocate the output buffer so the test doesn't cause extra allocations
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	sl := slog.New(otelhandler.NewOtelHandler(slog.NewTextHandler(buf, nil)))
	// Make sure the WithX methods don't cause the handler to get unwrapped
	sl = sl.WithGroup("g").With("key", "value")

	ctx := context.Background()
	ctx, span := tracer.Start(ctx, "example")
	allocsPerRun := testing.AllocsPerRun(1, func() {
		// 2 allocations come from creating two slog.Attrs and two come from converting the trace and span IDs from byte arrays to hex strings
		sl.InfoContext(ctx, "hello")
	})
	if allocsPerRun > 4 {
		t.Fatalf("too many allocations: %.0f", allocsPerRun)
	}
	span.End()

	output := buf.String()
	if !strings.Contains(output, "traceID=0102030405060708090a0b0c0d0e0f10") {
		t.Fatalf("log doesn't contain traceID: %s", output)
	}
	if !strings.Contains(output, "spanID=0102030405060708") {
		t.Fatalf("log doesn't contain spanID: %s", output)
	}
	t.Log(output)
}

func Example() {
	ctx := context.Background()
	sl := slog.New(otelhandler.NewOtelHandler(slog.NewTextHandler(os.Stderr, nil)))
	sl.InfoContext(ctx, "example")
}
