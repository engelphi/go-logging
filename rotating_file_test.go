package logging

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RotatingFileTestSuite struct {
	suite.Suite
	LogDir       string
	LogFile      string
	LogFileCount int
	LogFileSize  int64
}

func (suite *RotatingFileTestSuite) SetupTest() {
	suite.LogDir = "./logs"
	suite.LogFile = suite.LogDir + "/log"
	suite.LogFileCount = 2
	suite.LogFileSize = int64(1000)
}

func (suite *RotatingFileTestSuite) TearDownTest() {
	os.RemoveAll(suite.LogDir)
}

func (suite *RotatingFileTestSuite) TestNewRotatingFile() {
	os.MkdirAll(suite.LogDir, 0777)

	w, err := NewRotatingFile(suite.LogDir, suite.LogFileCount, suite.LogFileSize)
	defer w.Close()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.LogDir, w.logDirName)
	assert.Equal(suite.T(), suite.LogFileCount, w.fileCount)
	assert.Equal(suite.T(), suite.LogFileSize, w.perFileSizeInByte)
	assert.Equal(suite.T(), int64(0), w.currentFileSizeInByte)
	assert.FileExists(suite.T(), suite.LogFile)
}

func (suite *RotatingFileTestSuite) TestNewRotatingFileOpenIfAlreadyExists() {
	writeBuf := []byte("hallo")
	os.MkdirAll(suite.LogDir, 0777)

	w, err := NewRotatingFile(suite.LogDir, suite.LogFileCount, suite.LogFileSize)
	w.Write(writeBuf)
	w.Close()
	w, err = NewRotatingFile(suite.LogDir, suite.LogFileCount, suite.LogFileSize)
	defer w.Close()
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.LogDir, w.logDirName)
	assert.Equal(suite.T(), suite.LogFileCount, w.fileCount)
	assert.Equal(suite.T(), suite.LogFileSize, w.perFileSizeInByte)
	assert.Equal(suite.T(), int64(len(writeBuf)), w.currentFileSizeInByte)
	assert.FileExists(suite.T(), suite.LogFile)
}

func (suite *RotatingFileTestSuite) TestNewRotatingFileErrorsIfLogFileCountIsLessThanTwo() {
	suite.LogFileCount = 1
	os.MkdirAll(suite.LogDir, 0777)
	w, err := NewRotatingFile(suite.LogDir, suite.LogFileCount, suite.LogFileSize)
	assert.Nil(suite.T(), w)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "fileCount needs to be at least to be rotated", err.Error())
}

func (suite *RotatingFileTestSuite) TestNewRotatingFileErrorsIfLogDirIsNotExisting() {
	w, err := NewRotatingFile(suite.LogDir, suite.LogFileCount, suite.LogFileSize)
	assert.Nil(suite.T(), w)
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "stat ./logs: no such file or directory", err.Error())
}

func (suite *RotatingFileTestSuite) TestRotateLogFiles() {
	rotatedLogFileOne := suite.LogFile + ".1"
	rotatedLogFileTwo := suite.LogFile + ".2"
	suite.LogFileCount = 3
	os.MkdirAll(suite.LogDir, 0777)
	w, err := NewRotatingFile(suite.LogDir, suite.LogFileCount, suite.LogFileSize)
	require.NotNil(suite.T(), w)
	assert.Nil(suite.T(), err)
	defer w.Close()

	err = w.rotateLogFiles()
	assert.Nil(suite.T(), err)
	err = w.rotateLogFiles()
	assert.Nil(suite.T(), err)
	assert.FileExists(suite.T(), suite.LogFile)
	assert.FileExists(suite.T(), rotatedLogFileOne)
	assert.FileExists(suite.T(), rotatedLogFileTwo)
}

func (suite *RotatingFileTestSuite) TestWriteInvokeRotateLogFileIfRemainingLogSizeIsNotSufficient() {
	rotatedLogFileOne := suite.LogFile + ".1"
	suite.LogFileSize = int64(20)
	tenByteMessage := []byte("abcdefghij")
	elevenByteMessage := []byte("abcdefghijk")

	os.MkdirAll(suite.LogDir, 0777)
	w, err := NewRotatingFile(suite.LogDir, suite.LogFileCount, suite.LogFileSize)
	require.NotNil(suite.T(), w)
	assert.Nil(suite.T(), err)
	defer w.Close()
	n, err := w.Write(tenByteMessage)
	assert.Equal(suite.T(), len(tenByteMessage), n)
	assert.Nil(suite.T(), err)

	n, err = w.Write(elevenByteMessage)
	assert.Equal(suite.T(), len(elevenByteMessage), n)
	assert.Nil(suite.T(), err)
	assert.FileExists(suite.T(), rotatedLogFileOne)
	assert.FileExists(suite.T(), suite.LogFile)
}

func TestRotatingFileTestSuite(t *testing.T) {
	suite.Run(t, new(RotatingFileTestSuite))
}
