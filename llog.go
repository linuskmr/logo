package llog

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var (
	DebugMode = color.New(color.FgGreen).Sprint("DEBUG")
	InfoMode  = color.New(color.FgBlue).Sprint("INFO")
	WarnMode  = color.New(color.FgYellow).Sprint("WARN")
	ErrorMode = color.New(color.FgRed).Sprint("ERROR")
	PrintMode = "PRINT"
)

func spaceJoiner(v []interface{}) string {
	var out []string
	for _, elem := range v {
		out = append(out, fmt.Sprint(elem))
	}
	return strings.Join(out, " ")
}

func Info(v ...interface{}) {
	log(InfoMode, v)
}

func Error(v ...interface{}) {
	log(ErrorMode, v)
}

func Debug(v ...interface{}) {
	log(DebugMode, v)
}

func Print(v ...interface{}) {
	log(PrintMode, v)
}

func Warn(v ...interface{}) {
	log(WarnMode, v)
}

func log(mode string, v []interface{}) {
	output := NewEntry(mode, spaceJoiner(v)).ByteArr()
	output = append(output, []byte("\n")...)
	OutputWriter.Write(output)
}
