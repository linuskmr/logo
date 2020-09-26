package llog

import (
	"encoding/json"
	"github.com/fatih/color"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Mode uint8

const (
	DebugMode = Mode(iota)
	InfoMode
	WarnMode
	ErrorMode
	PrintMode
)

var modeText = [...]string{
	DebugMode: "DEBUG",
	InfoMode:  "INFO",
	WarnMode:  "WARN",
	ErrorMode: "ERROR",
	PrintMode: "PRINT",
}

var modeColors = [...]*color.Color{
	DebugMode: color.New(color.FgGreen).Add(color.Bold),
	InfoMode:  color.New(color.FgBlue).Add(color.Bold),
	WarnMode:  color.New(color.FgYellow).Add(color.Bold),
	ErrorMode: color.New(color.FgRed).Add(color.Bold),
	PrintMode: color.New(color.FgBlack).Add(color.Bold),
}

type Entry struct {
	Mode     Mode    `json:"mode"`
	Date     *string `json:"date,omitempty"`
	Time     *string `json:"time,omitempty"`
	Msg      *string `json:"msg,omitempty"`
	Filename *string `json:"filename,omitempty"`
	FuncName *string `json:"func_name,omitempty"`
}

func NewEntry(mode Mode, msg string) *Entry {
	entry := Entry{Mode: mode, Msg: &msg}

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
	params := []string{modeColors[entry.Mode].Sprintf("%-5s", modeText[entry.Mode])}
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
	type Alias Entry
	jsonByteArr, err := json.Marshal(&struct {
		Mode string `json:"mode"`
		*Alias
	}{
		Mode:  modeText[entry.Mode],
		Alias: (*Alias)(entry),
	})
	if err != nil {
		return []byte("Could not convert to Json: " + string(entry.String()))
	}
	return jsonByteArr
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
