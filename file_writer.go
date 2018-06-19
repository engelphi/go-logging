package logging

import (
	"io"
	"os"
)

// FileWriter a logging backend that logs to file
type FileWriter struct {
	filename string
	writer   io.Writer
}

// NewFileWriter Creates a new instance of the FileBackend
func NewFileWriter(filename string) *FileWriter {
	w := &FileWriter{filename: filename}
	file, err := os.Create(w.filename)
	if err != nil {
		return nil
	}
	w.writer = file
	return w
}

func (f FileWriter) log(s string) {
	f.writer.Write([]byte(s))
}
