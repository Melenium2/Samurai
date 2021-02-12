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

// Interface who provide access to application by models.Category
type ChartApi interface {
	Charts(ctx context.Context, chart models.Category) ([]string, error)
}

// Interface for uploading images to an external server and getting images
// filenames from this server
type ImageProcessingApi interface {
	Process(ctx context.Context, image []string) ([]string, error)
}

// Requester interface provide access to both ExternalApi and ChartApi
type Requester interface {
	ChartApi
	ExternalApi
}
