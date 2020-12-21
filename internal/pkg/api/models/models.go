package models

import (
	"fmt"
	"strings"
)

type Category string

func (c Category) Split() (string, string) {
	splited := strings.Split(string(c), "|")
	if len(splited) > 2 {
		panic("invalid category")
	}
	return splited[0], splited[1]
}

func NewCategory(cat, subcat string) Category {
	return Category(strings.ToLower(fmt.Sprintf("%s|%s", cat, subcat)))
}

type App struct {
	Bundle            string            `json:"bundle" db:"bundle"`
	DeveloperId       string            `json:"developerId" db:"developer_id"`
	Developer         string            `json:"developer" db:"developer"`
	Title             string            `json:"title" db:"title"`
	Categories        string            `json:"categories" db:"categories"`
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
}

type DeveloperContacts struct {
	Email    string `json:"email,omitempty"`
	Contacts string `json:"contacts,omitempty"`
}