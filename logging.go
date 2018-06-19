package logging

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

//--------------------------------------------------------------------------------------------------
var exit = os.Exit
var cfg config

func init() {
	cfg.backend = ConsoleWriter{}
	cfg.level = INFO
}

//--------------------------------------------------------------------------------------------------

func formatLogString(time string, level LogLevel, function, msg string) string {
	return fmt.Sprintf("[%s][%s][%s] %s\n", time, level, function, msg)
}

func getCallerContext() (string, error) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		return "", errors.New("unable to get calling context")
	}

	return fmt.Sprintf("%s#%d", file, line), nil
}

func logBase(level LogLevel, msg string) {
	if level < cfg.level {
		return
	}

	t := time.Now()
	caller, err := getCallerContext()
	if err != nil {
		log.Fatalf("%v", err)
	}
	cfg.backend.write(formatLogString(t.Local().String(), level, caller, msg))
}

//--------------------------------------------------------------------------------------------------

// Debug logs debug messages
func Debug(msg string, v ...interface{}) {
	logBase(DEBUG, fmt.Sprintf(msg, v...))
}

// Info logs info messages
func Info(msg string, v ...interface{}) {
	logBase(INFO, fmt.Sprintf(msg, v...))
}

// Warn logs warn messages
func Warn(msg string, v ...interface{}) {
	logBase(WARN, fmt.Sprintf(msg, v...))
}

// Error logs error messages
func Error(msg string, v ...interface{}) {
	logBase(ERROR, fmt.Sprintf(msg, v...))
}

// Fatal logs error messages and exits application
func Fatal(msg string, v ...interface{}) {
	logBase(FATAL, fmt.Sprintf(msg, v...))
	exit(1)
}
