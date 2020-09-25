package llog

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	infoPrefix  = color.New(color.FgBlue).Sprint("INFO ")
	printPrefix = "PRINT"
	errorPrefix = color.New(color.FgRed).Sprint("ERROR")
	debugPrefix = color.New(color.FgGreen).Sprint("DEBUG")
)

const (
	stdLongYear  = "2006"
	stdZeroMonth = "01"
	stdZeroDay   = "02"

	stdHour        = "15"
	stdZeroMinute  = "04"
	stdZeroSecond  = "05"
	stdMillisecond = "000"
)

const (
	dateFormat = stdLongYear + "-" + stdZeroMonth + "-" + stdZeroDay
	timeFormat = stdHour + ":" + stdZeroMinute + ":" + stdZeroSecond
)

const (
	Fdate = 1 << iota
	Ftime
)

var (
	Output io.Writer = os.Stdout
	Flags            = Fdate | Ftime
)

func shortFilename(filename string) string {
	filePath := strings.Split(filename, "/")
	return filePath[len(filePath)-1]
}

func currentTime() string {
	if Flags&Fdate != 0 && Flags&Ftime != 0 {
		return time.Now().Format(dateFormat + " " + timeFormat)
	}
	if Flags&Fdate != 0 {
		return time.Now().Format(dateFormat)
	}
	if Flags&Ftime != 0 {
		return time.Now().Format(timeFormat)
	}
	return ""
}

func header(mode string) string {
	caller, file, line, _ := runtime.Caller(2)
	functionName := runtime.FuncForPC(caller).Name()
	return fmt.Sprintf(`%s %s %s:%d %s:`, mode, currentTime(), shortFilename(file), line, shortFilename(functionName))
}

func spaceJoiner(v ...interface{}) string {
	var out []string
	for _, elem := range v {
		out = append(out, fmt.Sprint(elem))
	}
	return strings.Join(out, " ")
}

func Info(v ...interface{}) {
	fmt.Fprintln(Output, header(infoPrefix), spaceJoiner(v...))
}

func Error(v ...interface{}) {
	fmt.Fprintln(Output, header(errorPrefix), spaceJoiner(v...))
}

func Debug(v ...interface{}) {
	fmt.Fprintln(Output, header(debugPrefix), spaceJoiner(v...))
}

func Print(v ...interface{}) {
	fmt.Fprintln(Output, header(printPrefix), spaceJoiner(v...))
}
