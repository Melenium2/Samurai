package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type keywordsTable struct {
	db *pgx.Conn
}


func (k *keywordsTable) Insert(ctx context.Context, data interface{}) (int, error) {
	tx, err := k.db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	id, err := k.InsertTx(tx, ctx, data)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (k *keywordsTable) InsertTx(tx pgx.Tx, ctx context.Context, data interface{}) (int, error) {
	track, ok := data.(Track)
	if !ok {
		return 0, ErrWrongDataType
	}

	row := tx.QueryRow(
		ctx,
		fmt.Sprint("insert into keyword_tracking (bundleId, type, place, date) values ($1, $2, $3, $4) returning id"),
		track.BundleId,
		track.Type,
		track.Place,
		track.Date,
	)
	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func NewKeywordsTracking(db *pgx.Conn) *keywordsTable {
	return &keywordsTable{
		db: db,
	}
}