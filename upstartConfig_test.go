package upstartConfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	appName := GetName()
	assert.Equal(t, "upstartConfig.test", appName)
}

func TestGetPath(t *testing.T) {
	path := GetPath()
	assert.Contains(t, path, "github.com")
	assert.NotContains(t, path, "upstartConfig.test")
}

func TestGenerate_Defaults(t *testing.T) {
	upstart, _ := Generate()
	assert.Contains(t, upstart, "description \"upstartConfig.test\"")

	assert.Contains(t, upstart, "start on runlevel [2345]")
	assert.Contains(t, upstart, "stop on runlevel [!2345]")
	assert.Contains(t, upstart, "env enabled=1")
}
