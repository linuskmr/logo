package logo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	stdLog "log"
	"strconv"
	"strings"
	"testing"
)

func TestInfo(t *testing.T) {
	var stringBuilder strings.Builder
	Standard.Output = &stringBuilder
	const (
		mode           = ".*INFO.*"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "logo.TestInfo"
	)
	Info(mode)
	assert.Regexp(t, mode+` `+dateTimeMillis+` `+filename+`:\d+ `+method+`: `+mode, stringBuilder.String())
}

func TestDebug(t *testing.T) {
	var stringBuilder strings.Builder
	Standard.Output = &stringBuilder
	const (
		mode           = ".*DEBUG.*"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "logo.TestDebug"
	)
	Debug(mode)
	assert.Regexp(t, mode+` +`+dateTimeMillis+` `+filename+`:\d+ `+method+`: `+mode, stringBuilder.String())
}

func TestWarn(t *testing.T) {
	var stringBuilder strings.Builder
	Standard.Output = &stringBuilder
	const (
		mode           = ".*WARN.*"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "logo.TestWarn"
	)
	Warn(mode)
	assert.Regexp(t, mode+` +`+dateTimeMillis+` `+filename+`:\d+ `+method+`: `+mode, stringBuilder.String())
}

func TestError(t *testing.T) {
	var stringBuilder strings.Builder
	Standard.Output = &stringBuilder
	const (
		mode           = ".*ERROR.*"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "logo.TestError"
	)
	Error(mode)
	assert.Regexp(t, mode+` +`+dateTimeMillis+` `+filename+`:\d+ `+method+`: `+mode, stringBuilder.String())
}

func TestPrint(t *testing.T) {
	var stringBuilder strings.Builder
	Standard.Output = &stringBuilder
	const (
		mode           = "PRINT"
		dateTimeMillis = "\\d\\d\\d\\d-\\d\\d-\\d\\d \\d\\d:\\d\\d:\\d\\d.\\d\\d\\d"
		filename       = "logo_test.go"
		method         = "logo.TestPrint"
	)
	Print(mode)
	assert.Regexp(t, mode+` +`+dateTimeMillis+` `+filename+`:\d+ `+method+`: `+mode, stringBuilder.String())
}

func TestConfig(t *testing.T) {
	var stringBuilder strings.Builder
	Standard.Output = &stringBuilder
	const (
		mode   = "PRINT"
		date   = "\\d\\d\\d\\d-\\d\\d-\\d\\d"
		time   = "\\d\\d:\\d\\d:\\d\\d"
		millis = ".\\d\\d\\d"
	)
	Standard.Config(DateFlag)
	Print(mode)
	assert.Regexp(t, mode+" +"+date+": "+mode, stringBuilder.String())
	stringBuilder.Reset()

	Standard.Config(TimeFlag)
	Print(mode)
	assert.Regexp(t, mode+" +"+time+": "+mode, stringBuilder.String())
	stringBuilder.Reset()

	Standard.Config(MillisFlag)
	Print(mode)
	assert.Regexp(t, mode+" +"+millis+": "+mode, stringBuilder.String())
	stringBuilder.Reset()

	Standard.Config(DateFlag | TimeFlag)
	Print(mode)
	assert.Regexp(t, mode+" +"+date+" "+time+": "+mode, stringBuilder.String())
	stringBuilder.Reset()

	Standard.Config(DateFlag | TimeFlag | MillisFlag)
	Print(mode)
	assert.Regexp(t, mode+" +"+date+" "+time+millis+": "+mode, stringBuilder.String())
	stringBuilder.Reset()
}

func TestJson(t *testing.T) {
	var stringBuilder strings.Builder
	Standard.Output = &stringBuilder
	Standard.Config(1<<6 - 1)
	Print("Print")
	var expected = []string{
		`\s+"level": "PRINT"`,
		`\s+"date": "\d\d\d\d-\d\d-\d\d"`,
		`\s+"time": "\d\d:\d\d:\d\d.\d\d\d"`,
		`\s+"msg": "Print"`,
		`\s+"filename": "logo_test.go:\d+"`,
		`\s+"func_name": "logo.TestJson"`,
	}

	assert.Regexp(t, `{\n`+strings.Join(expected, `,\n?`)+`\n\s+}`, stringBuilder.String())
}

func BenchmarkInfo(b *testing.B) {
	Standard.Output = ioutil.Discard
	for i := 0; i < b.N; i++ {
		Info(i)
	}
}

func BenchmarkStdLog(b *testing.B) {
	stdLog.SetOutput(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		stdLog.Print(i)
	}
}

func BenchmarkFmtPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Fprintln(ioutil.Discard, i)
	}
}

func BenchmarkWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ioutil.Discard.Write([]byte(strconv.Itoa(i)))
	}
}
