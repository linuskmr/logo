package llog

import (
	"encoding/json"
	"github.com/fatih/color"
	"strings"
)

// A log Level
type Level uint8

const (
	// The log Level's following are sorted in ascending order of priority.
	// DebugLevel(AllLevels), InfoLevel, WarnLevel, ErrorLevel and PrintLevel. So
	// the highest Level is PrintLevel, the lowest is DebugLevel(PrintLevel). A
	// Entry is only logged if its level is greater or equal to the Level of the
	// Logger.
	DebugLevel = Level(iota)
	InfoLevel
	WarnLevel
	ErrorLevel
	PrintLevel
	AllLevels = DebugLevel
)

// Text and color for each log Level.
var levels = [...]struct {
	// The text of the Level.
	text string

	// The color of the level.
	color *color.Color
}{
	DebugLevel: {
		text:  "DEBUG",
		color: color.New(color.FgGreen).Add(color.Bold),
	},
	InfoLevel: {
		text:  "INFO",
		color: color.New(color.FgBlue).Add(color.Bold),
	},
	WarnLevel: {
		text:  "WARN",
		color: color.New(color.FgYellow).Add(color.Bold),
	},
	ErrorLevel: {
		text:  "ERROR",
		color: color.New(color.FgRed).Add(color.Bold),
	},
	PrintLevel: {
		text:  "PRINT",
		color: color.New(color.FgBlack).Add(color.Bold),
	},
}

// A Logger Entry.
type Entry struct {
	// Information about the Entry.
	Level    Level   `json:"level"`
	Date     *string `json:"date,omitempty"`
	Time     *string `json:"time,omitempty"`
	Msg      *string `json:"msg,omitempty"`
	Filename *string `json:"filename,omitempty"`
	FuncName *string `json:"func_name,omitempty"`
}

// Converts an Entry to a string, with the Level colored (See levels).
func (entry *Entry) String() string {
	params := []string{levels[entry.Level].color.Sprintf("%-5s", levels[entry.Level].text)}
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
	return strings.Join(params, " ") + ": " + *entry.Msg
}

// Converts an Entry to a Json byte arr.
func (entry *Entry) Json() []byte {
	type Alias Entry
	jsonByteArr, err := json.Marshal(&struct {
		Level string `json:"level"`
		*Alias
	}{
		Level: levels[entry.Level].text,
		Alias: (*Alias)(entry),
	})
	if err != nil {
		return []byte("Could not convert to Json: " + entry.String())
	}
	return jsonByteArr
}

// Shortens the filename to the last /.
func shortFilename(filename string) string {
	filePath := strings.Split(filename, "/")
	return filePath[len(filePath)-1]
}
