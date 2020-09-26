package llog

var Standard = New(DateFlag | TimeFlag | MillisFlag | FilenameFlag | FuncnameFlag)

// Logs a message with InfoLevel.
func Info(v ...interface{}) {
	Standard.logDistance(InfoLevel, 1, v...)
}

// Logs a message with ErrorLevel.
func Error(v ...interface{}) {
	Standard.logDistance(ErrorLevel, 1, v...)
}

// Logs a message with DebugLevel.
func Debug(v ...interface{}) {
	Standard.logDistance(DebugLevel, 1, v...)
}

// Logs a message with PrintLevel.
func Print(v ...interface{}) {
	Standard.logDistance(PrintLevel, 1, v...)
}

// Logs a message with WarnLevel.
func Warn(v ...interface{}) {
	Standard.logDistance(WarnLevel, 1, v...)
}

// Logs a message with the given Level.
func Log(level Level, v ...interface{}) {
	Standard.logDistance(level, 1, v...)
}
