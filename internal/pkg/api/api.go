package api

import (
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/mobilerpc"
	"context"
)

type Requester interface {
	mobilerpc.ChartApi
	inhuman.ExternalApi
}

type ApiImpl struct {
	mobile  mobilerpc.ChartApi
	inhuman inhuman.ExternalApi
}

func (a *ApiImpl) Charts(ctx context.Context, chart mobilerpc.Category) ([]string, error) {
	return a.mobile.Charts(ctx, chart)
}

func (a *ApiImpl) App(bundle string) (*inhuman.App, error) {
	return a.inhuman.App(bundle)
}

func (a *ApiImpl) Flow(key string) ([]inhuman.App, error) {
	return a.inhuman.Flow(key)
}






