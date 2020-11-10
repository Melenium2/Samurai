package config_test

import (
	"Samurai/config"
	"github.com/stretchr/testify/assert"
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

