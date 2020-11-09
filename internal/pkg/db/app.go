package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)


// Struct works with app_tracking table
// and can insert and get from table Apps
type appTable struct {
	db *pgx.Conn
}

// Insert App struct to table. Method create transaction and then commit results.
// In order not to repeat the logic of InsertTx method
func (a *appTable) Insert(ctx context.Context, data interface{}) error {
	tx, err := a.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	err = a.InsertTx(tx, ctx, data)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}


// Method get transaction status from param and execute insert method.
// If errors not presented return no error
func (a *appTable) InsertTx(tx pgx.Tx, ctx context.Context, data interface{}) error {
	app, ok := data.(App)
	if !ok {
		return ErrWrongDataType
	}
	_, err := tx.Exec(
		ctx,
		fmt.Sprint("insert into app_tracking (bundle, category, developerId, developer, geo, startAt, period)  values ($1, $2, $3, $4, $5, $6, $7)"),
		app.Bundle, app.Category, app.DeveloperId, app.Developer, app.Geo, app.StartAt, app.Period,
	)
	if err != nil {
		return err
	}

	return nil
}


// Get method can get App struct from table according to his bundle
func (a *appTable) Get(ctx context.Context, value interface{}) (interface{}, error) {
	bundle, ok := value.(string)
	if !ok {
		return nil, ErrWrongDataType
	}
	var app App
	err := a.db.QueryRow(ctx, fmt.Sprint("select * from app_tracking where bundle = $1"), bundle).
		Scan(
			&app.Id,
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

// Create new instance of appTable implementation
func NewAppTracking(db *pgx.Conn) *appTable {
	return &appTable{
		db: db,
	}
}
