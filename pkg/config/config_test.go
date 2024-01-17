package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfigValues(t *testing.T) {

	assert.Equal(t, Config.Port, "9101")
	assert.Equal(t, Config.LabelProject, "project")
}
