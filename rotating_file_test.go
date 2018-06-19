package logging

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRotatingFile(t *testing.T) {
	logDir := "./logs"
	logFile := logDir + "/log"
	logFileCount := 2
	logFileSize := int64(1000)
	os.MkdirAll(logDir, 0777)

	w, err := NewRotatingFile(logDir, logFileCount, logFileSize)
	defer w.Close()
	assert.Nil(t, err)
	assert.Equal(t, logDir, w.logDirName)
	assert.Equal(t, logFileCount, w.fileCount)
	assert.Equal(t, logFileSize, w.perFileSizeInByte)
	assert.Equal(t, int64(0), w.currentFileSizeInByte)
	assert.FileExists(t, logFile)

	os.RemoveAll(logDir)
}

func TestNewRotatingFileOpenIfAlreadyExists(t *testing.T) {
	logDir := "./logs"
	logFile := logDir + "/log"
	logFileCount := 2
	logFileSize := int64(1000)
	writeBuf := []byte("hallo")

	os.MkdirAll(logDir, 0777)

	w, err := NewRotatingFile(logDir, logFileCount, logFileSize)
	w.Write(writeBuf)
	w.Close()
	w, err = NewRotatingFile(logDir, logFileCount, logFileSize)

	assert.Nil(t, err)
	assert.Equal(t, logDir, w.logDirName)
	assert.Equal(t, logFileCount, w.fileCount)
	assert.Equal(t, logFileSize, w.perFileSizeInByte)
	assert.Equal(t, int64(len(writeBuf)), w.currentFileSizeInByte)
	assert.FileExists(t, logFile)

	os.RemoveAll(logDir)
}
