package db_test

import (
	"Samurai/config"
	"Samurai/internal/pkg/db"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

type getter_mock struct {
}

func (g getter_mock) Insert(ctx context.Context, data interface{}) (int, error) {
	panic("implement me")
}

func (g getter_mock) InsertTx(tx pgx.Tx, ctx context.Context, data interface{}) (int, error) {
	panic("implement me")
}

func (g getter_mock) Get(ctx context.Context, value interface{}) (interface{}, error) {
	panic("implement me")
}

type inserter_mock struct {
}

func (i inserter_mock) Insert(ctx context.Context, data interface{}) (int, error) {
	panic("implement me")
}

func (i inserter_mock) InsertTx(tx pgx.Tx, ctx context.Context, data interface{}) (int, error) {
	panic("implement me")
}

func TestTrackingDatabaseNew_ShouldCreateInstanceByConnAndConfig_NoError(t *testing.T) {
	c := config.New()
	repo := db.NewWithConfig(c.Database)

	app := NewApp()
	id, err := repo.Insert(context.Background(), app)
	assert.NoError(t, err)

	conn, cleaner := RealDb()
	defer cleaner("meta_tracking", "app_tracking")
	repoConn := db.NewWithConnection(conn)

	meta := NewMeta()
	meta.BundleId = id
	_, err = repoConn.Insert(context.Background(), meta)
	assert.NoError(t, err)
}

func TestTrackingDatabaseInsert_ShouldInsertApp_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("app_tracking")

	app := NewApp()
	repo := db.New(db.NewAppTracking(conn), inserter_mock{}, inserter_mock{}, inserter_mock{})

	_, err := repo.Insert(context.Background(), app)
	assert.NoError(t, err)
}

func TestTrackingDatabaseInsert_ShouldInsertMeta_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("meta_tracking", "app_tracking")

	app := NewApp()
	meta := NewMeta()
	repo := db.New(db.NewAppTracking(conn), db.NewMetaTracking(conn), inserter_mock{}, inserter_mock{})

	id, err := repo.Insert(context.Background(), app)

	meta.BundleId = id
	_, err = repo.Insert(context.Background(), meta)
	assert.NoError(t, err)
}

func TestTrackingDatabaseInsert_ShouldInsertKeywordAndCategory_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("category_tracking", "keyword_tracking", "app_tracking")

	app := NewApp()
	track := NewTrack()
	repo := db.New(db.NewAppTracking(conn), inserter_mock{}, db.NewKeywordsTracking(conn), db.NewCategoryTracking(conn))

	id, err := repo.Insert(context.Background(), app)
	track.BundleId = id

	_, err = repo.Insert(context.Background(), track)
	assert.NoError(t, err)
	track.Type = "key"
	_, err = repo.Insert(context.Background(), track)
	assert.NoError(t, err)
}

func TestTrackingDatabaseInsertTx_ShouldInsertInfoInTransaction_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("meta_tracking", "category_tracking", "keyword_tracking")

	app := NewApp()
	meta := NewMeta()
	track := NewTrack()
	repo := db.NewWithConnection(conn)

	id, err := repo.Insert(context.Background(), app)
	meta.BundleId = id
	track.BundleId = id

	ctx := context.Background()
	tx, _ := conn.Begin(ctx)

	_, err = repo.InsertTx(tx, context.Background(), meta)
	assert.NoError(t, err)
	_, err = repo.InsertTx(tx, context.Background(), track)
	assert.NoError(t, err)

	track.Type = "key"
	_, err = repo.InsertTx(tx, context.Background(), track)
	assert.NoError(t, err)

	assert.NoError(t, tx.Commit(ctx))
}
