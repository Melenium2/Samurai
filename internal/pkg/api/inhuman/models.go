package inhuman

import (
	"Samurai/internal/pkg/api/models"
	"fmt"
	"strings"
)

// StoreApp represent app bundle from external api
type StoreApp struct {
	ID             string              `json:"id,omitempty"`
	Developer      string              `json:"developer,omitempty"`
	DeveloperId    string              `json:"developerId,omitempty"`
	Title          string              `json:"title,omitempty"`
	Categories     []string            `json:"genres,omitempty"`
	Price          string              `json:"price,omitempty"`
	Image          string              `json:"image,omitempty"`
	Screenshots    map[string][]string `json:"screenshots,omitempty"`
	UserRating     StoreRating         `json:"userRating,omitempty"`
	Description    string              `json:"description,omitempty"`
	ReleaseDate    string              `json:"releaseDate,omitempty"`
	VersionHistory []AppVersionChanges `json:"versionHistory,omitempty"`
	AppSize        int                 `json:"appSize,omitempty"`
	Version        string              `json:"osVersion,omitempty"`
	ContentRating  ContentRating       `json:"contentRating,omitempty"`
	Website        string              `json:"website,omitempty"`
	PrivacyPolicy  string              `json:"privacyPolicy,omitempty"`
}

// Converts StoreApp to models.App
func (sa StoreApp) ToModel() models.App {
	screenshots := make([]models.Screenshots, len(sa.Screenshots))
	i := 0
	for k, v := range sa.Screenshots {
		s := models.Screenshots{
			Device:  k,
			Screens: v,
		}
		screenshots[i] = s
		i++
	}

	histogram := make([]string, len(sa.UserRating.Histogram))
	for i, v := range sa.UserRating.Histogram {
		histogram[i] = fmt.Sprintf("%.0f", v)
	}

	var recentChange string
	var lastUpdateDate string
	var version string
	{
		if len(sa.VersionHistory) > 0 {
			lastChange := sa.VersionHistory[0]
			recentChange = lastChange.ReleaseNotes
			lastUpdateDate = lastChange.Date
			version = lastChange.Version
		}
	}

	return models.App{
		Bundle:          sa.ID,
		DeveloperId:     sa.DeveloperId,
		Developer:       sa.Developer,
		Title:           sa.Title,
		Categories:      strings.Join(sa.Categories, ", "),
		Price:           sa.Price,
		Picture:         sa.Image,
		Screenshots:     screenshots,
		Rating:          fmt.Sprintf("%.1f", sa.UserRating.Value),
		ReviewCount:     fmt.Sprintf("%.0f", sa.UserRating.RatingCount),
		RatingHistogram: histogram,
		Description:     sa.Description,
		RecentChanges:   recentChange,
		ReleaseDate:     sa.ReleaseDate,
		LastUpdateDate:  lastUpdateDate,
		AppSize:         fmt.Sprint(sa.AppSize),
		Version:         version,
		AndroidVersion:  sa.Version,
		ContentRating:   sa.ContentRating.Value,
		DeveloperContacts: models.DeveloperContacts{
			Contacts: sa.Website,
		},
		PrivacyPolicy: sa.PrivacyPolicy,
	}
}

type AppVersionChanges struct {
	Version      string `json:"version,omitempty"`
	ReleaseNotes string `json:"releaseNotes,omitempty"`
	Date         string `json:"date,omitempty"`
}

type StoreRating struct {
	Value       float32   `json:"value,omitempty"`
	RatingCount float32   `json:"ratingCount,omitempty"`
	Histogram   []float32 `json:"histogram,omitempty"`
}

type ContentRating struct {
	Value string `json:"value,omitempty"`
}
