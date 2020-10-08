//
package llog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Logger struct {
	// Specifies whether the date should be displayed.
	Date bool

	// Specifies whether the time should be displayed.
	Time bool

	// Specifies whether the milliseconds should be displayed.
	Millis bool

	// Specifies whether the filename and line should be displayed that was logged.
	Filename bool

	// Specifies whether the struct and function should be displayed that was logged.
	Funcname bool

	// Specifies whether the output should be in Json format.
	Json bool

	// The writer to which the log should be written.
	OutputWriter io.Writer

	// The date format of the output.
	DateFormat string

	// The time format of the output.
	TimeFormat string

	// The log level of this logger. It only logs messages with a level greater or
	// equal to the log level. To log everything, choose AllLevels. By default AllLevels is selected.
	Level Level
}

// A flag with which the logger can be configured.
type Flag uint8

const (
	// Flag to configure the equivalent property in the logger.
	DateFlag = 1 << Flag(iota)
	TimeFlag
	MillisFlag
	FilenameFlag
	FuncnameFlag
	JsonFlag
)

// Creates a new Logger with the given flag.
func New(flag Flag) *Logger {
	newLogger := Logger{
		OutputWriter: os.Stdout,
		DateFormat:   "2006-01-02",
		TimeFormat:   "15:04:05",
	}
	newLogger.Config(flag)
	return &newLogger
}

// Configures the Logger according to the given flag.
func (l *Logger) Config(flag Flag) {
	l.Date = flag&DateFlag != 0
	l.Time = flag&TimeFlag != 0
	l.Millis = flag&MillisFlag != 0
	l.Filename = flag&FilenameFlag != 0
	l.Funcname = flag&FuncnameFlag != 0
	l.Json = flag&JsonFlag != 0
}

// Creates a new Entry for a this Logger.
func (l *Logger) entry(level Level, distance int, msg string) *Entry {
	distance++
	entry := Entry{Level: level, Msg: &msg}

	// Add date and time
	timeNow := time.Now()
	if l.Date {
		dateStr := timeNow.Format(l.DateFormat)
		entry.Date = &dateStr
	}
	if l.Time {
		timeStr := timeNow.Format(l.TimeFormat)
		entry.Time = &timeStr
	}
	if l.Millis {
		millisStr := timeNow.Format(".000")
		if entry.Time == nil {
			// Set entry.Time to empty string if it was nil to be able to attach milliseconds.
			emptyStr := ""
			entry.Time = &emptyStr
		}
		timeStr := *entry.Time + millisStr
		entry.Time = &timeStr
	}

	// Add file and line
	caller, file, line, _ := runtime.Caller(distance)
	if l.Filename {
		filenameStr := shortFilename(file) + ":" + strconv.Itoa(line)
		entry.Filename = &filenameStr
	}
	if l.Funcname {
		funcNameStr := shortFilename(runtime.FuncForPC(caller).Name())
		entry.FuncName = &funcNameStr
	}

	return &entry
}

// Logs a message with InfoLevel.
func (l *Logger) Info(v ...interface{}) {
	l.logDistance(InfoLevel, 1, v...)
}

// Logs a message with ErrorLevel.
func (l *Logger) Error(v ...interface{}) {
	l.logDistance(ErrorLevel, 1, v...)
}

// Logs a message with DebugLevel.
func (l *Logger) Debug(v ...interface{}) {
	l.logDistance(DebugLevel, 1, v...)
}

// Logs a message with PrintLevel.
func (l *Logger) Print(v ...interface{}) {
	l.logDistance(PrintLevel, 1, v...)
}

// Logs a message with WarnLevel.
func (l *Logger) Warn(v ...interface{}) {
	l.logDistance(WarnLevel, 1, v...)
}

// Logs a message with the given Level.
func (l *Logger) Log(level Level, v ...interface{}) {
	l.logDistance(level, 1, v...)
}

// Logs a message with a given log level and the distance to the original
// call (needed for the filename and line number of the log message).
func (l *Logger) logDistance(level Level, distance int, v ...interface{}) {
	if l.OutputWriter == nil || level < l.Level {
		return
	}
	distance++
	var output []byte

	// Format message in desired format
	if l.Json {
		output = l.entry(level, distance, spaceJoiner(v)).Json()
	} else {
		output = []byte(l.entry(level, distance, spaceJoiner(v)).String())
	}
	output = append(output, []byte("\n")...)
	l.OutputWriter.Write(output)
}

// Gets the string representation of an array of interface{} and joins it with spaces.
func spaceJoiner(v []interface{}) string {
	var out []string
	for _, elem := range v {
		out = append(out, fmt.Sprint(elem))
	}
	return strings.Join(out, " ")
}
