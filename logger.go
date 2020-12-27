package logo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Logger with various configuration options regarding how and where to output.
type Logger struct {
	// Date specifies whether the date should be displayed.
	Date bool
	// Time specifies whether the time should be displayed.
	Time bool
	// Millis specifies whether the milliseconds should be displayed.
	Millis bool
	// Filename specifies whether the filename and line should be displayed where the
	// Logger was called.
	Filename bool
	// Funcname specifies whether the struct and function should be displayed where
	// the Logger was called.
	Funcname bool
	// Json specifies whether the output should be in Json format.
	Json bool
	// Output is the writer to which the log should be written.
	Output io.Writer
	// DateFormat is the date format of the output.
	DateFormat string
	// TimeFormat is the time format of the output.
	TimeFormat string
	// Level is the log level of this logger. The logger only logs messages with a
	// level greater or equal to the log level. By default AllLevels is selected,
	// which logs everything.
	Level Level
}

// New creates a new Logger with the given flags.
func New(flag Flag) *Logger {
	logger := &Logger{
		Output:     os.Stdout,
		DateFormat: "2006-01-02",
		TimeFormat: "15:04:05",
	}
	logger.Config(flag)
	return logger
}

// Config configures the Logger according to the given flags.
func (l *Logger) Config(flag Flag) {
	l.Date = flag&DateFlag != 0
	l.Time = flag&TimeFlag != 0
	l.Millis = flag&MillisFlag != 0
	l.Filename = flag&FilenameFlag != 0
	l.Funcname = flag&FuncnameFlag != 0
	l.Json = flag&JsonFlag != 0
}

// entry creates a new entry for a this Logger.
func (l *Logger) newEntry(level Level, distance int, msg string) *entry {
	// Increment distance to get the right funcname and line number of the original
	// log call.
	distance++

	e := &entry{Level: level, Msg: msg}

	// Add date and time
	timeNow := time.Now()
	if l.Date {
		dateStr := timeNow.Format(l.DateFormat)
		e.Date = dateStr
	}
	if l.Time {
		timeStr := timeNow.Format(l.TimeFormat)
		e.Time = timeStr
	}
	if l.Millis {
		// Append millis to e.Time
		millisStr := timeNow.Format(".000")
		e.Time += millisStr
	}

	// Add file and line
	caller, file, line, _ := runtime.Caller(distance)
	if l.Filename {
		filenameStr := shortFilename(file) + ":" + strconv.Itoa(line)
		e.Filename = filenameStr
	}
	if l.Funcname {
		funcNameStr := shortFilename(runtime.FuncForPC(caller).Name())
		e.FuncName = funcNameStr
	}
	return e
}

// Info logs a message with InfoLevel.
func (l *Logger) Info(v ...interface{}) {
	l.doLog(InfoLevel, 1, v...)
}

// Error logs a message with ErrorLevel.
func (l *Logger) Error(v ...interface{}) {
	l.doLog(ErrorLevel, 1, v...)
}

// Debug logs a message with DebugLevel.
func (l *Logger) Debug(v ...interface{}) {
	l.doLog(DebugLevel, 1, v...)
}

// Print logs a message with PrintLevel.
func (l *Logger) Print(v ...interface{}) {
	l.doLog(PrintLevel, 1, v...)
}

// Warn logs a message with WarnLevel.
func (l *Logger) Warn(v ...interface{}) {
	l.doLog(WarnLevel, 1, v...)
}

// Log logs a message with the given Level.
func (l *Logger) Log(level Level, v ...interface{}) {
	l.doLog(level, 1, v...)
}

// doLog logs a message with a given log Level and the distance to the original
// call (needed for the filename and line number of the log message).
func (l *Logger) doLog(level Level, distance int, v ...interface{}) {
	// Do not log if the output would not be visible anyway or log level is to low
	if l.Output == nil || level < l.Level {
		return
	}

	// Increment distance to get the right funcname and line number of the original
	// log call.
	distance++

	// Format message in desired format
	var output []byte
	if l.Json {
		var err error
		output, err = json.MarshalIndent(l.newEntry(level, distance, spaceJoiner(v)), "  ", "")
		if err != nil {
			output = []byte(err.Error())
		}
	} else {
		output = []byte(l.newEntry(level, distance, spaceJoiner(v)).String())
	}
	output = append(output, []byte("\n")...)
	l.Output.Write(output)
}

// spaceJoiner converts the array entries to string and joins them with a
// whitespace.
func spaceJoiner(v []interface{}) string {
	var out []string
	for _, elem := range v {
		out = append(out, fmt.Sprint(elem))
	}
	return strings.Join(out, " ")
}
