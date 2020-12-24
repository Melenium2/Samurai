package inhuman_test

import (
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/models"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Чтобы нормально пртестировать, нужно мокнуть метод реквест
// либо передовать функцией, либо сделать структуру (что не хочется)
// без этого там нечего тестировать

func TestInhumanApiStore_App_ShouldReturnInstanceOfAppWithoutError(t *testing.T) {
	config := Config()
	api := inhuman.NewApiStore(config)
	bundle := "625334537"
	app, err := api.App(bundle)
	assert.NoError(t, err)
	assert.NotNil(t, app)
	assert.Equal(t, bundle, app.Bundle)
}

func TestInhumanApiStore_Flow_ShouldReturnListOfAppsFromMainPage(t *testing.T) {
	config := Config()
	api := inhuman.NewApiStore(config)
	list, err := api.Flow("car")
	assert.NoError(t, err)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0)
	t.Log(len(list))
}

func TestInhumanApiStore_Charts_ShouldReturnListOfAppsByCats(t *testing.T) {
	config := Config()
	api := inhuman.NewApiStore(config)
	list, err := api.Charts(context.Background(), models.NewCategory("games", "topfreeapplications"))
	assert.NoError(t, err)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0)
	t.Log(len(list))
	for _, v := range list {
		t.Log(v)
	}
}