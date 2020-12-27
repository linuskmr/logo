package logo

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	actual := New(FilenameFlag | FuncnameFlag | TimeFlag)
	expected := &Logger{
		Date:       false,
		Time:       true,
		Millis:     false,
		Filename:   true,
		Funcname:   true,
		Json:       false,
		Output:     os.Stdout,
		DateFormat: "2006-01-02",
		TimeFormat: "15:04:05",
		Level:      0,
	}
	assert.Equal(t, expected, actual)
}

func TestLogger_Config(t *testing.T) {
	logger := Logger{
		Date:       false,
		Time:       false,
		Millis:     false,
		Filename:   false,
		Funcname:   false,
		Json:       false,
		Output:     nil,
		DateFormat: "",
		TimeFormat: "",
		Level:      0,
	}
	logger.Config(DateFlag)
	assert.True(t, logger.Date)

	logger.Config(TimeFlag)
	assert.True(t, logger.Time)
	assert.False(t, logger.Date)

	logger.Config(MillisFlag | JsonFlag)
	assert.True(t, logger.Millis)
	assert.True(t, logger.Json)
}
