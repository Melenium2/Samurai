package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type metaTable struct {
	db *pgx.Conn
}


func (m *metaTable) Insert(ctx context.Context, data interface{}) (int, error) {
	tx, err := m.db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	id, err := m.InsertTx(tx, ctx, data)
	if err != nil {
		return 0, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *metaTable) InsertTx(tx pgx.Tx, ctx context.Context, data interface{}) (int, error) {
	info, ok := data.(Meta)
	if !ok {
		return 0, ErrWrongDataType
	}

	row := tx.QueryRow(
		ctx,
		fmt.Sprint("insert into meta_tracking values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18:developercontacts, $19, $20) returning id"),
		info.Bundle,
		info.Title,
		info.Price,
		info.Picture,
		info.Screenshots,
		info.Rating,
		info.ReviewCount,
		info.RatingHistogram,
		info.Description,
		info.ShortDescription,
		info.RecentChanges,
		info.ReleaseDate,
		info.LastUpdateDate,
		info.AppSize,
		info.Installs,
		info.Version,
		info.AndroidVersion,
		info.ContentRating,
		info.DeveloperContacts,
		info.PrivacyPolicy,
		info.Date,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func NewMetaTracking(db *pgx.Conn) *metaTable {
	return &metaTable{
		db: db,
	}
}
