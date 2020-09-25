package llog

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

type Entry struct {
	Mode     Mode
	Date     string
	Time     string
	Millis   string
	Msg      string
	Filename string
	Funcname string
}

func NewEntry(mode Mode, msg string) *Entry {
	entry := Entry{
		Msg: msg,
	}

	// Add date and time
	timeNow := time.Now()
	entry.Date = timeNow.Format(dateFormat)
	entry.Time = timeNow.Format(timeFormat)

	// Add file and line
	caller, file, line, _ := runtime.Caller(2)
	functionName := runtime.FuncForPC(caller).Name()
	if Filename {
		entry.Filename = fmt.Sprintf("%s:%d", shortFilename(file), line)
	}
	if Funcname {
		entry.Funcname = shortFilename(functionName)
	}

	return &entry
}

func (entry *Entry) String() string {
	params := []string{modeText[entry.Mode]}
	if Date {
		params = append(params, entry.Date)
	}
	if Time {
		if Millis {
			params = append(params, entry.Time+entry.Millis)
		} else {
			params = append(params, entry.Time)
		}
	}
	if Filename {
		params = append(params, entry.Filename)
	}
	if Funcname {
		params = append(params, entry.Funcname)
	}

	return strings.Join(params, " ") + ":" + entry.Msg
}

func shortFilename(filename string) string {
	filePath := strings.Split(filename, "/")
	return filePath[len(filePath)-1]
}
