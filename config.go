package llog

import (
	"io"
	"os"
)

type Flag uint8

const (
	FlagDate = 1 << Flag(iota)
	FlagTime
	FlagMillis
	FlagFilename
	FlagFuncname
	FlagJson
)

var (
	Date                 = true
	Time                 = true
	Millis               = true
	Filename             = true
	Funcname             = true
	Json                 = false
	Output     io.Writer = os.Stdout
	DateFormat           = stdLongYear + "-" + stdZeroMonth + "-" + stdZeroDay
	TimeFormat           = stdHour + ":" + stdZeroMinute + ":" + stdZeroSecond
)

func Config(flags Flag) {
	Date = flags&FlagDate != 0
	Time = flags&FlagTime != 0
	Millis = flags&FlagMillis != 0
	Filename = flags&FlagFilename != 0
	Funcname = flags&FlagFuncname != 0
	Json = flags&FlagJson != 0
}

const (
	stdLongYear    = "2006"
	stdZeroMonth   = "01"
	stdZeroDay     = "02"
	stdHour        = "15"
	stdZeroMinute  = "04"
	stdZeroSecond  = "05"
	stdMillisecond = ".000"
)
