package storage

import (
	"fmt"
	"database/sql"
)

// CreateConnectionPostgres .
func CreateConnectionPostgres(Hostname string, Port int, Username, Password, Database string, SSLMode string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		Hostname, Port, Username, Password, Database, SSLMode)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}