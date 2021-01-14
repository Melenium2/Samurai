package inhuman

import (
	"Samurai/internal/pkg/api/models"
	"Samurai/internal/pkg/api/request"
)

// Struct of external api instance
type inhumanApi struct {
	config Config
}

// Get application info by bundle from external api
func (api *inhumanApi) App(bundle string) (models.App, error) {
	var app models.GoogleApp
	err := request.Request("bundle", "post", request.WithData(map[string]string{
		"query": bundle,
		"hl":    api.config.Hl,
		"gl":    api.config.Gl,
	}), request.WithResponseType(&app), request.WithApikey(api.config.Key), request.WithUrl(api.config.Url))

	if err != nil {
		return models.App{}, err
	}

	return app.ToModel(), nil
}

// Method calls api and return top N application from main page
func (api *inhumanApi) Flow(key string) ([]models.App, error) {
	var list []models.GoogleApp
	err := request.Request("mainPage", "post", request.WithData(map[string]interface{}{
		"query": key,
		"hl":    api.config.Hl,
		"gl":    api.config.Gl,
		"count": api.config.ItemsCount,
	}), request.WithResponseType(&list), request.WithApikey(api.config.Key), request.WithUrl(api.config.Url))

	if err != nil {
		return nil, err
	}

	apps := make([]models.App, len(list))
	for i, v := range list {
		apps[i] = v.ToModel()
	}

	return apps, nil
}

// Create new instance of inhuman api
func NewApiPlay(config Config) *inhumanApi {
	return &inhumanApi{
		config: config,
	}
}
