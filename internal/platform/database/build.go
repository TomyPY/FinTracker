package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	username string
	password string
	host     string
	name     string
}

func NewDb(options ...func(*Database)) *Database {
	db := &Database{}

	for _, o := range options {
		o(db)
	}

	return db
}

func (db *Database) getConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", db.username, db.password, db.host, db.name)
}

func (db *Database) Init() (*sql.DB, error) {
	databaseSql, err := sql.Open("mysql", db.getConnectionString())
	if err != nil {
		log.Fatal("Connection failed with db")
		return nil, err
	}

	if err := databaseSql.Ping(); err != nil {
		log.Fatal("Ping to db failed")
		return nil, err
	}

	return databaseSql, nil
}
