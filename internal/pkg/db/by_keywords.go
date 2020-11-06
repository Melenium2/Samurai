package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mailru/go-clickhouse"
)

type keywordsTable struct {
	db *sql.DB
}

func (k *keywordsTable) Insert(ctx context.Context, data interface{}) error {
	track, ok := data.(Track)
	if !ok {
		return ErrWrongDataType
	}

	_, err := k.db.ExecContext(
		ctx,
		fmt.Sprint("insert into keyword_tracking values (?, ?, ?, ?)"),
		track.Bundle,
		track.Type,
		track.Place,
		clickhouse.Date(track.Date),
	)
	if err != nil {
		return err
	}

	return nil
}

func NewKeywordsTracking(db *sql.DB) *keywordsTable {
	return &keywordsTable{
		db: db,
	}
}