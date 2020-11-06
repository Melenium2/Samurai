package db

import (
	"time"
)

type App struct {
	Bundle      string    `json:"bundle,omitempty"`
	Category    string    `json:"category,omitempty"`
	DeveloperId string    `json:"developer_id,omitempty"`
	Developer   string    `json:"developer,omitempty"`
	Geo         string    `json:"geo,omitempty"`
	StartAt     time.Time `json:"start_at,omitempty"`
	Period      uint32    `json:"period,omitempty"`
}

type Meta struct {
	Bundle            string            `json:"bundle" db:"bundle"`
	Title             string            `json:"title" db:"title"`
	Price             string            `json:"price" db:"price"`
	Picture           string            `json:"picture" db:"picture"`
	Screenshots       []string          `json:"screenshots" db:"screenshots"`
	Rating            string            `json:"rating" db:"rating"`
	ReviewCount       string            `json:"reviewCount" db:"review_count"`
	RatingHistogram   []string          `json:"ratingHistogram" db:"rating_histogram"`
	Description       string            `json:"description" db:"description"`
	ShortDescription  string            `json:"shortDescription" db:"short_description"`
	RecentChanges     string            `json:"recentChanges" db:"recent_changes"`
	ReleaseDate       string            `json:"releaseDate" db:"release_date"`
	LastUpdateDate    string            `json:"lastUpdateDate" db:"last_update_date"`
	AppSize           string            `json:"appSize" db:"app_size"`
	Installs          string            `json:"installs" db:"installs"`
	Version           string            `json:"version" db:"version"`
	AndroidVersion    string            `json:"androidVersion" db:"android_version"`
	ContentRating     string            `json:"contentRating" db:"content_rating"`
	DeveloperContacts DeveloperContacts `json:"developerContacts" db:"developer_contacts"`
	PrivacyPolicy     string            `json:"privacyPolicy,omitempty"`
	Date              time.Time         `json:"date,omitempty"`
}

type Track struct {
	Bundle string    `json:"bundle,omitempty"`
	Type   string    `json:"type,omitempty"`
	Date   time.Time `json:"date,omitempty"`
	Place  int32     `json:"place,omitempty"`
}

type DeveloperContacts struct {
	Email    string `json:"email,omitempty"`
	Contacts string `json:"contacts,omitempty"`
}