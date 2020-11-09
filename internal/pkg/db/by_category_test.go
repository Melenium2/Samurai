package db_test

import (
	"Samurai/internal/pkg/db"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryTableInsert_ShouldInsertNewRow_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("category_tracking", "app_tracking")

	app := NewApp()
	track := NewTrack()

	repoApp := db.NewAppTracking(conn)
	id, err := repoApp.Insert(context.Background(), app)
	assert.NoError(t, err)

	track.BundleId = id

	repoCategory := db.NewCategoryTracking(conn)
	_, err = repoCategory.Insert(context.Background(), track)
	assert.NoError(t, err)
}

func TestCategoryTableInsertTx_ShouldInsertRowsInTransaction_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("category_tracking", "app_tracking")

	app := NewApp()
	track := NewTrack()

	ctx := context.Background()
	repoApp := db.NewAppTracking(conn)
	repoCategory := db.NewCategoryTracking(conn)

	tx, _ := conn.Begin(ctx)
	id, err := repoApp.InsertTx(tx, ctx, app)
	assert.NoError(t, err)
	track.BundleId = id
	_, err = repoCategory.InsertTx(tx, ctx, track)
	assert.NoError(t, err)
	assert.NoError(t, tx.Commit(ctx))
}

func TestCategoryTableInsertTx_ShouldNotSaveCategoryIfRollback_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("category_tracking", "app_tracking")

	app := NewApp()
	track := NewTrack()

	ctx := context.Background()
	repoApp := db.NewAppTracking(conn)
	repoCategory := db.NewCategoryTracking(conn)

	tx, _ := conn.Begin(ctx)
	id, err := repoApp.InsertTx(tx, ctx, app)
	assert.NoError(t, err)
	track.BundleId = id
	_, err = repoCategory.InsertTx(tx, ctx, track)
	assert.NoError(t, err)

	var ids int
	conn.
		QueryRow(context.Background(), "select count(*) from category_tracking").
		Scan(&ids)
	assert.Equal(t, 1, ids)

	assert.NoError(t, tx.Rollback(ctx))

	conn.
		QueryRow(context.Background(), "select count(*) from category_tracking").
		Scan(&ids)
	assert.Equal(t, 0, ids)
}