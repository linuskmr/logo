package logo

import (
	"encoding/json"
	"strings"
)

// A Logger entry.
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
	params := []string{}

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
	type Alias entry
	return json.Marshal(&struct {
		Level string `json:"level"`
		*Alias
	}{
		Level: e.Level.String(),
		Alias: (*Alias)(e),
	})
}

// shortFilename cuts filename at the last /.
func shortFilename(filename string) string {
	filePath := strings.Split(filename, "/")
	return filePath[len(filePath)-1]
}
