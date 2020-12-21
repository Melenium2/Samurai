package inhuman

import (
	"Samurai/internal/pkg/api/models"
	"context"
	"fmt"
)

type inhumanApiStore struct {
	config Config
}

func (i *inhumanApiStore) App(bundle string) (models.App, error) {
	var app StoreApp
	err := Request("ios_bundle", "GET", WithQueryParams(map[string]interface{} {
		"q": bundle,
		"l": i.config.Gl,
	}), WithApikey(i.config.Key), WithUrl(i.config.Url), WithResponseType(&app))
	if err != nil {
		return models.App{}, err
	}

	return CreateFromStore(app), nil
}

func (i *inhumanApiStore) Flow(key string) ([]models.App, error) {
	var list []StoreApp
	err := Request("ios_apps_list", "GET", WithQueryParams(map[string]interface{} {
		"q": key,
		"lang": i.config.Hl,
		"geo": i.config.Gl,
	}), WithApikey(i.config.Key), WithUrl(i.config.Url), WithResponseType(&list))
	if err != nil {
		return nil, err
	}

	apps := make([]models.App, len(list))
	for i, v := range list {
		apps[i] = CreateFromStore(v)
	}

	return apps, nil
}

func (i *inhumanApiStore) Charts(ctx context.Context, chart models.Category) ([]string, error) {
	cat, subcat := chart.Split()
	var list []map[string]interface{}
	err := Request("ios/collections", "POST", WithData(map[string]interface{} {
		"cat": cat,
		"subcat": subcat,
		"count": 200,
		"country": i.config.Gl,
	}), WithUrl(i.config.Url), WithApikey(i.config.Key), WithResponseType(&list))
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



