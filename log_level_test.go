package logging

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogLevelString(t *testing.T) {
	var invalidLower = DEBUG - 1
	var invalidUpper = FATAL + 1
	assert.Equal(t, "Unknown", fmt.Sprint(invalidUpper))
	assert.Equal(t, "Unknown", fmt.Sprint(invalidLower))
	assert.Equal(t, "DEBUG", fmt.Sprint(DEBUG))
	assert.Equal(t, "INFO", fmt.Sprint(INFO))
	assert.Equal(t, "WARN", fmt.Sprint(WARN))
	assert.Equal(t, "ERROR", fmt.Sprint(ERROR))
	assert.Equal(t, "FATAL", fmt.Sprint(FATAL))
}
