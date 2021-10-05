// Package logo contains an extension of the log package from the standard library.
// Its main features are:
// - Logging filename, function and line number where the logger was called.
// - Multiple log Level's
// - Easy start (You don't need to create a logger instance)
// - Colored output of log Level's via ANSI-Code
// - String and JSON output
// - Can write to any io.Writer
//
// Example:
// 		logo.Print("Hallo")
package logo

var (
	// Default is the Logger used by all functions that don't get called on a concrete logger instance.
	// For example, if you write logo.Info("Logo"), that uses this Default Logger.
	Default = New(DateFlag | TimeFlag | MillisFlag | FilenameFlag | FuncnameFlag)
)

// Info logs a message with InfoLevel.
func Info(v ...interface{}) {
	Default.doLog(InfoLevel, 1, v...)
}

// Error logs a message with ErrorLevel.
func Error(v ...interface{}) {
	Default.doLog(ErrorLevel, 1, v...)
}

// Debug logs a message with DebugLevel.
func Debug(v ...interface{}) {
	Default.doLog(DebugLevel, 1, v...)
}

// Print logs a message with PrintLevel.
func Print(v ...interface{}) {
	Default.doLog(PrintLevel, 1, v...)
}

// Warn logs a message with WarnLevel.
func Warn(v ...interface{}) {
	Default.doLog(WarnLevel, 1, v...)
}

// Log logs a message with the given Level.
func Log(level Level, v ...interface{}) {
	Default.doLog(level, 1, v...)
}

// Config configures the Default Logger according to the given flags.
func Config(flag Flag) {
	Default.Config(flag)
}
