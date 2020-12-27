package logo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEntry_String(t *testing.T) {
	e := entry{
		Level:    DebugLevel,
		Date:     "",
		Time:     "",
		Msg:      "",
		Filename: "",
		FuncName: "",
	}
	assert.Regexp(t, ".*DEBUG.*: .*", e.String())
}
