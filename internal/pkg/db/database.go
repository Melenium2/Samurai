package db

import (
	"Samurai/config"
	"context"
	"database/sql"
	"fmt"
	"strings"
)


var ErrWrongDataType = fmt.Errorf("wrong data type")

type Inserter interface {
	Insert(ctx context.Context, data interface{}) error
	InsertTx(tx *sql.Tx, ctx context.Context,  data interface{}) error
}

type Getter interface {
	Inserter
	Get(ctx context.Context, value interface{}) (interface{}, error)
}

type Tracking interface {
}

type TrackingDatabase struct {
	app      Getter
	meta     Inserter
	keywords Inserter
	category Inserter
}

func (t *TrackingDatabase) Insert(ctx context.Context, data interface{}) error {
	switch v := data.(type) {
	case App:
		return t.app.Insert(ctx, v)
	case Meta:
		return t.meta.Insert(ctx, v)
	case Track:
		splited := strings.Split(v.Type, "_")
		if len(splited) >= 2 {
			return t.category.Insert(ctx, v)
		}

		return t.keywords.Insert(ctx, v)
	}

	return ErrWrongDataType
}

func (t *TrackingDatabase) InsertTx(tx *sql.Tx, ctx context.Context,  data interface{}) error {
	switch v := data.(type) {
	case App:
		return t.app.InsertTx(tx, ctx, v)
	case Meta:
		return t.meta.InsertTx(tx, ctx, v)
	case Track:
		splited := strings.Split(v.Type, "_")
		if len(splited) >= 2 {
			return t.category.InsertTx(tx, ctx, v)
		}

		return t.keywords.InsertTx(tx, ctx, v)
	}

	return ErrWrongDataType
}

func NewWithConfig(config config.DBConfig) *TrackingDatabase {
	url, err := ConnectionUrl(config)
	if err != nil {
		panic(err)
	}
	conn, err := Connect(url, "clickhouse")
	if err != nil {
		panic(err)
	}

	return NewWithConnection(conn)
}

func NewWithConnection(db *sql.DB) *TrackingDatabase {
	return New(
		NewAppTracking(db),
		NewMetaTracking(db),
		NewKeywordsTracking(db),
		NewCategoryTracking(db),
	)
}

func New(app Getter, meta Inserter, keys Inserter, cats Inserter) *TrackingDatabase {
	return &TrackingDatabase{
		app, meta, keys, cats,
	}
}
