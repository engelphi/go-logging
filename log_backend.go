package logging

import (
	"errors"
	"io"
	"os"
)

// LogBackend Interface that logging backends need to fulfill
type LogBackend interface {
	write(s string) error
}

type testingLogBackend struct {
	LogHistory []string
}

func (backend *testingLogBackend) write(s string) error {
	backend.LogHistory = append(backend.LogHistory, s)
	return nil
}

// LogWriterBackend generic write which uses the io.Writer interface
type LogWriterBackend struct {
	backend io.Writer
}

func (logger LogWriterBackend) write(msg string) error {
	rawMsg := []byte(msg)
	n, err := logger.backend.Write(rawMsg)
	if err != nil {
		return err
	}

	if n != len(rawMsg) {
		return errors.New("unable to write complete msg to log")
	}

	return nil
}

// NewLogWriterBackend creates a new backend which uses the given io.Writer
func NewLogWriterBackend(writer io.Writer) *LogWriterBackend {
	w := &LogWriterBackend{backend: writer}
	return w
}

// NewConsoleLogWriter creates a new logger backend which writes to the console
func NewConsoleLogWriter() *LogWriterBackend {
	return NewLogWriterBackend(os.Stdout)
}
