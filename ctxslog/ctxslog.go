package ctxslog

import (
	"context"

	"golang.org/x/exp/slog"
)

// CtxSlog is a simplified API that doesn't support With and WithGroup and supports only four log levels, but it ensures that calls to the logger always pass a context.
type CtxSlog struct {
	logger *slog.Logger
}

func (l CtxSlog) Debug(ctx context.Context, msg string, attrs ...any) {
	l.logger.DebugContext(ctx, msg, attrs...)
}

func (l CtxSlog) Info(ctx context.Context, msg string, attrs ...any) {
	l.logger.InfoContext(ctx, msg, attrs...)
}

func (l CtxSlog) Warn(ctx context.Context, msg string, attrs ...any) {
	l.logger.WarnContext(ctx, msg, attrs...)
}

func (l CtxSlog) Error(ctx context.Context, msg string, attrs ...any) {
	l.logger.ErrorContext(ctx, msg, attrs...)
}

func (l CtxSlog) With(args ...any) *CtxSlog {
	return &CtxSlog{logger: l.logger.With(args...)}
}

func (l CtxSlog) WithGroup(group string) *CtxSlog {
	return &CtxSlog{logger: l.logger.WithGroup(group)}
}

func New(l *slog.Logger) *CtxSlog {
	return &CtxSlog{logger: l}
}
