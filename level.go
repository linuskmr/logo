package logo

// Level is a log level of a Logger.
type Level uint8

// The log Level's following are sorted in ascending order of priority.
// DebugLevel(AllLevels), InfoLevel, WarnLevel, ErrorLevel and PrintLevel. So
// the highest Level is PrintLevel, the lowest is DebugLevel(PrintLevel). A
// newEntry is only logged if its level is greater or equal to the Level of the
// Logger.
const (
	DebugLevel = Level(iota)
	InfoLevel
	WarnLevel
	ErrorLevel
	PrintLevel
	AllLevels = DebugLevel
)

// levelNames stores the name and the colorizedName for a log Level.
type levelNames struct {
	name          string
	colorizedName string
}

// String returns the name of the log Level.
func (l Level) String() string {
	return levels[l].name
}

// ColorizedString returns the colorized name of the log Level.
func (l Level) ColorizedString() string {
	return levels[l].colorizedName
}

// levels contains text and colorizedName for each log Level.
var levels = [...]levelNames{
	DebugLevel: {
		name:          "DEBUG",
		colorizedName: "\033[32mDEBUG\033[0m",
	},
	InfoLevel: {
		name:          "INFO",
		colorizedName: "\033[34mINFO\033[0m",
	},
	WarnLevel: {
		name:          "WARN",
		colorizedName: "\033[33mWARN\033[0m",
	},
	ErrorLevel: {
		name:          "ERROR",
		colorizedName: "\033[31mERROR\033[0m",
	},
	PrintLevel: {
		name:          "PRINT",
		colorizedName: "PRINT",
	},
}
