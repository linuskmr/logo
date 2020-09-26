package llog

import (
	"encoding/json"
	"github.com/fatih/color"
	"strings"
)

// A log level
type Level uint8

const (
	// The log Level's following are sorted in ascending order of priority. AllLevels
	// = PrintLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel. So the highest
	// LogLevel is ErrorLevel, the lowest is AllLevels = PrintLevel. A message is only logged if
	// its level is greater or equal to the loglevel of the logger.
	PrintLevel = Level(iota)
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	AllLevels = PrintLevel
)

// Texts for the log Level's.
var levelTexts = [...]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PrintLevel: "PRINT",
}

// Colors for the Level's.
var levelColors = [...]*color.Color{
	DebugLevel: color.New(color.FgGreen).Add(color.Bold),
	InfoLevel:  color.New(color.FgBlue).Add(color.Bold),
	WarnLevel:  color.New(color.FgYellow).Add(color.Bold),
	ErrorLevel: color.New(color.FgRed).Add(color.Bold),
	PrintLevel: color.New(color.FgBlack).Add(color.Bold),
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

// Converts an Entry to a string, with the Level colored (See levelColors).
func (entry *Entry) String() string {
	params := []string{levelColors[entry.Level].Sprintf("%-5s", levelTexts[entry.Level])}
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
		Level: levelTexts[entry.Level],
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
