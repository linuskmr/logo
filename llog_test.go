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
	fmt.Println(stringBuilder.String())
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
	fmt.Println(stringBuilder.String())
	assert.Regexp(t, "("+mode+")\\s+"+dateTimeMillis+"\\s("+filename+"):\\d+\\s("+method+": "+mode+")", stringBuilder.String())
}

func TestWarn(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode           = "WARN"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "llog_test.go"
		method         = "llog.TestWarn"
	)
	Warn(mode)
	fmt.Println(stringBuilder.String())
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
	fmt.Println(stringBuilder.String())
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
	fmt.Println(stringBuilder.String())
	assert.Regexp(t, "("+mode+")\\s+"+dateTimeMillis+"\\s("+filename+"):\\d+\\s("+method+": "+mode+")", stringBuilder.String())
}

func TestConfig(t *testing.T) {
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
	Config(FlagDate)
	Print(mode)
	fmt.Println(stringBuilder.String())
	assert.Regexp(t, "("+mode+")\\s+"+date+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()

	Config(FlagTime)
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+time+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()

	Config(FlagMillis)
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+millis+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()

	Config(FlagDate | FlagTime)
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+date+"\\s"+time+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()

	Config(FlagDate | FlagTime | FlagMillis)
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+date+"\\s"+time+millis+":\\s("+mode+")", stringBuilder.String())
	stringBuilder.Reset()
}
