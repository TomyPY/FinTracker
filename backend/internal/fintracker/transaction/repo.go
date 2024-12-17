package transaction

import "database/sql"

type repository struct {
	db *sql.DB
}

type Repository interface {
}
