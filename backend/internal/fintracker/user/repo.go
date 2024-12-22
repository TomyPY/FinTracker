package user

import (
	"context"
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
	GetByUsername(ctx context.Context, username string) (User, error)
	GetByID(ctx context.Context, id uint64) (User, error)
	Create(ctx context.Context, user User) error
	Delete(id int) error
	Update(id int, user User) (int, error)
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByUsername(ctx context.Context, username string) (User, error) {

	stmt, err := r.db.PrepareContext(ctx, getUserQuery)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, username)

	u := User{}
	err = row.Scan(&u.ID, &u.Username, &u.Password, &u.Role)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNotFound
		}
		return User{}, err
	}

	return u, nil
}

func (r *repository) GetByID(ctx context.Context, id uint64) (User, error) {
	stmt, err := r.db.PrepareContext(ctx, getUserByIDQuery)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	u := User{}
	err = row.Scan(&u.ID, &u.Username, &u.Password, &u.Role)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNotFound
		}
		return User{}, err
	}

	return u, nil
}

func (r *repository) Create(ctx context.Context, user User) error {
	//Create sql tx
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// //Create a wallet for the user
	// queryWallet := "INSERT INTO wallets(money) VALUES (?)"
	// stmtWallet, err := tx.Prepare(queryWallet)
	// if err != nil {
	// 	tx.Rollback()
	// 	return 0, err
	// }
	// defer stmtWallet.Close()

	// result, err := stmtWallet.Exec(0)
	// if err != nil {
	// 	tx.Rollback()
	// 	return 0, err
	// }

	// //Store walletID
	// walletId, err := result.LastInsertId()
	// if err != nil {
	// 	tx.Rollback()
	// 	return 0, err
	// }

	// //Create the user

	stmtUser, err := tx.PrepareContext(ctx, createUserQuery)
	if err != nil {

		return err
	}
	defer stmtUser.Close()

	_, err = stmtUser.ExecContext(ctx, user.Username, user.Password)
	if err != nil {

		return err
	}

	return tx.Commit()
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
