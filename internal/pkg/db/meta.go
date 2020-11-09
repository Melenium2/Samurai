package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mailru/go-clickhouse"
)

type metaTable struct {
	db *sql.DB
}


func (m *metaTable) Insert(ctx context.Context, data interface{}) error {
	info, ok := data.(Meta)
	if !ok {
		return ErrWrongDataType
	}
	_, err := m.db.ExecContext(
		ctx,
		fmt.Sprint("insert into meta_tracking values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"),
		info.Bundle,
		info.Title,
		info.Price,
		info.Picture,
		clickhouse.Array(info.Screenshots),
		info.Rating,
		info.ReviewCount,
		clickhouse.Array(info.RatingHistogram),
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
		clickhouse.Array([]string{ info.DeveloperContacts.Email }),
		clickhouse.Array([]string{ info.DeveloperContacts.Contacts }),
		info.PrivacyPolicy,
		clickhouse.Date(info.Date),
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *metaTable) InsertTx(tx *sql.Tx, ctx context.Context, data interface{}) error {
	info, ok := data.(Meta)
	if !ok {
		return ErrWrongDataType
	}
	_, err := tx.ExecContext(
		ctx,
		fmt.Sprint("insert into meta_tracking values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"),
		info.Bundle,
		info.Title,
		info.Price,
		info.Picture,
		clickhouse.Array(info.Screenshots),
		info.Rating,
		info.ReviewCount,
		clickhouse.Array(info.RatingHistogram),
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
		clickhouse.Array([]string{ info.DeveloperContacts.Email }),
		clickhouse.Array([]string{ info.DeveloperContacts.Contacts }),
		info.PrivacyPolicy,
		clickhouse.Date(info.Date),
	)

	if err != nil {
		return err
	}

	return nil
}

func NewMetaTracking(db *sql.DB) *metaTable {
	return &metaTable{
		db: db,
	}
}
