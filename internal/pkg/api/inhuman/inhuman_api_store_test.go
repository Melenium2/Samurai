package inhuman_test

import (
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/models"
	"context"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// Чтобы нормально пртестировать, нужно мокнуть метод реквест
// либо передовать функцией, либо сделать структуру (что не хочется)
// без этого там нечего тестировать

func TestInhumanApiStore_App_ShouldReturnInstanceOfAppWithoutError(t *testing.T) {
	config := Config()
	config.Hl = "de"
	config.Gl = "de"

	api := inhuman.NewApiStore(config)
	bundle := "1504417378"
	app, err := api.App(bundle)
	assert.NoError(t, err)
	assert.NotNil(t, app)
	assert.Equal(t, bundle, app.Bundle)

	t.Log(app.Screenshots)
}

func TestInhumanApiStore_Flow_ShouldReturnListOfAppsFromMainPage(t *testing.T) {
	config := Config()
	config.Hl = "ru"
	config.Gl = "ru"
	config.ItemsCount = int(math.Min(float64(config.ItemsCount), 200))

	api := inhuman.NewApiStore(config)
	list, err := api.Flow("bank")
	assert.NoError(t, err)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0)
	t.Log(len(list))

	for _, v := range list {
		t.Log(v.Bundle, " ", v.Title)
	}
}

func TestInhumanApiStore_Charts_ShouldReturnListOfAppsByCats(t *testing.T) {
	config := Config()
	config.Hl = "en"
	config.Gl = "uk"

	api := inhuman.NewApiStore(config)
	list, err := api.Charts(context.Background(), models.NewCategory("FINANCE", "topfreeapplications"))
	assert.NoError(t, err)
	assert.NotNil(t, list)
	assert.Greater(t, len(list), 0)
	t.Log(len(list))
	for _, v := range list {
		t.Log(v)
	}
}