# slogexamples

`slogexamples` is a collection of examples showing how to extend slog. They are a follow up to https://medium.com/anchorage/three-logging-features-to-improve-your-slog-f72300a7fb66, showing some of the techniques mentioned briefly and some of the example logger features that can be implemented on top of slog. Three methods are mentioned in the article, but this repo covers a fourth one too.

All of these examples stay as close as possible to the 0 allocations goal of slog. They also all have example usage code.

## Hooking into io.Writer

[testoutputter](./testoutputter/) shows how to intercept the logger's calls to the underlying io.Writer and do something useful. It sends all logs to `t.Log()`, which allows tests to output correctly.

One limitation of all attempts to do this with slog is that the correct call site can't be printed. See https://github.com/golang/go/issues/59928 for more details.

## Wrapping slog

[ctxslog](./ctxslog/) is an example of a slog wrapper. It exposes only the methods `Debug/Info/Warn/Error`, but it forces the caller to pass the logger through the context. This is a more restricted way to use slog, but it's slightly more convenient to use in codebases where the context is expected to be passed everywhere and tracing or cancelation is very important.

## Custom slog.Handler

[otelhandler](./otelhandler/) is an example of a handler that acts as a middleware and adds additional attributes to log entries. In particular, it adds TraceID and SpanID if logger is used within the context of an open telemetry trace. This allows for correlating logs and traces sent to different systems.

There's a lot more to writing custom handlers. This guide is very helpful. https://github.com/golang/example/tree/master/slog-handler-guide

## Bonus: Hooking into slog.HandlerOptions.ReplaceAttr

[testerrorer](./testerrorer/) hooks into the slog handler's ReplaceAttr callback. This function is called on every attribute before it's formatted for rendering and testerrorer uses the opportunity to check if anything is logged at error level and fail the test if so. One limitation is that it looks at all attributes with the type `slog.Level` that are logged rather than the actual level of the log. This is to make sure that if the level attribute is renamed by another ReplaceAttr function that wraps this one, that doesn't break the functionality. The down side is that a call like `logger.Info("example", "l", slog.LevelError)` will cause the test to error incorrectly.