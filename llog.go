package llog

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var modeText = [...]string{
	DebugMode: "DEBUG",
	InfoMode:  "INFO",
	WarnMode:  "WARN",
	ErrorMode: "ERROR",
	PrintMode: "PRINT",
}

var modeColors = [...]*color.Color{
	DebugMode: color.New(color.FgGreen),
	InfoMode:  color.New(color.FgBlue),
	WarnMode:  color.New(color.FgYellow),
	ErrorMode: color.New(color.FgRed),
	PrintMode: color.New(color.FgBlack),
}

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

func log(mode Mode, v []interface{}) {
	output := NewEntry(mode, spaceJoiner(v)).ByteArr()
	output = append(output, []byte("\n")...)
	OutputWriter.Write(output)
}
