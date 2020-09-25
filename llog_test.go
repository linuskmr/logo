package llog

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode           = "INFO"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "llog_test.go"
		method         = "llog.TestInfo"
	)
	Info(mode)
	assert.Regexp(t, "("+mode+")\\s+"+dateTimeMillis+"\\s("+filename+"):\\d+\\s("+method+": "+mode+")", stringBuilder.String())
}

func TestDebug(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode           = "DEBUG"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "llog_test.go"
		method         = "llog.TestDebug"
	)
	Debug(mode)
	assert.Regexp(t, "("+mode+")\\s+"+dateTimeMillis+"\\s("+filename+"):\\d+\\s("+method+": "+mode+")", stringBuilder.String())
}

func TestError(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode           = "ERROR"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "llog_test.go"
		method         = "llog.TestError"
	)
	Error(mode)
	assert.Regexp(t, "("+mode+")\\s+"+dateTimeMillis+"\\s("+filename+"):\\d+\\s("+method+": "+mode+")", stringBuilder.String())
}

func TestPrint(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode           = "PRINT"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "llog_test.go"
		method         = "llog.TestPrint"
	)
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+dateTimeMillis+"\\s("+filename+"):\\d+\\s("+method+": "+mode+")", stringBuilder.String())
}

func TestFlags(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode     = "PRINT"
		date     = "\\d\\d\\d\\d-\\d\\d-\\d\\d"
		time     = "\\d\\d:\\d\\d:\\d\\d"
		millis   = ".\\d\\d\\d"
		filename = "llog_test.go"
		method   = "llog.TestPrint"
	)
	Flags = FlagDate
	Print(mode)
	fmt.Println(stringBuilder.String())
	assert.Regexp(t, "("+mode+")\\s+"+date+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()

	Flags = FlagTime
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+time+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()

	Flags = FlagMillis
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+millis+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()

	Flags = FlagDate | FlagTime
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+date+"\\s"+time+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()

	Flags = FlagDate | FlagTime | FlagMillis
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+date+"\\s"+time+millis+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()
}
