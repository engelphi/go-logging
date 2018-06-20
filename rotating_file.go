package logging

import (
	"errors"
	"os"
	"strconv"
)

// RotatingFile an io.Writer that creates rotating files
type RotatingFile struct {
	logDirName            string
	currentFile           *os.File
	currentFileSizeInByte int64
	fileCount             int
	perFileSizeInByte     int64
}

const logBaseName = "log"

func (fw *RotatingFile) rotateLogFiles() error {
	fw.currentFile.Close()

	logFilePath := fw.logDirName + "/" + logBaseName

	for fileIdx := fw.fileCount - 1; fileIdx > 1; fileIdx-- {
		oldName := logFilePath + "." + strconv.Itoa(fileIdx-1)
		newName := logFilePath + "." + strconv.Itoa(fileIdx)
		err := os.Rename(oldName, newName)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		}
	}

	newName := logFilePath + "." + strconv.Itoa(1)
	err := os.Rename(logFilePath, newName)
	if err != nil {
		return err
	}

	fw.currentFile, err = os.Create(logFilePath)
	fw.currentFileSizeInByte = 0
	return err
}

// NewRotatingFile creates a io.Writer that writes rotating log files
func NewRotatingFile(logDirName string, fileCount int, perFileSizeInByte int64) (*RotatingFile, error) {
	if stat, err := os.Stat(logDirName); os.IsNotExist(err) || !stat.IsDir() {
		return nil, err
	}

	if fileCount < 2 {
		return nil, errors.New("fileCount needs to be at least to be rotated")
	}

	w := &RotatingFile{logDirName: logDirName, fileCount: fileCount, perFileSizeInByte: perFileSizeInByte}

	logFilePath := logDirName + "/" + logBaseName

	if stat, err := os.Stat(logFilePath); os.IsNotExist(err) {
		w.currentFile, err = os.Create(logFilePath)
		if err != nil {
			return nil, err
		}
		w.currentFileSizeInByte = 0
	} else {
		w.currentFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, stat.Mode())
		if err != nil {
			return nil, err
		}
		w.currentFileSizeInByte = stat.Size()
	}

	return w, nil
}

// Write implements the io.Writer interface for RotatingFileWriter
func (fw *RotatingFile) Write(p []byte) (int, error) {
	bufLength := len(p)
	if (fw.currentFileSizeInByte + int64(bufLength)) > fw.perFileSizeInByte {
		err := fw.rotateLogFiles()
		if err != nil {
			return 0, err
		}
	}

	n, err := fw.currentFile.Write(p)
	if err != nil {
		return n, err
	}

	fw.currentFileSizeInByte += int64(n)
	if n != bufLength {
		return n, errors.New("unable to write complete message")
	}

	return bufLength, nil
}

// Close Implements the Closer interface for RotatingFile
func (fw RotatingFile) Close() error {
	return fw.currentFile.Close()
}
