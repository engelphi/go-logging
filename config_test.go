package logging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetBackend(t *testing.T) {
	backend := testingLogWriter{}
	SetBackend(&backend)
	assert.Equal(t, &backend, cfg.backend)
}

func TestLogLevelInitialValue(t *testing.T) {
	assert.Equal(t, INFO, cfg.level)
}

func TestSetLogLevelValidValues(t *testing.T) {
	SetLogLevel(DEBUG)
	assert.Equal(t, DEBUG, cfg.level)

	SetLogLevel(INFO)
	assert.Equal(t, INFO, cfg.level)

	SetLogLevel(WARN)
	assert.Equal(t, WARN, cfg.level)

	SetLogLevel(ERROR)
	assert.Equal(t, ERROR, cfg.level)

	SetLogLevel(FATAL)
	assert.Equal(t, FATAL, cfg.level)
}

func TestSetLogLevelInvalidValues(t *testing.T) {
	invalidLower := DEBUG - 1
	invalidUpper := FATAL + 1
	SetLogLevel(INFO)
	assert.Equal(t, INFO, cfg.level)
	SetLogLevel(invalidLower)
	assert.Equal(t, INFO, cfg.level)
	SetLogLevel(invalidUpper)
	assert.Equal(t, INFO, cfg.level)
}
