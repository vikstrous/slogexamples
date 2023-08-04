package ctxslog2_test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/vikstrous/slogexamples/ctxslog2"
	"golang.org/x/exp/slog"
)

func TestCtxSlog2(t *testing.T) {
	// Pre-allocate the output buffer so the test doesn't cause extra allocations
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	sl := slog.New(slog.NewTextHandler(buf, nil))
	ctx := ctxslog2.Put(context.Background(), sl)
	allocsPerRun := testing.AllocsPerRun(1, func() {
		ctxslog2.Info(ctx, "example")
	})
	if allocsPerRun > 0 {
		t.Fatalf("extra allocations introduced %.0f", allocsPerRun)
	}
	if !strings.Contains(buf.String(), "example") {
		t.Fatal("did not log")
	}
	ctx = ctxslog2.With(ctx, "key", "value")
	allocsPerRun = testing.AllocsPerRun(1, func() {
		ctxslog2.Info(ctx, "example")
	})
	if allocsPerRun > 0 {
		t.Fatalf("extra allocations introduced %.0f", allocsPerRun)
	}
	if !strings.Contains(buf.String(), "value") {
		t.Fatal("did not log the value from With")
	}
}

func Example() {
	ctx := ctxslog2.Put(context.Background(), slog.New(slog.NewTextHandler(os.Stdout, nil)))
	ctxslog2.Info(ctx, "example")
}
