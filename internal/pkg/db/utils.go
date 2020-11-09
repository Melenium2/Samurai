package db

import (
	"Samurai/config"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"strings"
)

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

func Connect(url string) (*pgx.Conn, error) {
	connect, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return connect, nil
}

func InitSchema(connection *pgx.Conn, schemafile string) error {
	b, err := ioutil.ReadFile(schemafile)
	if err != nil {
		return err
	}
	schema := string(b)
	tables := strings.Split(schema, ";\n")

	ctx := context.Background()
	for _, v := range tables {
		_, err = connection.Exec(ctx, v)
		if err != nil {
			if strings.Contains(err.Error(), "Code: 57") {
				newsql := strings.ReplaceAll(v,"CREATE TABLE", "ATTACH TABLE")
				if _, err = connection.Exec(ctx, newsql); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}


	return nil
}
