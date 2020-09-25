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
	stdMillisecond = ".000"
)

const (
	dateFormat = stdLongYear + "-" + stdZeroMonth + "-" + stdZeroDay
	timeFormat = stdHour + ":" + stdZeroMinute + ":" + stdZeroSecond
)

const (
	FlagDate = 1 << iota
	FlagTime
	FlagMillis
	FlagFilename
	FlagFuncName
)

var (
	Output io.Writer = os.Stdout
	Flags            = FlagDate | FlagTime | FlagMillis | FlagFilename | FlagFuncName
)

func shortFilename(filename string) string {
	filePath := strings.Split(filename, "/")
	return filePath[len(filePath)-1]
}

func currentTime() string {
	var output strings.Builder
	timeNow := time.Now()
	if Flags&FlagDate != 0 && Flags&FlagTime != 0 {
		output.WriteString(timeNow.Format(dateFormat + " " + timeFormat))
	} else if Flags&FlagDate != 0 {
		output.WriteString(timeNow.Format(dateFormat))
	} else if Flags&FlagTime != 0 {
		output.WriteString(timeNow.Format(timeFormat))
	}
	if Flags&FlagMillis != 0 {
		output.WriteString(timeNow.Format(stdMillisecond))
	}
	return output.String()
}

func header(mode string) string {
	caller, file, line, _ := runtime.Caller(2)
	functionName := runtime.FuncForPC(caller).Name()

	var headers []string
	headers = append(headers, mode)
	headers = append(headers, currentTime())
	if Flags&FlagFilename != 0 {
		headers = append(headers, fmt.Sprintf("%s:%d", shortFilename(file), line))
	}
	if Flags&FlagFuncName != 0 {
		headers = append(headers, shortFilename(functionName))
	}
	return strings.Join(headers, " ") + ":"
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
