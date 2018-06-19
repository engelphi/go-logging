package logging

import (
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

func TestNewFileWriter(t *testing.T) {
	filename := "Test.log"
	fileWriter := NewFileWriter(filename)
	assert.Equal(t, filename, fileWriter.filename)
	assert.FileExists(t, filename)
}

func TestLog(t *testing.T) {
	filename := "Test.log"
	fileWriter := NewFileWriter(filename)
	assert.Equal(t, filename, fileWriter.filename)
	assert.FileExists(t, filename)

	testWriter := new(MockWriter)
	testMessage := "test"
	testWriter.On("Write", []byte(testMessage)).Return(5, nil)
	fileWriter.writer = testWriter
	fileWriter.log(testMessage)
	testWriter.AssertExpectations(t)
}
