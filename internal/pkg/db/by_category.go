package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type categoryTable struct {
	db *pgx.Conn
}


func (k *categoryTable) Insert(ctx context.Context, data interface{}) error {
	track, ok := data.(Track)
	if !ok {
		return ErrWrongDataType
	}

	_, err := k.db.Exec(
		ctx,
		fmt.Sprint("insert into category_tracking values ($1, $2, $3, $4)"),
		track.Bundle,
		track.Type,
		track.Place,
		track.Date,
	)
	if err != nil {
		return err
	}

	return nil
}

func (k *categoryTable) InsertTx(tx pgx.Tx, ctx context.Context, data interface{}) error {
	track, ok := data.(Track)
	if !ok {
		return ErrWrongDataType
	}

	_, err := tx.Exec(
		ctx,
		fmt.Sprint("insert into category_tracking values ($1, $2, $3, $4)"),
		track.Bundle,
		track.Type,
		track.Place,
		track.Date,
	)
	if err != nil {
		return err
	}

	return nil
}

func NewCategoryTracking(db *pgx.Conn) *categoryTable {
	return &categoryTable{
		db: db,
	}
}