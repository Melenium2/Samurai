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

func TestNewProxy_ShouldReturnCorrectInstance(t *testing.T) {
	p := config.NewProxy("http://1d213:asddasd@14.23.51.22:1322")
	assert.Equal(t, "http://1d213:asddasd@14.23.51.22:1322", p.Http)
	assert.Equal(t, "https://1d213:asddasd@14.23.51.22:1322", p.Https)

	p = config.NewProxy("https://1d213:asddasd@14.23.51.22:1322")
	assert.Equal(t, "http://1d213:asddasd@14.23.51.22:1322", p.Http)
	assert.Equal(t, "https://1d213:asddasd@14.23.51.22:1322", p.Https)
}

