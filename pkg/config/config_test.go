package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfigValues(t *testing.T) {
	assert.Equal(t, Config.Database, "127.0.0.1")
	assert.Equal(t, Config.DatabaseName, "eventstracker")
	assert.Equal(t, Config.DatabasePort, "27017")
	assert.Equal(t, Config.LogLevel, "info")
	assert.Equal(t, Config.Port, "9101")
}
