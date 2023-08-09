package testoutputter2

import (
	"bytes"
	"log/slog"
	"sync"
	"testing"
)

// TestOutputterSlog
type TestOutputterSlog struct {
	logger *slog.Logger
	tb     testing.TB
	// The lock and buffer are shared by all copies of the logger
	lock *sync.Mutex
	buf  *bytes.Buffer
}

func (t TestOutputterSlog) Debug(msg string, args ...any) {
	t.tb.Helper()
	t.log(slog.LevelDebug, msg, args...)
}

func (t TestOutputterSlog) Info(msg string, args ...any) {
	t.tb.Helper()
	t.log(slog.LevelInfo, msg, args...)
}

func (t TestOutputterSlog) Warn(msg string, args ...any) {
	t.tb.Helper()
	t.log(slog.LevelWarn, msg, args...)
}

func (t TestOutputterSlog) Error(msg string, args ...any) {
	t.tb.Helper()
	t.log(slog.LevelError, msg, args...)
}

func (t TestOutputterSlog) log(level slog.Level, msg string, args ...any) {
	t.tb.Helper()
	t.lock.Lock()
	defer t.lock.Unlock()
	t.buf.Reset()
	t.logger.Log(nil, level, msg, args...)
	t.tb.Logf(t.buf.String())
}

func (t TestOutputterSlog) With(args ...any) *TestOutputterSlog {
	t2 := t
	t2.logger = t2.logger.With(args...)
	return &t2
}

func (t TestOutputterSlog) WithGroup(group string) *TestOutputterSlog {
	t2 := t
	t2.logger = t2.logger.WithGroup(group)
	return &t2
}

func New(tb testing.TB) *TestOutputterSlog {
	buf := new(bytes.Buffer)
	sl := slog.New(slog.NewTextHandler(buf, nil))
	return &TestOutputterSlog{logger: sl, tb: tb, lock: new(sync.Mutex), buf: buf}
}
