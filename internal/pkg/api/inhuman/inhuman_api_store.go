package inhuman

import (
	"Samurai/internal/pkg/api/models"
	"Samurai/internal/pkg/api/request"
	"context"
	"fmt"
)

type inhumanApiStore struct {
	config Config
}

// App method requests to external api and gets models.App by given string
func (ias *inhumanApiStore) App(bundle string) (models.App, error) {
	var app StoreApp
	err := request.Request("ios_bundle", "GET", request.WithQueryParams(map[string]interface{}{
		"q": bundle,
		"l": ias.config.Gl,
	}), request.WithApikey(ias.config.Key), request.WithUrl(ias.config.Url), request.WithResponseType(&app))
	if err != nil {
		return models.App{}, err
	}

	return app.ToModel(), nil
}

// Flow method request to external api and return []models.App by given term
func (ias *inhumanApiStore) Flow(key string) ([]models.App, error) {
	var list []StoreApp
	err := request.Request("ios_apps_list", "GET", request.WithQueryParams(map[string]interface{}{
		"q":    key,
		"o":    ias.config.ItemsCount,
		"lang": ias.config.Hl,
		"geo":  ias.config.Gl,
	}), request.WithApikey(ias.config.Key), request.WithUrl(ias.config.Url), request.WithResponseType(&list))
	if err != nil {
		return nil, err
	}

	apps := make([]models.App, len(list))
	for i, v := range list {
		apps[i] = v.ToModel()
	}

	return apps, nil
}

// Charts method request to external api and return []string which contains
// bundleid of app. []string contains bundle from given models.Category
func (ias *inhumanApiStore) Charts(ctx context.Context, chart models.Category) ([]string, error) {
	cat, subcat := chart.Split()
	var list []map[string]interface{}
	err := request.Request("ios/collections", "POST", request.WithData(map[string]interface{}{
		"cat":     cat,
		"subCat":  subcat,
		"count":   ias.config.ItemsCount,
		"country": ias.config.Gl,
	}), request.WithUrl(ias.config.Url), request.WithApikey(ias.config.Key), request.WithResponseType(&list))
	if err != nil {
		return nil, err
	}

	bundles := make([]string, len(list))
	for i, v := range list {
		bundle, ok := v["id"].(string)
		if !ok {
			return nil, fmt.Errorf("can not cast %v with field id to string", v)
		}
		bundles[i] = bundle
	}

	return bundles, nil
}

func NewApiStore(config Config) *inhumanApiStore {
	return &inhumanApiStore{
		config: config,
	}
}
