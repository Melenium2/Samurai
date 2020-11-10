package db_test

import (
	"Samurai/internal/pkg/db"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func NewMeta() db.Meta {
	return db.Meta{
		BundleId:         0,
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

func TestMetaTableInsert_ShouldInsertNewRow_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("meta_tracking", "app_tracking")

	app := NewApp()
	meta := NewMeta()

	repoApp := db.NewAppTracking(conn)
	repoMeta := db.NewMetaTracking(conn)

	id, err := repoApp.Insert(context.Background(), app)
	assert.NoError(t, err)

	meta.BundleId = id
	_, err = repoMeta.Insert(context.Background(), meta)
	assert.NoError(t, err)
}

func TestMetaTableInsert_ShouldInsertSomeNewRowWithOneBundleId_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("meta_tracking", "app_tracking")

	app := NewApp()
	meta := NewMeta()

	repoApp := db.NewAppTracking(conn)
	repoMeta := db.NewMetaTracking(conn)

	id, err := repoApp.Insert(context.Background(), app)
	assert.NoError(t, err)

	meta.BundleId = id
	_, err = repoMeta.Insert(context.Background(), meta)
	assert.NoError(t, err)
	_, err = repoMeta.Insert(context.Background(), meta)
	assert.NoError(t, err)
	_, err = repoMeta.Insert(context.Background(), meta)
	assert.NoError(t, err)
}

func TestMetaTableInsertTx_ShouldInsertSomeRowsInTransaction_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("meta_tracking", "app_tracking")

	app := NewApp()
	meta := NewMeta()

	repoApp := db.NewAppTracking(conn)
	repoMeta := db.NewMetaTracking(conn)

	id, err := repoApp.Insert(context.Background(), app)
	assert.NoError(t, err)

	meta.BundleId = id
	ctx := context.Background()
	tx, _ := conn.Begin(ctx)
	_, err = repoMeta.InsertTx(tx, ctx, meta)
	assert.NoError(t, err)
	_, err = repoMeta.InsertTx(tx, ctx, meta)
	assert.NoError(t, err)
	_, err = repoMeta.InsertTx(tx, ctx, meta)
	assert.NoError(t, err)
	_, err = repoMeta.InsertTx(tx, ctx, meta)
	assert.NoError(t, err)

	assert.NoError(t, tx.Commit(ctx))
}

func TestMetaTableInsert_ShouldReturnErrorCozConnectionToDbIsOut_Error(t *testing.T) {
	conn, _ := RealDb()

	app := NewApp()

	repoApp := db.NewAppTracking(conn)
	conn.Close(context.Background())

	id, err := repoApp.Insert(context.Background(), app)
	assert.Error(t, err)
	assert.Equal(t, 0, id)
}