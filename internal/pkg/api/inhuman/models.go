package inhuman

import (
	"Samurai/internal/pkg/api/models"
	"encoding/json"
	"fmt"
	"strings"
)

type StoreApp struct {
	ID             string              `json:"id,omitempty"`
	Developer      string              `json:"developer,omitempty"`
	DeveloperId    string              `json:"developer_id,omitempty"`
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

type AppVersionChanges struct {
	Version      string `json:"version,omitempty"`
	ReleaseNotes string `json:"releaseNotes,omitempty"`
	Date         string `json:"date,omitempty"`
}

type StoreRating struct {
	Value       float32 `json:"value,omitempty"`
	RatingCount float32   `json:"ratingCount,omitempty"`
	Histogram   []float32 `json:"histogram,omitempty"`
}

type ContentRating struct {
	Value string `json:"value,omitempty"`
}

func CreateFromStore(m StoreApp) models.App {
	screenshots := make([]string, len(m.Screenshots))
	i := 0
	for k, v := range m.Screenshots {
		s := map[string]interface{}{
			"Device":      k,
			"Screenshots": v,
		}
		b, err := json.Marshal(s)
		if err != nil {
			panic(err)
		}
		screenshots[i] = string(b)
		i++
	}

	histogram := make([]string, len(m.UserRating.Histogram))
	for i, v := range m.UserRating.Histogram {
		histogram[i] = fmt.Sprintf("%.0f", v)
	}

	var recentChange string
	var lastUpdateDate string
	var version string
	{
		if len(m.VersionHistory) > 0 {
			lastChange := m.VersionHistory[0]
			recentChange = lastChange.ReleaseNotes
			lastUpdateDate = lastChange.Date
			version = lastChange.Version
		}
	}

	return models.App{
		Bundle:            m.ID,
		DeveloperId:       m.DeveloperId,
		Developer:         m.Developer,
		Title:             m.Title,
		Categories:        strings.Join(m.Categories, ", "),
		Price:             m.Price,
		Picture:           m.Image,
		Screenshots:       screenshots,
		Rating:            fmt.Sprintf("%.0f", m.UserRating.Value),
		ReviewCount:       fmt.Sprintf("%.0f", m.UserRating.RatingCount),
		RatingHistogram:   histogram,
		Description:       m.Description,
		RecentChanges:     recentChange,
		ReleaseDate:       m.ReleaseDate,
		LastUpdateDate:    lastUpdateDate,
		AppSize:           fmt.Sprint(m.AppSize),
		Version:           version,
		AndroidVersion:    m.Version,
		ContentRating:     m.ContentRating.Value,
		DeveloperContacts: models.DeveloperContacts{
			Contacts: m.Website,
		},
		PrivacyPolicy:     m.PrivacyPolicy,
	}
}

