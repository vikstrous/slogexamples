package testoutputter

import (
	"io"
	"testing"
)

type testOutputter struct {
	TB   testing.TB
	Next io.Writer
}

func (t testOutputter) Write(b []byte) (int, error) {
	// The last bytes is a newline that we don't want to include because tb.Log assumes that there's no trailing newline
	t.TB.Logf(string(b[:len(b)-1]))
	if t.Next == nil {
		return len(b), nil
	}
	return t.Next.Write(b)
}

// NewTestOutputter creates an io.Writer that can be used with slog's TextHandler or JSONHandler implementatiosn to redirect their output to the test's logs so they can be displayed correctly with concurrent tests, subtests, etc.
func NewTestOutputter(tb testing.TB, next io.Writer) io.Writer {
	return testOutputter{tb, next}
}
