package logging

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//--------------------------------------------------------------------------------------------------
func TestFormatLogString(t *testing.T) {
	date := time.Date(2018, time.June, 18, 12, 0, 0, 0, time.UTC)
	expected := "[2018-06-18 12:00:00 +0000 UTC][INFO][TestFormatLogString] Test\n"
	actual := formatLogString(date.String(), INFO, "TestFormatLogString", "Test")
	assert.Equal(t, expected, actual)
}

func TestGetCallerContext(t *testing.T) {
	getCallerContext := func() (string, error) {
		t := func() (string, error) {
			return getCallerContext()
		}
		return t()
	}

	expected := "logging_test.go#27"
	actual, err := getCallerContext()
	assert.Nil(t, err)
	assert.Contains(t, actual, expected)
}

//--------------------------------------------------------------------------------------------------
type testingBackend struct {
	LogHistory []string
}

func (backend *testingBackend) log(s string) {
	backend.LogHistory = append(backend.LogHistory, s)
}

func TestInfo(t *testing.T) {
	backend := testingBackend{}
	SetBackend(&backend)

	Info("Testmessage")
	message := backend.LogHistory[0]
	assert.Contains(t, message, "Testmessage")
	assert.Contains(t, message, "[INFO]")
	assert.Contains(t, message, "logging_test.go#45")
}

func TestDebug(t *testing.T) {
	backend := testingBackend{}
	SetBackend(&backend)
	SetLogLevel(DEBUG)

	Debug("Testmessage")
	message := backend.LogHistory[0]
	assert.Contains(t, message, "Testmessage")
	assert.Contains(t, message, "[DEBUG]")
	assert.Contains(t, message, "logging_test.go#57")
}

func TestWarn(t *testing.T) {
	backend := testingBackend{}
	SetBackend(&backend)

	Warn("Testmessage")
	message := backend.LogHistory[0]
	assert.Contains(t, message, "Testmessage")
	assert.Contains(t, message, "[WARN]")
	assert.Contains(t, message, "logging_test.go#68")
}

func TestError(t *testing.T) {
	backend := testingBackend{}
	SetBackend(&backend)

	Error("Testmessage")
	message := backend.LogHistory[0]
	assert.Contains(t, message, "Testmessage")
	assert.Contains(t, message, "[ERROR]")
	assert.Contains(t, message, "logging_test.go#79")
}

func TestFatal(t *testing.T) {
	backend := testingBackend{}
	SetBackend(&backend)

	oldOsExit := exit
	defer func() { exit = oldOsExit }()

	var got int
	myExit := func(code int) {
		got = code
	}

	exit = myExit
	expectedExitCode := 1

	Fatal("Testmessage")
	message := backend.LogHistory[0]
	assert.Contains(t, message, "Testmessage")
	assert.Contains(t, message, "[FATAL]")
	assert.Contains(t, message, "logging_test.go#101")
	assert.Equal(t, expectedExitCode, got)
}
