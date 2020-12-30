package db

import (
	"Samurai/config"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"strings"
)

// ConnectionUrl creates database connection url by given config.DBConfig
func ConnectionUrl(config config.DBConfig) (string, error) {
	url := "postgresql://"

	if config.User != "" && config.Password != "" {
		url += config.User + ":" + config.Password + "@"
	}

	if config.Address == "" {
		return "", errors.New("empty db address")
	}
	url += config.Address

	if config.Port == "" {
		return "", errors.New("empty db port")
	}
	url += ":" + config.Port + "/"

	db := config.Database
	if db == "" {
		db = "default"
	}
	url += db

	return url, nil
}

// Connect connect to database by url
func Connect(url string) (*pgx.Conn, error) {
	connect, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return connect, nil
}

// InitSchema creates database schema from shemafile string
func InitSchema(connection *pgx.Conn, schemafile string) error {
	b, err := ioutil.ReadFile(schemafile)
	if err != nil {
		return err
	}
	schema := string(b)
	tables := strings.Split(schema, ";\r\n")

	ctx := context.Background()
	for _, v := range tables {
		_, err = connection.Exec(ctx, v)
		if err != nil {
			if strings.Contains(err.Error(), "developercontacts") && strings.Contains(err.Error(), "42710")  {
				continue
			}
			return err
		}
	}


	return nil
}
