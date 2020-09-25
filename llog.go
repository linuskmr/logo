package llog

import (
	"github.com/fatih/color"
	"log"
	"os"
)

const flags = log.LstdFlags | log.Lshortfile

var (
	infoLog  = log.New(os.Stdout, color.New(color.FgBlue).Sprint("INFO  "), flags)
	printLog = log.New(os.Stdout, "PRINT ", flags)
	errorLog = log.New(os.Stdout, color.New(color.FgRed).Sprint("ERROR "), flags)
	debugLog = log.New(os.Stdout, color.New(color.FgGreen).Sprint("DEBUG "), flags)
)

func Info(v ...interface{}) {
	infoLog.Println(v...)
}

func Error(v ...interface{}) {
	errorLog.Println(v...)
}

func Debug(v ...interface{}) {
	debugLog.Println(v...)
}

func Print(v ...interface{}) {
	printLog.Println(v...)
}
