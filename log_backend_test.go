package logging

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWriter struct {
	mock.Mock
}

func (m *MockWriter) Write(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func TestNewLogWriterBackend(t *testing.T) {
	testWriter := new(MockWriter)
	logBackend := NewLogWriterBackend(testWriter)
	assert.Equal(t, testWriter, logBackend.backend)
}

func TestNewConsoleLogWriter(t *testing.T) {
	logBackend := NewConsoleLogWriter()
	assert.Equal(t, os.Stdout, logBackend.backend)
}

func TestLogWriterBackendWrite(t *testing.T) {
	testWriter := new(MockWriter)
	logBackend := NewLogWriterBackend(testWriter)
	msg := "Test"
	testMessage := []byte(msg)
	testWriter.On("Write", testMessage).Return(len(testMessage), nil)
	assert.Nil(t, logBackend.write(msg))
	testWriter.AssertExpectations(t)
}

func TestLogWriterBackendWriteErrorOccured(t *testing.T) {
	testWriter := new(MockWriter)
	logBackend := NewLogWriterBackend(testWriter)
	msg := "Test"
	testMessage := []byte(msg)
	err := errors.New("test")
	testWriter.On("Write", testMessage).Return(len(testMessage), err)
	assert.Equal(t, err, logBackend.write(msg))
	testWriter.AssertExpectations(t)
}

func TestLogWriterBackendWriteWrittenLessBytesThanExpected(t *testing.T) {
	testWriter := new(MockWriter)
	logBackend := NewLogWriterBackend(testWriter)
	msg := "Test"
	testMessage := []byte(msg)
	testWriter.On("Write", testMessage).Return(len(testMessage)-1, nil)

	err := logBackend.write(msg)
	assert.NotNil(t, err)
	assert.Equal(t, "unable to write complete msg to log", err.Error())
	testWriter.AssertExpectations(t)
}
