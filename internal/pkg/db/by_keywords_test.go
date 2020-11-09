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

func NewTrack() db.Track {
	return db.Track{
		Bundle: "com.bundle.go",
		Type:   "game_free_how",
		Date:   time.Now(),
		Place:  43,
	}
}

func TestInsert_ShouldInsertNewRowToKeywordsTable_NoError(t *testing.T) {
	sqlDb, mock := MockDb()

	track := NewTrack()

	mock.ExpectExec("^insert into keyword_tracking values \\(\\?, \\?, \\?, \\?\\)$").
		WithArgs(
			track.Bundle,
			track.Type,
			track.Place,
			clickhouse.Date(track.Date),
		).
		WillReturnResult(sqlmock.NewErrorResult(nil))

	tracking := db.NewKeywordsTracking(sqlDb)
	assert.NoError(t, tracking.Insert(context.Background(), track))
	assert.NoError(t, mock.ExpectationsWereMet())
}
