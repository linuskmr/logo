package logo

import (
	"encoding/json"
	"strings"
)

// entry is a logging message.
type entry struct {
	Level    Level  `json:"level"`
	Date     string `json:"date,omitempty"`
	Time     string `json:"time,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Filename string `json:"filename,omitempty"`
	FuncName string `json:"func_name,omitempty"`
}

// Converts an entry to a string, with the Level colored.
func (e *entry) String() string {
	var params []string
	params = append(params, e.Level.ColorizedString())
	if e.Date != "" {
		params = append(params, e.Date)
	}
	if e.Time != "" {
		params = append(params, e.Time)
	}
	if e.Filename != "" {
		params = append(params, e.Filename)
	}
	if e.FuncName != "" {
		params = append(params, e.FuncName)
	}
	return strings.Join(params, " ") + ": " + e.Msg
}

func (e *entry) MarshalJSON() ([]byte, error) {
	// Overwrite level with a string representation of level
	type alias entry
	return json.Marshal(&struct {
		// The new overwritten field
		Level string `json:"level"`
		// Fill with the remaining fields from alias aka entry
		*alias
	}{
		Level: e.Level.String(),
		alias: (*alias)(e),
	})
}

// shortFilename cuts filename at the last /.
func shortFilename(filename string) string {
	filePath := strings.Split(filename, "/")
	return filePath[len(filePath)-1]
}
