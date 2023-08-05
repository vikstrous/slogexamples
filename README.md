# slogexamples

`slogexamples` is a collection of examples showing how to extend slog. They are a follow up to [this blog post on the Anchorage Digital blog](https://medium.com/anchorage/three-logging-features-to-improve-your-slog-f72300a7fb66), showing some of the techniques for extending `slog` mentioned there.

All of these examples stay as close as possible to the 0 allocations goal of slog. They also all have example usage code.

The [docs](https://pkg.go.dev/github.com/vikstrous/slogexamples) have the usage examples and API docs rendered nicely. Navigate into each package directory to see the code and the usage examples.

## Hooking into io.Writer

[testoutputter](https://github.com/vikstrous/slogexamples/blob/master/testoutputter/testoutputter.go) shows how to intercept the logger's calls to the underlying io.Writer and do something useful. It sends all logs to `t.Log()`, which ensures that test output is readable when using parallel tests, subtests or when one test of many fails.

One limitation of most attempts to use `t.Log()` with slog is that the correct call site can't be printed. See [this issue](https://github.com/golang/go/issues/59928) for more details. The only way to correctly redirect logs to `t.Log()` is to use a wrapper around slog that calls `t.Log()` outside slog's code. An example of this is provided in [testoutputter2](https://github.com/vikstrous/slogexamples/blob/master/testoutputter2/testoutputter.go), which uses a wrapper around slog to do this. There are some obvious down sides of this approach, so I would personally prefer wrong line numbers over the `testoutputter2` solution.

## Wrapping slog

[ctxslog](https://github.com/vikstrous/slogexamples/blob/master/ctxslog/ctxslog.go) is an example of a slog wrapper. It forces the caller to pass the context in every logger call. This is a more restricted way to use slog, but it's slightly more convenient to use in codebases where the context is expected to be passed everywhere and tracing or cancelation is very important.

[ctxslog2](https://github.com/vikstrous/slogexamples/blob/master/ctxslog2/ctxslog.go) is an even more restrictive wrapper that forces the logger itself to be passed through the context. It hides all direct access to the logger and requires the user to call functions of the package rather than logger methods. This looks and feels like using a global logger instance, but the logger is actually in the context, which is better than a global logger because it can be faked in tests.

## Custom slog.Handler

[otelhandler](https://github.com/vikstrous/slogexamples/blob/master/otelhandler/otelhandler.go) is an example of a handler that acts as a middleware and adds additional attributes to log entries. In particular, it adds `TraceID` and `SpanID` to logs emitted within the context of an open telemetry trace. This allows for correlating logs and traces sent to different systems. See the [original blog post](https://medium.com/anchorage/three-logging-features-to-improve-your-slog-f72300a7fb66) for screenshots of what this looks like in Google Cloud.

There's a lot more to writing custom handlers than what's shown here. [This guide](https://github.com/golang/example/tree/master/slog-handler-guide), from the author of slog, is very helpful.

## Bonus: Hooking into slog.HandlerOptions.ReplaceAttr

[testerrorer](https://github.com/vikstrous/slogexamples/blob/master/testerrorer/testerrorer.go) hooks into the `slog.TextHandler`'s `ReplaceAttr` callback. This function is called on every attribute before it's formatted for rendering and `testerrorer` uses the opportunity to check if anything is logged at error level and fail the test if so.

One limitation is that it looks at all attributes with the type `slog.Level` that are logged rather than just the level of the log. This is to make sure that if the level attribute is renamed by another `ReplaceAttr` function, that doesn't break the functionality. The down side is that a call like `logger.Info("example", "l", slog.LevelError)` will cause the test to error incorrectly. The implementation can be trivially modified to look for the level attributed based on its name instead if the name is considered more stable.