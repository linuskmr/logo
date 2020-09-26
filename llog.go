package llog

import (
	"fmt"
	"strings"
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

func log(mode Mode, v []interface{}) {
	output := NewEntry(mode, spaceJoiner(v)).ByteArr()
	output = append(output, []byte("\n")...)
	OutputWriter.Write(output)
}
