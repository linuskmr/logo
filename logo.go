// With the package llog you can log. You can log in MarshalJSON or as string to any
// io.Writer. There are different LogLevels (see Level) with corresponding
// methods.
package logo

var Standard = New(DateFlag | TimeFlag | MillisFlag | FilenameFlag | FuncnameFlag)

// Info logs a message with InfoLevel.
func Info(v ...interface{}) {
	Standard.doLog(InfoLevel, 1, v...)
}

// Error logs a message with ErrorLevel.
func Error(v ...interface{}) {
	Standard.doLog(ErrorLevel, 1, v...)
}

// Debug logs a message with DebugLevel.
func Debug(v ...interface{}) {
	Standard.doLog(DebugLevel, 1, v...)
}

// Print logs a message with PrintLevel.
func Print(v ...interface{}) {
	Standard.doLog(PrintLevel, 1, v...)
}

// Warn logs a message with WarnLevel.
func Warn(v ...interface{}) {
	Standard.doLog(WarnLevel, 1, v...)
}

// Log logs a message with the given Level.
func Log(level Level, v ...interface{}) {
	Standard.doLog(level, 1, v...)
}

// Config configures the Standard Logger according to the given flags.
func Config(flag Flag) {
	Standard.Config(flag)
}
