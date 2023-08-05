package ctxslog

import (
	"context"

	"golang.org/x/exp/slog"
)

// CtxSlog is a logger with an API that requires a context when logging
type CtxSlog struct {
	logger *slog.Logger
}

func (l CtxSlog) Debug(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

func (l CtxSlog) Info(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

func (l CtxSlog) Warn(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

func (l CtxSlog) Error(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
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
