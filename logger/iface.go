package log

// Logger interface defining methods for logging.
type Logger interface {
	// Errorf logs error information using format and additional arguments.
	LogErrorf(format string, v ...any)

	// Infof logs informational message using format and additional arguments.
	LogInfof(format string, v ...any)

	// Error logs error information.
	LogError(v ...any)

	// Info logs information.
	LogInfo(v ...any)
}
