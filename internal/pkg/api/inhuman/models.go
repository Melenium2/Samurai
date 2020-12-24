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

func (sp StoreApp) ToModel() models.App {
	screenshots := make([]string, len(sp.Screenshots))
	i := 0
	for k, v := range sp.Screenshots {
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

	histogram := make([]string, len(sp.UserRating.Histogram))
	for i, v := range sp.UserRating.Histogram {
		histogram[i] = fmt.Sprintf("%.0f", v)
	}

	var recentChange string
	var lastUpdateDate string
	var version string
	{
		if len(sp.VersionHistory) > 0 {
			lastChange := sp.VersionHistory[0]
			recentChange = lastChange.ReleaseNotes
			lastUpdateDate = lastChange.Date
			version = lastChange.Version
		}
	}

	return models.App{
		Bundle:            sp.ID,
		DeveloperId:       sp.DeveloperId,
		Developer:         sp.Developer,
		Title:             sp.Title,
		Categories:        strings.Join(sp.Categories, ", "),
		Price:             sp.Price,
		Picture:           sp.Image,
		Screenshots:       screenshots,
		Rating:            fmt.Sprintf("%.1f", sp.UserRating.Value),
		ReviewCount:       fmt.Sprintf("%.0f", sp.UserRating.RatingCount),
		RatingHistogram:   histogram,
		Description:       sp.Description,
		RecentChanges:     recentChange,
		ReleaseDate:       sp.ReleaseDate,
		LastUpdateDate:    lastUpdateDate,
		AppSize:           fmt.Sprint(sp.AppSize),
		Version:           version,
		AndroidVersion:    sp.Version,
		ContentRating:     sp.ContentRating.Value,
		DeveloperContacts: models.DeveloperContacts{
			Contacts: sp.Website,
		},
		PrivacyPolicy:     sp.PrivacyPolicy,
	}
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
	return m.ToModel()
}

