// Package otelhandler shows a slog.Handler wrapper that adds extra attributes to log statements.
package otelhandler

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/otel/trace"
)

// OtelHandler adds extra fields to the record being logged containing traceID and spanID so that logs can be correlated with traces
type OtelHandler struct {
	slog.Handler
}

func NewOtelHandler(h slog.Handler) OtelHandler {
	return OtelHandler{h}
}

var _ slog.Handler = OtelHandler{}

func (o OtelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return OtelHandler{o.Handler.WithAttrs(attrs)}
}

func (o OtelHandler) WithGroup(group string) slog.Handler {
	return OtelHandler{o.Handler.WithGroup(group)}
}

func (o OtelHandler) Handle(ctx context.Context, r slog.Record) error {
	if spanContext := trace.SpanContextFromContext(ctx); spanContext.IsValid() {
		r.AddAttrs(
			slog.String("traceID", spanContext.TraceID().String()),
			slog.String("spanID", spanContext.SpanID().String()),
		)
	}
	return o.Handler.Handle(ctx, r)
}
