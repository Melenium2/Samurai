package api

import (
	"Samurai/config"
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/mobilerpc"
	"context"
	"strings"
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

func New(config config.ApiConfig, lang string) *ApiImpl {
	hlgl := strings.Split(lang, "_")
	rpc := mobilerpc.New(mobilerpc.Config{
		Address:    config.GrpcAddress,
		Port:       config.GrpcPort,
		RpcAccount: config.GrpcAccount,
	})
	api := inhuman.New(inhuman.Config{
		Url:       config.Url,
		Key:       config.Key,
		Hl:        strings.ToLower(hlgl[0]),
		Gl:        strings.ToLower(hlgl[1]),
		AppsCount: 250,
	})
	return &ApiImpl{
		rpc, api,
	}
}
