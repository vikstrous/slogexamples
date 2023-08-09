package testoutputter_test

import (
	"log/slog"
	"strings"
	"testing"

	"github.com/vikstrous/slogexamples/testoutputter"
)

type TestT struct {
	*testing.T
	Output string
}

func (t *TestT) Logf(s string, args ...any) {
	t.Output = s
}

func TestOutputter(t *testing.T) {
	wrappedT := &TestT{T: t}
	o := testoutputter.NewTestOutputter(wrappedT, nil)
	logger := slog.New(slog.NewTextHandler(o, nil))
	allocsPerRun := testing.AllocsPerRun(1, func() {
		logger.Info("hello")
	})
	if allocsPerRun > 0 {
		t.Fatalf("extra allocations introduced %.0f", allocsPerRun)
	}
	if !strings.Contains(wrappedT.Output, "hello") {
		t.Fatal("Did not log hello")
	}
}

func Example() {
	var t testing.TB
	logger := slog.New(slog.NewTextHandler(testoutputter.NewTestOutputter(t, nil), nil))
	logger.Info("example")
}
