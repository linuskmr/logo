package llog

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"runtime"
	"strings"
)

const flags = log.LstdFlags | log.Lshortfile

var (
	infoPrefix  = color.New(color.FgBlue).Sprint("INFO")
	printPrefix = "PRINT"
	errorPrefix = color.New(color.FgRed).Sprint("ERROR ")
	debugPrefix = color.New(color.FgGreen).Sprint("DEBUG ")
)

func shortFilename(filename string) string {
	filePath := strings.Split(filename, "/")
	return filePath[len(filePath)-1]
}

func header(mode string) string {
	caller, file, line, _ := runtime.Caller(2)
	functionName := runtime.FuncForPC(caller).Name()
	return fmt.Sprintf("%-5s %s:%d %s:", mode, shortFilename(file), line, shortFilename(functionName))
}

func Info(v ...interface{}) {
	fmt.Fprintln(os.Stdout, header(infoPrefix), v)
}

func Error(v ...interface{}) {
	fmt.Fprintln(os.Stdout, header(errorPrefix), v)
}

func Debug(v ...interface{}) {
	fmt.Fprintln(os.Stdout, header(debugPrefix), v)
}

func Print(v ...interface{}) {
	fmt.Fprintln(os.Stdout, header(printPrefix), v)
}
