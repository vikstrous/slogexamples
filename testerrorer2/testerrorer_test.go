package testerrorer2_test

import (
	"context"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/vikstrous/slogexamples/testerrorer2"
)

type TestT struct {
	*testing.T
	DidError bool
}

func (t *TestT) Fatalf(s string, args ...any) {
	t.DidError = true
}

func TestErrorerErrors(t *testing.T) {
	wrappedT := &TestT{T: t}
	logger := slog.New(testerrorer2.NewTestErrorerHandler(wrappedT, slog.NewJSONHandler(io.Discard, nil)))
	ctx := context.Background()
	allocsPerRun := testing.AllocsPerRun(1, func() {
		logger.InfoContext(ctx, "example")
	})
	if allocsPerRun > 1 {
		t.Fatalf("extra allocations introduced in info path %.0f", allocsPerRun)
	}
	if wrappedT.DidError {
		t.Fatal("error below error level")
	}
	allocsPerRun = testing.AllocsPerRun(1, func() {
		logger.Error("example")
	})
	if !wrappedT.DidError {
		t.Fatal("did not error at error level")
	}
	if allocsPerRun > 7 {
		t.Fatalf("extra allocations introduced in error path: %.0f", allocsPerRun)
	}
	wrappedT.DidError = false
	logger.Log(ctx, slog.LevelError+1, "example")
	if !wrappedT.DidError {
		t.Fatal("did not error above error level")
	}
	wrappedT.DidError = false
	logger.Info("example", "l", slog.LevelError)
	if wrappedT.DidError {
		t.Fatal("error level attribute in list triggered a test fialure")
	}
}

func Example() {
	var t *testing.T
	logger := slog.New(testerrorer2.NewTestErrorerHandler(t, slog.NewTextHandler(os.Stderr, nil)))
	logger.Info("use the logger as normal")
}
