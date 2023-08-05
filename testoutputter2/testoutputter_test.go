package testoutputter2_test

import (
	"strings"
	"testing"

	"github.com/vikstrous/slogexamples/testoutputter2"
)

type TestT struct {
	*testing.T
	Output string
}

func (t *TestT) Logf(s string, args ...any) {
	t.Helper()
	t.Output = s
	t.T.Logf(s, args...)
}

func TestOutputter(t *testing.T) {
	wrappedT := &TestT{T: t}
	logger := testoutputter2.New(wrappedT)
	// We don't test allocs because there are a lot but it doesn't matter as much in tests
	logger.Info("hello")
	if !strings.Contains(wrappedT.Output, "hello") {
		t.Fatal("Did not log hello")
	}
	// There's no easy way to assert on what's printed, but expect this test to print "testputter_test.go:25" on the log line
}

func Example() {
	var t testing.TB
	logger := testoutputter2.New(t)
	logger.Info("example")
}
