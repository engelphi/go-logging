package logging

// LogLevel Type that represents the different log levels
type LogLevel int

// The different possible log levels
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (level LogLevel) String() string {
	levels := [...]string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
	if level < DEBUG || level > FATAL {
		return "Unknown"
	}

	return levels[level]
}
