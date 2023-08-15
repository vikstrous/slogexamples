package testerrorer2

import (
	"context"
	"encoding/json"
	"log/slog"
	"testing"

	"github.com/vikstrous/slogevent"
)

type testErrorer struct {
	TB testing.TB
}

func (t testErrorer) EventHandler(ctx context.Context, e slogevent.Event) {
	if e.Level >= slog.LevelError {
		attrs, _ := json.Marshal(e.Attrs)
		t.TB.Fatalf("%s; %s", e.Message, string(attrs))
	}
}

// NewTestErrorerHandler creates a handler that fails tests with a descriptive error any time an error log is about to be printed.
func NewTestErrorerHandler(tb testing.TB, next slog.Handler) slog.Handler {
	return slogevent.NewHandler(testErrorer{TB: tb}.EventHandler, next)
}
