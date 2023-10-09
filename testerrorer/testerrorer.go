// Package testerrorer shows how to use slog.HandlerOptions.ReplaceAttr to fail tests when errors are logged.
package testerrorer

import (
	"log/slog"
	"testing"
)

type testErrorer struct {
	TB   testing.TB
	Next func(groups []string, a slog.Attr) slog.Attr
}

func (t *testErrorer) replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if groups == nil && a.Key == slog.LevelKey && a.Value.Kind() == slog.KindAny {
		if level, ok := a.Value.Any().(slog.Level); ok {
			if level >= slog.LevelError {
				t.TB.Errorf("An error was logged.")
			}

			// Preserve 0 allocations using the strategy described in https://github.com/golang/go/issues/61774#issuecomment-1750763596
			a.Value = slog.StringValue(level.String())
		}
	}
	if t.Next == nil {
		return a
	}
	return t.Next(groups, a)
}

// NewTestErrorer creates a function that can be used as slog's ReplaceAttr function in the standard library implementations of TextHandler or JSONHandler
func NewTestErrorer(tb testing.TB, next func(groups []string, a slog.Attr) slog.Attr) func(groups []string, a slog.Attr) slog.Attr {
	return (&testErrorer{tb, next}).replaceAttr
}
