package wallet

import "github.com/TomyPY/FinTracker/internal/transaction"

type service struct {
	r Repository
}

type Service interface {
	GetAll() ([]Wallet, error)
	Get(id int) (Wallet, error)
	GetReportTransactions(transactionId *int, walletId int) ([]transaction.Transaction, error)
	SubstractMoney(wallet_id int, tx transaction.Transaction) (transaction.Transaction, error)
	AddMoney(wallet_id int, tx transaction.Transaction) (transaction.Transaction, error)
	Delete(id int) error
}

func NewService(r Repository) Service {
	return &service{
		r: r,
	}
}

func (s *service) GetAll() ([]Wallet, error) {
	wallets, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func (s *service) Get(id int) (Wallet, error) {
	wallet, err := s.r.Get(id)
	if err != nil {
		return Wallet{}, err
	}

	return wallet, nil
}

func (s *service) AddMoney(wallet_id int, tx transaction.Transaction) (transaction.Transaction, error) {
	tx, err := s.r.AddMoney(wallet_id, tx)
	if err != nil {
		return transaction.Transaction{}, err
	}

	return tx, nil
}

func (s *service) SubstractMoney(wallet_id int, tx transaction.Transaction) (transaction.Transaction, error) {
	tx, err := s.r.SubstractMoney(wallet_id, tx)
	if err != nil {
		return transaction.Transaction{}, err
	}

	return tx, nil
}

func (s *service) GetReportTransactions(transactionId *int, walletId int) ([]transaction.Transaction, error) {
	txs, err := s.r.GetReportTransactions(transactionId, walletId)
	if err != nil {
		return nil, err
	}
	return txs, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
