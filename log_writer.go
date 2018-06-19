package logging

// LogWriter Interface that logging backends need to fulfill
type LogWriter interface {
	write(s string)
}

type testingLogWriter struct {
	LogHistory []string
}

func (backend *testingLogWriter) write(s string) {
	backend.LogHistory = append(backend.LogHistory, s)
}
