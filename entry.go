package llog

import (
	"encoding/json"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Mode     *string `json:"mode,omitempty"`
	Date     *string `json:"date,omitempty"`
	Time     *string `json:"time,omitempty"`
	Msg      *string `json:"msg,omitempty"`
	Filename *string `json:"filename,omitempty"`
	FuncName *string `json:"func_name,omitempty"`
}

func NewEntry(mode string, msg string) *Entry {
	entry := Entry{Mode: &mode, Msg: &msg}

	// Add date and time
	timeNow := time.Now()
	if Date {
		dateStr := timeNow.Format(DateFormat)
		entry.Date = &dateStr
	}
	if Time {
		timeStr := timeNow.Format(TimeFormat)
		entry.Time = &timeStr
	}
	if Millis {
		millisStr := timeNow.Format(stdMillisecond)
		if entry.Time == nil {
			emptyStr := ""
			entry.Time = &emptyStr
		}
		timeStr := *entry.Time + millisStr
		entry.Time = &timeStr
	}

	// Add file and line
	caller, file, line, _ := runtime.Caller(3)
	if Filename {
		filenameStr := shortFilename(file) + ":" + strconv.Itoa(line)
		entry.Filename = &filenameStr
	}
	if Funcname {
		funcNameStr := shortFilename(runtime.FuncForPC(caller).Name())
		entry.FuncName = &funcNameStr
	}

	return &entry
}

func (entry *Entry) String() []byte {
	params := make([]string, 6)
	params = append(params, *entry.Mode)
	if entry.Date != nil {
		params = append(params, *entry.Date)
	}
	if entry.Time != nil && *entry.Time != "" {
		params = append(params, *entry.Time)
	}
	if entry.Filename != nil {
		params = append(params, *entry.Filename)
	}
	if entry.FuncName != nil {
		params = append(params, *entry.FuncName)
	}

	return []byte(strings.Join(params, " ") + ": " + *entry.Msg)
}

func (entry *Entry) Json() []byte {
	jsonStr, err := json.Marshal(entry)
	if err != nil {
		return []byte("Could not convert to Json: " + string(entry.String()))
	}
	return jsonStr
}

func (entry *Entry) ByteArr() []byte {
	if Json {
		return entry.Json()
	} else {
		return entry.String()
	}
}

func shortFilename(filename string) string {
	filePath := strings.Split(filename, "/")
	return filePath[len(filePath)-1]
}