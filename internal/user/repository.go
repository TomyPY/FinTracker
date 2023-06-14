package user

import (
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("user doesnt exist")
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	GetAll() ([]User, error)
	Get(id int) (User, error)
	Create(user User) (int, error)
	Delete(id int) error
	Update(id int, user User) (int, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]User, error) {
	query := "SELECT username, wallet_id FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}

		err := rows.Scan(&u.Username, &u.WalletId)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *repository) Get(id int) (User, error) {
	query := "SELECT username, wallet_id FROM users WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	u := User{}
	err = row.Scan(&u.Username, &u.WalletId)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNotFound
		}
		return User{}, err
	}

	return u, nil
}

func (r *repository) Create(user User) (int, error) {
	//Create sql tx
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	//Create a wallet for the user
	queryWallet := "INSERT INTO wallets(money) VALUES (?)"
	stmtWallet, err := tx.Prepare(queryWallet)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmtWallet.Close()

	result, err := stmtWallet.Exec(0)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	//Store walletID
	walletId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	//Create the user
	queryUser := "INSERT INTO users(username, wallet_id) VALUES (?,?)"
	stmtUser, err := tx.Prepare(queryUser)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmtUser.Close()

	result, err = stmtUser.Exec(user.Username, walletId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return int(userId), nil

}

func (r *repository) Delete(id int) error {
	query := "DELETE FROM users WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	aff, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if aff <= 0 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) Update(id int, user User) (int, error) {
	//Create the user
	query := "UPDATE users SET username=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, id)
	if err != nil {
		return 0, err
	}

	aff, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if aff <= 0 {
		return 0, ErrNotFound
	}

	return id, nil
}
