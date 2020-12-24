package inhuman

import (
	"Samurai/internal/pkg/api/models"
	"context"
	"fmt"
)

type inhumanApiStore struct {
	config Config
}

func (ias *inhumanApiStore) App(bundle string) (models.App, error) {
	var app StoreApp
	err := Request("ios_bundle", "GET", WithQueryParams(map[string]interface{}{
		"q": bundle,
		"l": ias.config.Gl,
	}), WithApikey(ias.config.Key), WithUrl(ias.config.Url), WithResponseType(&app))
	if err != nil {
		return models.App{}, err
	}

	return app.ToModel(), nil
}

func (ias *inhumanApiStore) Flow(key string) ([]models.App, error) {
	var list []StoreApp
	err := Request("ios_apps_list", "GET", WithQueryParams(map[string]interface{}{
		"q":    key,
		"o":    ias.config.ItemsCount,
		"lang": ias.config.Hl,
		"geo":  ias.config.Gl,
	}), WithApikey(ias.config.Key), WithUrl(ias.config.Url), WithResponseType(&list))
	if err != nil {
		return nil, err
	}

	apps := make([]models.App, len(list))
	for i, v := range list {
		apps[i] = v.ToModel()
	}

	return apps, nil
}

func (ias *inhumanApiStore) Charts(ctx context.Context, chart models.Category) ([]string, error) {
	cat, subcat := chart.Split()
	var list []map[string]interface{}
	err := Request("ios/collections", "POST", WithData(map[string]interface{}{
		"cat":     cat,
		"subCat":  subcat,
		"count":   ias.config.ItemsCount,
		"country": ias.config.Gl,
	}), WithUrl(ias.config.Url), WithApikey(ias.config.Key), WithResponseType(&list))
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
