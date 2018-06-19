package logging

type config struct {
	backend LogWriter
	level   LogLevel
}

// SetBackend Sets the backend that should be used by the logger
func SetBackend(backend LogWriter) {
	cfg.backend = backend
}

// SetLogLevel Sets the log level that controls which messages are logged
func SetLogLevel(level LogLevel) {
	if level < DEBUG || level > FATAL {
		return
	}
	cfg.level = level
}
