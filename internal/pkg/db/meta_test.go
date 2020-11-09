package db_test

import (
	"Samurai/internal/pkg/db"
	"time"
)

func NewMeta() db.Meta {
	return db.Meta{
		Bundle:           "com.bundle.go",
		Title:            "title",
		Price:            "4.3$",
		Picture:          "wowpiture",
		Screenshots:      []string{"1", "2", "3"},
		Rating:           "4.3",
		ReviewCount:      "40900",
		RatingHistogram:  []string{"1", "2", "3"},
		Description:      "some text",
		ShortDescription: "text",
		RecentChanges:    "wow it is text?",
		ReleaseDate:      "2020-01-01",
		LastUpdateDate:   "2020-12-12",
		AppSize:          "49M",
		Installs:         "30000000+",
		Version:          "1.3.rd",
		AndroidVersion:   "4.3+",
		ContentRating:    "18+",
		DeveloperContacts: db.DeveloperContacts{
			Email:    "email@email.com",
			Contacts: "Alaska 321",
		},
		PrivacyPolicy: "http://url.com",
		Date:          time.Now(),
	}
}
