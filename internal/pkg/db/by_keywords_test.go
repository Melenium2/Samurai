package db_test

import (
	"Samurai/internal/pkg/db"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func NewTrack() db.Track {
	return db.Track{
		BundleId: 1,
		Type:     "game_free_how",
		Date:     time.Now(),
		Place:    43,
	}
}

func TestKeywordTableInsert_ShouldInsertNewRow_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("keyword_tracking", "app_tracking")

	app := NewApp()
	track := NewTrack()

	repoApp := db.NewAppTracking(conn)
	id, err := repoApp.Insert(context.Background(), app)
	assert.NoError(t, err)

	track.BundleId = id

	repoKeyword := db.NewKeywordsTracking(conn)
	_, err = repoKeyword.Insert(context.Background(), track)
	assert.NoError(t, err)
}

func TestKeywordTableInsertTx_ShouldInsertRowsInTransaction_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("keyword_tracking", "app_tracking")

	app := NewApp()
	track := NewTrack()

	ctx := context.Background()
	repoApp := db.NewAppTracking(conn)
	repoKeyword := db.NewKeywordsTracking(conn)

	tx, _ := conn.Begin(ctx)
	id, err := repoApp.InsertTx(tx, ctx, app)
	assert.NoError(t, err)
	track.BundleId = id
	_, err = repoKeyword.InsertTx(tx, ctx, track)
	assert.NoError(t, err)
	assert.NoError(t, tx.Commit(ctx))
}

func TestKeywordTableInsertTx_ShouldNotSaveCategoryIfRollback_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("keyword_tracking", "app_tracking")

	app := NewApp()
	track := NewTrack()

	ctx := context.Background()
	repoApp := db.NewAppTracking(conn)
	repoKeyword := db.NewKeywordsTracking(conn)

	tx, _ := conn.Begin(ctx)
	id, err := repoApp.InsertTx(tx, ctx, app)
	assert.NoError(t, err)
	track.BundleId = id
	_, err = repoKeyword.InsertTx(tx, ctx, track)
	assert.NoError(t, err)

	var ids int
	conn.
		QueryRow(context.Background(), "select count(*) from keyword_tracking").
		Scan(&ids)
	assert.Equal(t, 1, ids)

	assert.NoError(t, tx.Rollback(ctx))

	conn.
		QueryRow(context.Background(), "select count(*) from keyword_tracking").
		Scan(&ids)
	assert.Equal(t, 0, ids)
}

