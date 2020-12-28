package api

import (
	"Samurai/internal/pkg/api/models"
	"context"
)

// Interface who provide api to external api
type ExternalApi interface {
	App(bundle string) (models.App, error)
	Flow(key string) ([]models.App, error)
}

type ChartApi interface {
	Charts(ctx context.Context, chart models.Category) ([]string, error)
}

type Requester interface {
	ChartApi
	ExternalApi
}
