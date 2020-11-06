package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/mailru/go-clickhouse"
)

type appTable struct {
	db *sql.DB
}

func (a *appTable) Insert(ctx context.Context, data interface{}) error {
	app, ok := data.(App)
	if !ok {
		return ErrWrongDataType
	}
	_, err := a.db.ExecContext(
		ctx,
		fmt.Sprint("insert into app_tracking values (?, ?, ?, ?, ?, ?, ?)"),
		app.Bundle, app.Category, app.DeveloperId, app.Developer, app.Geo, clickhouse.Date(app.StartAt), app.Period,
	)
	if err != nil {
		return err
	}

	return nil
}

func (a *appTable) Get(ctx context.Context, value interface{}) (interface{}, error) {
	bundle, ok := value.(string)
	if !ok {
		return nil, ErrWrongDataType
	}
	var app App
	err := a.db.QueryRowContext(ctx, fmt.Sprint("select * from app_tracking where bundle = $1"), bundle).
		Scan(
			&app.Bundle,
			&app.Category,
			&app.DeveloperId,
			&app.Developer,
			&app.Geo,
			&app.StartAt,
			&app.Period,
		)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func NewAppTracking(db *sql.DB) *appTable {
	return &appTable{
		db: db,
	}
}
