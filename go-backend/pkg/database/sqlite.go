package database

import (
	"database/sql"
	"fmt"
)

func InitDB(sqlite3fileName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", sqlite3fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(255) NOT NULL
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	return db, nil
}
