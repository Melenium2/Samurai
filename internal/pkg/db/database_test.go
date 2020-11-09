package db_test

import (
	"Samurai/internal/pkg/db"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mailru/go-clickhouse"
	"github.com/stretchr/testify/assert"
	"testing"
)

type getter_mock struct {
}

func (g getter_mock) Insert(ctx context.Context, data interface{}) error {
	panic("implement me")
}

func (g getter_mock) Get(ctx context.Context, value interface{}) (interface{}, error) {
	panic("implement me")
}

type inserter_mock struct {
}

func (i inserter_mock) Insert(ctx context.Context, data interface{}) error {
	panic("implement me")
}

func TestTrackingDatabase_ShouldInsertMetaToMetaTable_NoError(t *testing.T) {
	sqlDb, mock := MockDb()
	repo := db.New(getter_mock{}, db.NewMetaTracking(sqlDb), inserter_mock{}, inserter_mock{})

	info := NewMeta()
	mock.ExpectExec("^insert into meta_tracking values \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)$").
		WithArgs(info.Bundle,
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

	assert.NoError(t, repo.Insert(context.Background(), info))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTrackingDatabase_ShouldInsertAppToAppTable_NoError(t *testing.T) {
	sqlDb, mock := MockDb()
	repo := db.New(db.NewAppTracking(sqlDb), inserter_mock{}, inserter_mock{}, inserter_mock{})

	app := NewApp()
	mock.ExpectExec("insert into app_tracking values \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").
		WithArgs(app.Bundle, app.Category, app.DeveloperId, app.Developer, app.Geo, clickhouse.Date(app.StartAt), app.Period).
		WillReturnResult(sqlmock.NewErrorResult(nil))

	assert.NoError(t, repo.Insert(context.Background(), app))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTrackingDatabase_ShouldInsertCategoryToCategoryTable_NoError(t *testing.T) {
	sqlDb, mock := MockDb()
	repo := db.New(getter_mock{}, inserter_mock{}, inserter_mock{}, db.NewCategoryTracking(sqlDb))

	track := NewTrack()
	mock.ExpectExec("^insert into category_tracking values \\(\\?, \\?, \\?, \\?\\)$").
		WithArgs(
			track.Bundle,
			track.Type,
			track.Place,
			clickhouse.Date(track.Date),
		).
		WillReturnResult(sqlmock.NewErrorResult(nil))

	assert.NoError(t, repo.Insert(context.Background(), track))
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestTrackingDatabase_ShouldInsertKeywordToKeywordsTable_NoError(t *testing.T) {
	sqlDb, mock := MockDb()
	repo := db.New(getter_mock{}, inserter_mock{}, db.NewKeywordsTracking(sqlDb), inserter_mock{})

	track := NewTrack()
	track.Type = "key"
	mock.ExpectExec("^insert into keyword_tracking values \\(\\?, \\?, \\?, \\?\\)$").
		WithArgs(
			track.Bundle,
			track.Type,
			track.Place,
			clickhouse.Date(track.Date),
		).
		WillReturnResult(sqlmock.NewErrorResult(nil))

	assert.NoError(t, repo.Insert(context.Background(), track))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestTrackingDatabase_ShouldReturnErrorCozInvalidInsertType_Error(t *testing.T) {
	repo := db.New(getter_mock{}, inserter_mock{}, inserter_mock{}, inserter_mock{})

	assert.Error(t, repo.Insert(context.Background(), []string{}))
}