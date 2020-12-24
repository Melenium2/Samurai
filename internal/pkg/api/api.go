package api

import (
	"Samurai/internal/pkg/api/models"
	"context"
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

func New(chart ChartApi, common ExternalApi) *ApiImpl {
	return &ApiImpl{
		chart, common,
	}
}

func NewRequester(requester Requester) *ApiImpl {
	chart := requester.(ChartApi)
	common := requester.(ExternalApi)
	return New(chart, common)
}
