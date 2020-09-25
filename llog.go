package llog

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var modeText = []string{
	DebugMode: color.New(color.FgGreen).Sprint("DEBUG"),
	InfoMode:  color.New(color.FgBlue).Sprint("INFO "),
	WarnMode:  color.New(color.FgYellow).Sprint("WARN "),
	ErrorMode: color.New(color.FgRed).Sprint("ERROR"),
	PrintMode: "PRINT",
}

type Mode uint8

const (
	DebugMode = Mode(iota)
	InfoMode
	WarnMode
	ErrorMode
	PrintMode
)

func spaceJoiner(v ...interface{}) string {
	var out []string
	for _, elem := range v {
		out = append(out, fmt.Sprint(elem))
	}
	return strings.Join(out, " ")
}

func Info(v ...interface{}) {
	fmt.Fprintln(Output, NewEntry(InfoMode, spaceJoiner(v...)))
}

func Error(v ...interface{}) {
	fmt.Fprintln(Output, NewEntry(InfoMode, spaceJoiner(v...)))
}

func Debug(v ...interface{}) {
	fmt.Fprintln(Output, NewEntry(InfoMode, spaceJoiner(v...)))
}

func Print(v ...interface{}) {
	fmt.Fprintln(Output, NewEntry(InfoMode, spaceJoiner(v...)))
}

func Warn(v ...interface{}) {
	fmt.Fprintln(Output, NewEntry(InfoMode, spaceJoiner(v...)))
}
