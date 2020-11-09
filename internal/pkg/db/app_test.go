package db_test

import (
	"Samurai/internal/pkg/db"
	"context"
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mailru/go-clickhouse"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func MockDb() (*sql.DB, sqlmock.Sqlmock) {
	db, m, _ := sqlmock.New()

	return db, m
}

func NewApp() db.App {
	t, _ := time.Parse("2006-01-01", "2020-01-01")
	return db.App{
		Bundle:      "com.bundle.go",
		Category:    "FINANCE",
		DeveloperId: "92834848476158744",
		Developer:   "Random",
		Geo:         "ru_ru",
		StartAt:     t,
		Period:      31,
	}
}

func TestInsert_ShouldInsertNewRecordToDb_NoError(t *testing.T) {
	sqlDb, mock := MockDb()

	app := NewApp()
	mock.ExpectExec("^insert into app_tracking values \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?\\)$").
		WithArgs(
			app.Bundle,
			app.Category,
			app.DeveloperId,
			app.Developer,
			app.Geo,
			clickhouse.Date(app.StartAt),
			app.Period,
		).WillReturnResult(sqlmock.NewErrorResult(nil))

	tracking := db.NewAppTracking(sqlDb)
	assert.NoError(t, tracking.Insert(context.Background(), app))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGet_ShouldGetValueFromDbByBundle_NoError(t *testing.T) {
	sqlDb, mock := MockDb()

	app := NewApp()
	mock.ExpectQuery("^select \\* from app_tracking where bundle = \\$1$").
		WithArgs("com.bundle.go").
		WillReturnRows(
			sqlmock.NewRows([]string{"bundle", "category", "developerId", "developer", "geo", "startAt", "period"}).
				AddRow(
					app.Bundle,
					app.Category,
					app.DeveloperId,
					app.Developer,
					app.Geo,
					clickhouse.Date(app.StartAt),
					app.Period,
				),
		)

	traking := db.NewAppTracking(sqlDb)
	aIn, err := traking.Get(context.Background(), "com.bundle.go")
	assert.NoError(t, err)
	assert.NotNil(t, aIn)

	a, ok := aIn.(db.App)
	assert.True(t, ok)
	assert.Equal(t, app.Bundle, a.Bundle)
}

func TestGet_ShouldReturnNoRowInDbError_NoError(t *testing.T) {
	sqlDb, mock := MockDb()

	mock.ExpectQuery("^select \\* from app_tracking where bundle = \\$1$").
		WithArgs("com.bundle.go123").
		WillReturnRows(sqlmock.NewRows([]string{})).
		WillReturnError(fmt.Errorf(""))

	traking := db.NewAppTracking(sqlDb)
	aIn, err := traking.Get(context.Background(), "com.bundle.go123")
	assert.Error(t, err)

	a, ok := aIn.(db.App)
	assert.False(t, ok)
	assert.Empty(t, a.Bundle)
}

func TestInsertTx_ShouldInsertNewRowInTx_NoError(t *testing.T) {
	sqlDb, mock := MockDb()

	app := NewApp()
	mock.ExpectBegin()
	mock.ExpectExec("^insert into app_tracking values \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?\\)$").
		WithArgs(
			app.Bundle,
			app.Category,
			app.DeveloperId,
			app.Developer,
			app.Geo,
			clickhouse.Date(app.StartAt),
			app.Period,
		).
		WillReturnResult(sqlmock.NewErrorResult(nil))
	mock.ExpectCommit()

	repo := db.NewAppTracking(sqlDb)
	tx, _ := sqlDb.Begin()
	assert.NoError(t, repo.InsertTx(tx, context.Background(), app))
	assert.NoError(t, tx.Commit())
	assert.NoError(t, mock.ExpectationsWereMet())
}
