package llog

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode     = "INFO"
		dateTime = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d"
		filename = "llog_test.go:18"
		method   = "llog.TestInfo"
	)
	Info(mode)
	assert.Regexp(t, "("+mode+")\\s+"+dateTime+"\\s("+filename+")\\s("+method+": "+mode+")", stringBuilder.String())

}

func TestDebug(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode     = "DEBUG"
		dateTime = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d"
		filename = "llog_test.go:32"
		method   = "llog.TestDebug"
	)
	Debug(mode)
	assert.Regexp(t, "("+mode+")\\s+"+dateTime+"\\s("+filename+")\\s("+method+": "+mode+")", stringBuilder.String())
}

func TestError(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode     = "ERROR"
		dateTime = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d"
		filename = "llog_test.go:45"
		method   = "llog.TestError"
	)
	Error(mode)
	assert.Regexp(t, "("+mode+")\\s+"+dateTime+"\\s("+filename+")\\s("+method+": "+mode+")", stringBuilder.String())
}

func TestPrint(t *testing.T) {
	stringBuilder := strings.Builder{}
	Output = &stringBuilder
	const (
		mode     = "PRINT"
		dateTime = "\\d\\d\\d\\d-\\d\\d-\\d\\d\\s\\d\\d:\\d\\d:\\d\\d"
		filename = "llog_test.go:58"
		method   = "llog.TestPrint"
	)
	Print(mode)
	assert.Regexp(t, "("+mode+")\\s+"+dateTime+"\\s("+filename+")\\s("+method+": "+mode+")", stringBuilder.String())
}
