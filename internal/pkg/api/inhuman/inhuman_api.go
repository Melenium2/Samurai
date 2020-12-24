package inhuman

import (
	"Samurai/internal/pkg/api/models"
)


// Struct of external api instance
type inhumanApi struct {
	config Config
}

// Get application info by bundle from external api
func (api *inhumanApi) App(bundle string) (models.App, error) {
	var app models.App
	err := Request("bundle", "post", WithData(map[string]string{
		"query": bundle,
		"hl": api.config.Hl,
		"gl": api.config.Gl,
	}), WithResponseType(&app), WithApikey(api.config.Key), WithUrl(api.config.Url))

	if err != nil {
		return models.App{}, err
	}

	return app, nil
}

// Method calls api and return top N application from main page
func (api *inhumanApi) Flow(key string) ([]models.App, error) {
	apps := make([]models.App, 0)
	err := Request("mainPage", "post", WithData(map[string]interface{} {
		"query": key,
		"hl": api.config.Hl,
		"gl": api.config.Gl,
		"count": api.config.ItemsCount,
	}), WithResponseType(&apps), WithApikey(api.config.Key), WithUrl(api.config.Url))

	if err != nil {
		return nil, err
	}

	return apps, nil
}

// Create new instance of inhuman api
func NewApiPlay(config Config) *inhumanApi {
	return &inhumanApi{
		config: config,
	}
}

