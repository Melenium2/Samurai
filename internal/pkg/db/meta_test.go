package db_test

import (
	"Samurai/internal/pkg/db"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mailru/go-clickhouse"
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestInsert_ShouldCreateNewMetaRecordInDatabase_NoError(t *testing.T) {
	sqlDb, mock := MockDb()

	info := NewMeta()
	mock.ExpectExec("^insert into meta_tracking values \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)$").
		WithArgs(
			info.Bundle,
			info.Title,
			info.Price,
			info.Picture,
			clickhouse.Array(info.Screenshots),
			info.Rating,
			info.ReviewCount,
			clickhouse.Array(info.RatingHistogram),
			info.Description,
			info.ShortDescription,
			info.RecentChanges,
			info.ReleaseDate,
			info.LastUpdateDate,
			info.AppSize,
			info.Installs,
			info.Version,
			info.AndroidVersion,
			info.ContentRating,
			clickhouse.Array([]string{info.DeveloperContacts.Email}),
			clickhouse.Array([]string{info.DeveloperContacts.Contacts}),
			info.PrivacyPolicy,
			clickhouse.Date(info.Date),
		).
		WillReturnResult(sqlmock.NewErrorResult(nil))

	tracking := db.NewMetaTracking(sqlDb)
	assert.NoError(t, tracking.Insert(context.Background(), info))
}
