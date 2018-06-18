package logging

import (
	"strings"
	"testing"
	"time"
)

//--------------------------------------------------------------------------------------------------
func TestFormatLogString(t *testing.T) {
	date := time.Date(2018, time.June, 18, 12, 0, 0, 0, time.UTC)
	expected := "[2018-06-18 12:00:00 +0000 UTC][INFO][TestFormatLogString] Test\n"
	actual := formatLogString(date.String(), INFO, "TestFormatLogString", "Test")
	if expected != actual {
		t.Errorf("Unexepected log string:\nExpected: %s\nActual: %s\n", expected, actual)
	}
}

func TestGetCallerContext(t *testing.T) {
	getCallerContext := func() (string, error) {
		t := func() (string, error) {
			return getCallerContext()
		}
		return t()
	}

	expected := "logging_test.go#28"
	actual, err := getCallerContext()
	if err != nil {
		t.Errorf("getCallerContext failed: %v\n", err)
	}

	if !strings.Contains(actual, expected) {
		t.Errorf("getCallerContext failed to retrieve correct caller:\nExpected: %s, Actual: %s\n", expected, actual)
	}
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
	if !(strings.Contains(message, "Testmessage") && strings.Contains(message, "[INFO]") && strings.Contains(message, "logging_test.go#51")) {
		t.Errorf("Did not found expected log message")
	}
}

func TestDebug(t *testing.T) {
	backend := testingBackend{}
	SetBackend(&backend)
	SetLogLevel(DEBUG)

	Debug("Testmessage")
	message := backend.LogHistory[0]
	if !(strings.Contains(message, "Testmessage") && strings.Contains(message, "[DEBUG]") && strings.Contains(message, "logging_test.go#63")) {
		t.Errorf("Did not found expected log message")
	}
}

func TestWarn(t *testing.T) {
	backend := testingBackend{}
	SetBackend(&backend)

	Warn("Testmessage")
	message := backend.LogHistory[0]
	if !(strings.Contains(message, "Testmessage") && strings.Contains(message, "[WARN]") && strings.Contains(message, "logging_test.go#74")) {
		t.Errorf("Did not found expected log message")
	}
}

func TestError(t *testing.T) {
	backend := testingBackend{}
	SetBackend(&backend)

	Error("Testmessage")
	message := backend.LogHistory[0]
	if !(strings.Contains(message, "Testmessage") && strings.Contains(message, "[ERROR]") && strings.Contains(message, "logging_test.go#85")) {
		t.Errorf("Did not found expected log message")
	}
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
	if !(strings.Contains(message, "Testmessage") && strings.Contains(message, "[FATAL]") && strings.Contains(message, "logging_test.go#107")) {
		t.Errorf("Did not found expected log message")
	}

	if got != expectedExitCode {
		t.Errorf("Unexpected error code")
	}
}
