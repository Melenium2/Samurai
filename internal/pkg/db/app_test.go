package db_test

import (
	"Samurai/config"
	"Samurai/internal/pkg/db"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
	"time"
)

func RealDb() (*pgx.Conn, func(names ...string)) {
	c := config.New("../../../config/dev.yml")
	url, err := db.ConnectionUrl(c.Database)
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		panic(err)
	}

	if err := db.InitSchema(conn, "../../../config/schema.sql"); err != nil {
		panic(err)
	}

	return conn, func(names ...string) {
		_, err := conn.Exec(context.Background(), fmt.Sprintf("truncate table %s CASCADE", strings.Join(names, ",")))
		if err != nil {
			log.Print(err)
		}
	}
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

func TestAppTableInsert_ShouldInsertRowWithoutError_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("app_tracking")

	repo := db.NewAppTracking(conn)

	app := NewApp()

	_, err := repo.Insert(context.Background(), app)
	assert.NoError(t, err)
}

func TestAppTableGet_ShouldGetLastInsertedRow_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("app_tracking")

	repo := db.NewAppTracking(conn)
	app := NewApp()
	app.Bundle = "super.bundle.for.test"

	_, err := repo.Insert(context.Background(), app)
	assert.NoError(t, err)
	res, err := repo.Get(context.Background(), "super.bundle.for.test")
	assert.NoError(t, err)
	assert.NotNil(t, res)
	a, ok := res.(db.App)
	assert.True(t, ok)
	assert.Equal(t, app.Bundle, a.Bundle)
}

func TestAppTableGet_ShouldReturnErrorCozEmptyTable_Error(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("app_tracking")

	repo := db.NewAppTracking(conn)
	app := NewApp()
	app.Bundle = "super.bundle.for.test"

	res, err := repo.Get(context.Background(), "super.bundle.for.test")
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestAppTableInsertTx_ShouldInsertRowWithoutError_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("app_tracking")

	repo := db.NewAppTracking(conn)

	app := NewApp()

	ctx := context.Background()
	tx, _ := conn.Begin(ctx)

	_, err := repo.InsertTx(tx, ctx, app)
	assert.NoError(t, err)
	assert.NoError(t, tx.Commit(ctx))
}

func TestAppTableInsertTx_ShouldInsertSomeRows_NoError(t *testing.T) {
	conn, cleaner := RealDb()
	defer cleaner("app_tracking")

	repo := db.NewAppTracking(conn)

	app := NewApp()
	ctx := context.Background()
	var num int
	sql := "select count(*) from app_tracking"

	tx, _ := conn.Begin(ctx)
	_, err := repo.InsertTx(tx, ctx, app)
	assert.NoError(t, err)
	_, err = repo.InsertTx(tx, ctx, app)
	assert.NoError(t, err)
	_, err = repo.InsertTx(tx, ctx, app)
	assert.NoError(t, err)
	assert.NoError(t, tx.Commit(ctx))

	conn.QueryRow(context.Background(), sql).Scan(&num)
	assert.Equal(t, 3, num)
}