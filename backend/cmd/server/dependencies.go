package main

import (
	"github.com/TomyPY/FinTracker/cmd/server/config"
	"github.com/TomyPY/FinTracker/internal/fintracker/auth"
	"github.com/TomyPY/FinTracker/internal/fintracker/session"
	"github.com/TomyPY/FinTracker/internal/fintracker/transaction"
	"github.com/TomyPY/FinTracker/internal/fintracker/user"
	"github.com/TomyPY/FinTracker/internal/fintracker/wallet"
	"github.com/TomyPY/FinTracker/platform/database"
)

type Dependencies struct {
	walletRepo      wallet.Repository
	userRepo        user.Repository
	sessionRepo     session.Repository
	transactionRepo transaction.Repository
	auth            auth.Authenticator
}

func BuildDependencies(cfg *config.Config) (*Dependencies, error) {

	db, err := database.NewDb(cfg)
	if err != nil {
		return nil, err
	}

	sessionRepo := session.NewRepository(db)

	return &Dependencies{
		userRepo:    user.NewRepository(db),
		sessionRepo: sessionRepo,
		auth: auth.NewAuthenticator(
			cfg.Secrets.AccessTokenSecret,
			cfg.Secrets.RefreshTokenSecret,
			cfg.Secrets.EncryptTokenSecret,
			sessionRepo,
		),
	}, nil
}
