package database

import (
	"database/sql"
	"os"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func Connect() (*sql.DB, error) {

	// Connect to Clickhouse database
	connection := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"clickhouse:9000"}, // container name and port
		Auth: clickhouse.Auth{
			Database: os.Getenv("CLICKHOUSE_DATABASE"),
			Username: os.Getenv("CLICKHOUSE_USER"),
			Password: os.Getenv("CLICKHOUSE_PASSWORD"),
		},
	})

	err := connection.Ping()
	if err != nil {
		return nil, err
	}

	return connection, nil
}
