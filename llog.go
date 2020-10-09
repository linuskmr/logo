// With the package llog you can log. You can log in Json or as string to any
// io.Writer. There are different LogLevels (see Level) with corresponding
// methods.
package llog

var standard = New(DateFlag | TimeFlag | MillisFlag | FilenameFlag | FuncnameFlag)

// Logs a message with InfoLevel.
func Info(v ...interface{}) {
	standard.logDistance(InfoLevel, 1, v...)
}

// Logs a message with ErrorLevel.
func Error(v ...interface{}) {
	standard.logDistance(ErrorLevel, 1, v...)
}

// Logs a message with DebugLevel.
func Debug(v ...interface{}) {
	standard.logDistance(DebugLevel, 1, v...)
}

// Logs a message with PrintLevel.
func Print(v ...interface{}) {
	standard.logDistance(PrintLevel, 1, v...)
}

// Logs a message with WarnLevel.
func Warn(v ...interface{}) {
	standard.logDistance(WarnLevel, 1, v...)
}

// Logs a message with the given Level.
func Log(level Level, v ...interface{}) {
	standard.logDistance(level, 1, v...)
}

// Configures the Logger according to the given flag.
func Config(flag Flag) {
	standard.Config(flag)
}
