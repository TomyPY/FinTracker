package wallet

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/TomyPY/FinTracker/internal/fintracker/transaction"
)

var (
	ErrNotFound             = errors.New("wallet doesnt exist")
	ErrTxNotInserted        = errors.New("tx wasnt inserted")
	ErrInvalidWalletId      = errors.New("invalid wallet id")
	ErrInvalidTransactionId = errors.New("invalid transaction id")
)

type repository struct {
	db *sql.DB
}

type Repository interface {
	GetAll() ([]Wallet, error)
	Get(id int) (Wallet, error)
	GetReportTransactions(transactionId *int, walletId int) ([]transaction.Transaction, error)
	SubstractMoney(wallet_id int, tx transaction.Transaction) (transaction.Transaction, error)
	AddMoney(wallet_id int, tx transaction.Transaction) (transaction.Transaction, error)
	Delete(id int) error
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Wallet, error) {

	// Make SQL query
	rows, err := r.db.Query("SELECT id, money FROM wallets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//Store SQL data
	wallets := make([]Wallet, 0)
	for rows.Next() {
		w := Wallet{}

		err := rows.Scan(&w.ID, &w.Money)
		if err != nil {
			return nil, err
		}

		wallets = append(wallets, w)
	}
	return wallets, nil
}

func (r *repository) Get(id int) (Wallet, error) {
	// Make SQL query
	query := "SELECT id, money FROM wallets WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return Wallet{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)

	//Store SQL data
	w := Wallet{}
	err = row.Scan(&w.ID, &w.Money)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = ErrNotFound
		}
		return Wallet{}, err
	}

	return w, nil
}

func (r *repository) GetReportTransactions(transactionId *int, walletId int) ([]transaction.Transaction, error) {
	// Make query string
	var query string
	var args []any

	if transactionId != nil {
		query = "SELECT id, amount, type, datetime FROM transactions WHERE id=? AND wallet_id"
		args = append(args, *transactionId, walletId)
	} else {
		query = "SELECT id, amount, type, datetime FROM transactions WHERE wallet_id = ?"
		args = append(args, walletId)
	}

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	//TODO: DETECT WHEN TRANSACTION OR WALLET ISNT FOUND (MAKE TWO ERRORS DIFFERENT ONE FOR TX ANOTHER FOR WALLET)
	//Make SQL query
	rows, err := stmt.Query(args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
	}

	//Store SQL data
	transactionsReports := make([]transaction.Transaction, 0)
	for rows.Next() {
		tx := transaction.Transaction{}
		err := rows.Scan(&tx.ID, &tx.Amount, &tx.Type, &tx.Datetime)
		if err != nil {
			return nil, err
		}
		transactionsReports = append(transactionsReports, tx)
	}

	return transactionsReports, nil
}

func (r *repository) AddMoney(wallet_id int, tx transaction.Transaction) (transaction.Transaction, error) {
	query := "UPDATE wallet SET money = money + ? WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return transaction.Transaction{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(tx.Amount, wallet_id)
	if err != nil {
		return transaction.Transaction{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return transaction.Transaction{}, err
	}

	if rowsAffected <= 0 {
		return transaction.Transaction{}, ErrNotFound
	}
	tx.Datetime = fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))

	err = saveTransaction(r.db, tx)
	if err != nil {
		return transaction.Transaction{}, err
	}

	return tx, nil
}

func (r *repository) SubstractMoney(wallet_id int, tx transaction.Transaction) (transaction.Transaction, error) {
	query := "UPDATE wallet SET money = money - ? WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return transaction.Transaction{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(tx.Amount, wallet_id)
	if err != nil {
		return transaction.Transaction{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return transaction.Transaction{}, err
	}

	if rowsAffected <= 0 {
		return transaction.Transaction{}, ErrNotFound
	}
	tx.Datetime = fmt.Sprint(time.Now().UnixNano() / int64(time.Millisecond))

	err = saveTransaction(r.db, tx)
	if err != nil {
		return transaction.Transaction{}, err
	}

	return tx, nil
}

func (r *repository) Delete(id int) error {
	query := "DELETE FROM wallets WHERE id=?"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return ErrNotFound
	}

	return nil
}

func saveTransaction(db *sql.DB, tx transaction.Transaction) error {
	query := "INSERT INTO transactions(amount,type,datetime) VALUES (?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(tx.Amount, tx.Type, tx.Datetime)
	if err != nil {
		return err
	}

	return nil
}
