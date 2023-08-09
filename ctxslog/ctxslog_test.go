package ctxslog_test

import (
	"bytes"
	"context"
	"log/slog"
	"os"
	"strings"
	"testing"

	"github.com/vikstrous/slogexamples/ctxslog"
)

func TestCtxSlog(t *testing.T) {
	ctx := context.Background()
	// Pre-allocate the output buffer so the test doesn't cause extra allocations
	buf := bytes.NewBuffer(make([]byte, 0, 1000))
	sl := slog.New(slog.NewTextHandler(buf, nil))
	l := ctxslog.New(sl)
	allocsPerRun := testing.AllocsPerRun(1, func() {
		l.Info(ctx, "example")
	})
	if allocsPerRun > 0 {
		t.Fatalf("extra allocations introduced %.0f", allocsPerRun)
	}
	if !strings.Contains(buf.String(), "example") {
		t.Fatal("did not log")
	}
}

func Example() {
	ctx := context.Background()
	l := ctxslog.New(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	l.Info(ctx, "example")
}
