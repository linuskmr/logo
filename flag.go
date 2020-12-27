package logo

// Flag with which the logger can be configured.
type Flag uint8

// Different Flag's to configure the equivalent property in the logger.
const (
	DateFlag = 1 << Flag(iota)
	TimeFlag
	MillisFlag
	FilenameFlag
	FuncnameFlag
	JsonFlag
)
