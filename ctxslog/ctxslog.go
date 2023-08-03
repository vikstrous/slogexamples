package ctxslog

import (
	"context"

	"golang.org/x/exp/slog"
)

// loggerWithCtx is a simplified API to slow that doesn't support With and WithGroup and supports only four log levels, but it forces the user to pass the logger through the context and ensures that calls to the logger always pass a context.
// The only way to create one is with PutNew and the only way to use it is with statements like ctxslog.Get(ctx).Info("example")
type loggerWithCtx struct {
	logger *slog.Logger
	ctx    context.Context
}

func (l loggerWithCtx) Debug(msg string, attrs ...any) {
	l.logger.DebugContext(l.ctx, msg, attrs...)
}

func (l loggerWithCtx) Info(msg string, attrs ...any) {
	l.logger.InfoContext(l.ctx, msg, attrs...)
}

func (l loggerWithCtx) Warn(msg string, attrs ...any) {
	l.logger.WarnContext(l.ctx, msg, attrs...)
}

func (l loggerWithCtx) Error(msg string, attrs ...any) {
	l.logger.ErrorContext(l.ctx, msg, attrs...)
}

type key struct{}

func Get(ctx context.Context) loggerWithCtx {
	l := ctx.Value(key{})
	if l == nil {
		panic("no ctxslog found in context")
	}
	sl := l.(*slog.Logger)
	return loggerWithCtx{logger: sl, ctx: ctx}
}

func PutNew(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, key{}, l)
}
