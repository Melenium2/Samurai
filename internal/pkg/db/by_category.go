package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mailru/go-clickhouse"
)

type categoryTable struct {
	db *sql.DB
}

func (k *categoryTable) Insert(ctx context.Context, data interface{}) error {
	track, ok := data.(Track)
	if !ok {
		return ErrWrongDataType
	}

	_, err := k.db.ExecContext(
		ctx,
		fmt.Sprint("insert into category_tracking values (?, ?, ?, ?)"),
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

func NewCategoryTracking(db *sql.DB) *categoryTable {
	return &categoryTable{
		db: db,
	}
}