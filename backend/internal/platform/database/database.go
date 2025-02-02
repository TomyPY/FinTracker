package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDb(conn string) (*sql.DB, error) {

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
