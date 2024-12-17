package database

import (
	"database/sql"
	"fmt"

	"github.com/TomyPY/FinTracker/cmd/server/config"
	_ "github.com/go-sql-driver/mysql"
)

func NewDb(cfg *config.Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", buildConnectionString(cfg))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func buildConnectionString(cfg *config.Config) string {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.DB.UsernameEnv, cfg.DB.PasswordEnv, cfg.DB.HostEnv, cfg.DB.NameEnv)
	fmt.Println("conn:", conn)
	return conn
}
