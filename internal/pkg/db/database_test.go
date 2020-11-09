package db_test

import (
	"context"
	"database/sql"
)

type getter_mock struct {
}

func (g getter_mock) InsertTx(tx *sql.Tx, ctx context.Context, data interface{}) error {
	panic("implement me")
}

func (g getter_mock) Insert(ctx context.Context, data interface{}) error {
	panic("implement me")
}

func (g getter_mock) Get(ctx context.Context, value interface{}) (interface{}, error) {
	panic("implement me")
}

type inserter_mock struct {
}

func (i inserter_mock) InsertTx(tx *sql.Tx, ctx context.Context, data interface{}) error {
	panic("implement me")
}

func (i inserter_mock) Insert(ctx context.Context, data interface{}) error {
	panic("implement me")
}

