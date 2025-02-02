package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/TomyPY/FinTracker/internal/fintracker/auth"
	"github.com/TomyPY/FinTracker/internal/fintracker/encrypt"
	"github.com/TomyPY/FinTracker/internal/fintracker/session"
	"github.com/TomyPY/FinTracker/internal/fintracker/user"
	"github.com/TomyPY/FinTracker/internal/platform/database"
)

type Dependencies struct {
	userRepo    user.Repository
	sessionRepo session.Repository
	auth        auth.Authenticator
	encrypter   encrypt.Encrypter
}

func BuildDependencies(myConf *config) (*Dependencies, error) {
	cfg := myConf.cfg

	db, err := buildDatabase(cfg)
	if err != nil {
		return nil, err
	}

	secrets := cfg.Secrets

	var encrypter encrypt.Encrypter
	if myConf.env == "development" {
		encrypter = encrypt.NewFakeEncrypter()
	} else {
		encrypter = encrypt.NewEncrypter(os.Getenv(secrets.EncryptTokenEnv))
	}

	sessionRepo := session.NewRepository(db)
	authService := auth.NewAuthenticator(
		os.Getenv(secrets.AccessTokenEnv),
		os.Getenv(secrets.RefreshTokenEnv),
		encrypter,
		sessionRepo,
	)

	return &Dependencies{
		userRepo:    user.NewRepository(db),
		sessionRepo: sessionRepo,
		auth:        authService,
		encrypter:   encrypter,
	}, nil
}

func buildDatabase(cfg *envConf) (*sql.DB, error) {
	db, err := database.NewDb(
		buildConnectionString(cfg),
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func buildConnectionString(cfg *envConf) string {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", cfg.DB.User, os.Getenv(cfg.DB.PasswordEnv), cfg.DB.Host, cfg.DB.Name)
	return conn
}
