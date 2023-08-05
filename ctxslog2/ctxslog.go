// Package ctxslog2 provides functions for logging that force the caller to pass the logger through the context.
// Even With and WithGroup functions return a context rather than a logger, ensuring that there is no direct access to the logger in functions that don't pass the context correctly.
// In codebases where context propagation is important, expected and required, this is a helpful constraint to enforce.
package ctxslog2

import (
	"context"

	"golang.org/x/exp/slog"
)

func Debug(ctx context.Context, msg string, attrs ...any) {
	get(ctx).DebugContext(ctx, msg, attrs...)
}

func Info(ctx context.Context, msg string, attrs ...any) {
	get(ctx).InfoContext(ctx, msg, attrs...)
}

func Warn(ctx context.Context, msg string, attrs ...any) {
	get(ctx).WarnContext(ctx, msg, attrs...)
}

func Error(ctx context.Context, msg string, attrs ...any) {
	get(ctx).ErrorContext(ctx, msg, attrs...)
}

func With(ctx context.Context, args ...any) context.Context {
	return Put(ctx, get(ctx).With(args...))
}

func WithGroup(ctx context.Context, group string) context.Context {
	return Put(ctx, get(ctx).WithGroup(group))
}

type key struct{}

func get(ctx context.Context) *slog.Logger {
	l := ctx.Value(key{})
	if l == nil {
		panic("no ctxslog found in context")
	}
	return l.(*slog.Logger)
}

func Put(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, key{}, l)
}
