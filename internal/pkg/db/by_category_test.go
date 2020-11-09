package db_test

import (
	"Samurai/internal/pkg/db"
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mailru/go-clickhouse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert_ShouldInsertNewRowToCategoryTable_NoError(t *testing.T) {
	sqlDb, mock := MockDb()

	track := NewTrack()

	mock.ExpectExec("^insert into category_tracking values \\(\\?, \\?, \\?, \\?\\)$").
		WithArgs(
			track.Bundle,
			track.Type,
			track.Place,
			clickhouse.Date(track.Date),
		).
		WillReturnResult(sqlmock.NewErrorResult(nil))

	tracking := db.NewCategoryTracking(sqlDb)
	assert.NoError(t, tracking.Insert(context.Background(), track))
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestInsertTx_ShouldInsertNewRowToCategoryTable_NoError(t *testing.T) {
	sqlDb, mock := MockDb()

	track := NewTrack()

	mock.ExpectBegin()
	mock.ExpectExec("^insert into category_tracking values \\(\\?, \\?, \\?, \\?\\)$").
		WithArgs(
			track.Bundle,
			track.Type,
			track.Place,
			clickhouse.Date(track.Date),
		).
		WillReturnResult(sqlmock.NewErrorResult(nil))
	mock.ExpectCommit()

	tracking := db.NewCategoryTracking(sqlDb)
	tx, _ := sqlDb.Begin()
	assert.NoError(t, tracking.InsertTx(tx, context.Background(), track))
	assert.NoError(t, tx.Commit())
	assert.NoError(t, mock.ExpectationsWereMet())
}



