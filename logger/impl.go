package log

import (
	"fmt"
	"log"
)

// logger implement the Logger interface.
type logger struct {
	defaulLogger *log.Logger
}

func NewLogger() Logger {
	return &logger{
		defaulLogger: log.Default(),
	}
}

// Infof log informational messages.
func (l *logger) LogInfo(v ...any) {
	l.defaulLogger.Printf("[Info]: %s\n", v)
}

// Error log error messages.
func (l *logger) LogError(v ...any) {
	l.defaulLogger.Printf("[Error]: %s\n", v)
}

// Errorf log error messages using format and additional arguments.
func (l *logger) LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)

	l.defaulLogger.Printf("[Error]: %s\n", msg)
}

// Infof log informational messages using format and additional arguments.
func (l *logger) LogInfof(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	l.defaulLogger.Printf("[Info]: %s\n", msg)
}
