package inhuman_test

import (
	"Samurai/internal/pkg/api/inhuman"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInhumanApiStore_App_ShouldReturnInstanceOfAppWithoutError(t *testing.T) {
	config := Config()
	api := inhuman.NewApiStore(config)
	bundle := "625334537"
	app, err := api.App(bundle)
	assert.NoError(t, err)
	assert.NotNil(t, app)
	assert.Equal(t, bundle, app.Bundle)
}