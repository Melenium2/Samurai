package db_test

import (
	"Samurai/config"
	"Samurai/internal/pkg/db"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitSchema_ShouldCreateNewSchemaAndExecQueryToTable(t *testing.T) {
	c := config.New("../../../config/dev.yml")
	url, err := db.ConnectionUrl(c.Database)
	assert.NoError(t, err)
	conn, err := db.Connect(url)
	assert.NoError(t, err)
	assert.NotNil(t, conn)
	assert.NoError(t, db.InitSchema(conn, "../../../config/schema.sql"))

	_, err = conn.Exec(context.Background(), "select count(*) from app_tracking")
	assert.NoError(t, err)
}
