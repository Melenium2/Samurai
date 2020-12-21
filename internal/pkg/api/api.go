package api

import (
	"Samurai/config"
	"Samurai/internal/pkg/api/inhuman"
	"Samurai/internal/pkg/api/mobilerpc"
	"Samurai/internal/pkg/api/models"
	"context"
	"strings"
)


type ApiImpl struct {
	chart  ChartApi
	common ExternalApi
}

func (a *ApiImpl) Charts(ctx context.Context, chart models.Category) ([]string, error) {
	return a.chart.Charts(ctx, chart)
}

func (a *ApiImpl) App(bundle string) (models.App, error) {
	return a.common.App(bundle)
}

func (a *ApiImpl) Flow(key string) ([]models.App, error) {
	return a.common.Flow(key)
}

func New(config config.ApiConfig, lang string) *ApiImpl {
	hlgl := strings.Split(lang, "_")
	rpc := mobilerpc.New(mobilerpc.Config{
		Address:    config.GrpcAddress,
		Port:       config.GrpcPort,
		RpcAccount: config.GrpcAccount,
	})
	api := inhuman.NewApiPlay(inhuman.Config{
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
