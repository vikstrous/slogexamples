package ctxslog_test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/vikstrous/slogexamples/ctxslog"
	"golang.org/x/exp/slog"
)

func TestCtxSlog(t *testing.T) {
	// Pre-allocate the output buffer so the test doesn't cause extra allocations
	buf := bytes.NewBuffer(make([]byte, 1000))
	sl := slog.New(slog.NewTextHandler(buf, nil))
	ctx := ctxslog.PutNew(context.Background(), sl)
	allocsPerRun := testing.AllocsPerRun(1, func() {
		ctxslog.Get(ctx).Info("example")
	})
	if allocsPerRun > 0 {
		t.Fatalf("extra allocations introduced %.0f", allocsPerRun)
	}
	if !strings.Contains(buf.String(), "example") {
		t.Fatal("did not log")
	}
}

func ExampleGet() {
	ctx := ctxslog.PutNew(context.Background(), slog.New(slog.NewTextHandler(os.Stdout, nil)))
	ctxslog.Get(ctx).Info("example")
}
