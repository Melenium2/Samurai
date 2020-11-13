package config_test

import (
	"Samurai/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig_ShouldCreateValidInstance_NoError(t *testing.T) {
	c := config.New()
	assert.NotEmpty(t, c.Api.Url)
	assert.NotEmpty(t, c.Api.Key)
	assert.NotEmpty(t, c.Database.Database)
	assert.NotEmpty(t, c.Database.User)
	assert.NotEmpty(t, c.Database.Password)
	assert.NotEmpty(t, c.Database.Schema)
	assert.NotEmpty(t, c.Database.Address)
	assert.NotEmpty(t, c.Database.Port)
}

func TestConfig_ShouldOverwriteAlreadyExistingParams(t *testing.T) {
	c := config.New()
	assert.NotEmpty(t, c.Api.Key)
	os.Setenv("api_key", "hello")
	c = config.New()
	assert.Equal(t, "hello", c.Api.Key)
}

