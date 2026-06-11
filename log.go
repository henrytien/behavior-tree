package behaviortree

import (
	"fmt"
	"io"
)

// Logger is the sink for the library's optional diagnostic output.
//
// By default the library is silent: diagnostic messages are discarded so a
// behavior tree never pollutes the host application's stdout. Enable logging
// during development with SetLogOutput (e.g. SetLogOutput(os.Stderr)) or plug
// in a custom Logger with SetLogger.
type Logger interface {
	Printf(format string, args ...interface{})
}

// stdLogger writes formatted lines to an io.Writer.
type stdLogger struct {
	w io.Writer
}

func (l stdLogger) Printf(format string, args ...interface{}) {
	fmt.Fprintf(l.w, format+"\n", args...)
}

// nopLogger discards everything. It is the default.
type nopLogger struct{}

func (nopLogger) Printf(string, ...interface{}) {}

var logger Logger = nopLogger{}

// SetLogger installs a custom logger. Passing nil restores the silent default.
func SetLogger(l Logger) {
	if l == nil {
		logger = nopLogger{}
		return
	}
	logger = l
}

// SetLogOutput enables diagnostic logging to the given writer. Passing nil
// disables logging (restores the silent default).
func SetLogOutput(w io.Writer) {
	if w == nil {
		logger = nopLogger{}
		return
	}
	logger = stdLogger{w: w}
}

// Logf emits a diagnostic line through the configured logger. It is a no-op
// unless logging has been enabled via SetLogger / SetLogOutput.
func Logf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}
