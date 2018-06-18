package logging

import "os"

// FileBackend a logging backend that logs to file
type FileBackend struct {
	filename string
	fp       *os.File
}

// NewFileBackend Creates a new instance of the FileBackend
func NewFileBackend(filename string) *FileBackend {
	w := &FileBackend{filename: filename}
	var err error
	w.fp, err = os.Create(w.filename)
	if err != nil {
		return nil
	}
	return w
}

func (f FileBackend) log(s string) {
	f.fp.WriteString(s)
}
