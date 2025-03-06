package database

import (
	"fmt"
	"os"
	"strings"
)

func Write(data map[string]any) error {

	// Connect to clickhosue
	conn, err := Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	// Prepare SQL
	columns := []string{"user_id", "severity", "user_ip", "action", "content", "agent", "timestamp"}
	placeholders := make([]string, len(columns))
	values := make([]any, len(columns))

	for i, col := range columns {
		placeholders[i] = "?"
		values[i] = data[col]
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", os.Getenv("CLICKHOUSE_TABLE"),
		strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	// Execute the query
	_, err = conn.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("failed to insert log: %v", err)
	}

	return nil
}
