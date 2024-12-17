package auth

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/TomyPY/FinTracker/internal/fintracker/session"
	"github.com/TomyPY/FinTracker/internal/fintracker/user"
	"github.com/TomyPY/FinTracker/platform/encrypt"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ExpirationTimeAT = time.Minute * 15
	ExpirationTimeRT = time.Hour * 1
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Authenticator interface {
	Create(user *user.User) (Tokens, error)
	Auth(t string) error
	Refresh(ctx context.Context, t string) (string, error)
}

type auth struct {
	session      session.Repository
	atSecret     string
	rfsSecret    string
	encryptToken string
}

func NewAuthenticator(atSecret, rfsSecret, encryptToken string, repo session.Repository) Authenticator {
	return &auth{
		session:      repo,
		atSecret:     atSecret,
		rfsSecret:    rfsSecret,
		encryptToken: encryptToken,
	}
}

func (a *auth) Auth(t string) error {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(a.atSecret), nil
	})

	// Check if token is valid
	if err != nil || !token.Valid {
		return ErrUnauthorized
	}

	return nil
}

func (a *auth) Create(user *user.User) (Tokens, error) {
	atClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID, // Subject (user identifier)
		"role": user.Role,
		"iss":  "fin-tracker",                                         // Issuer
		"exp":  time.Now().Unix() + int64(ExpirationTimeAT.Seconds()), // Expiration time
		"iat":  time.Now().Unix(),                                     // Issued at
	})

	accessToken, err := atClaims.SignedString([]byte(a.atSecret))
	if err != nil {
		return Tokens{}, err
	}

	slog.Info("accessToken", "at", accessToken)

	rfsClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID, // Subject (user identifier)
		"role": user.Role,
		"iss":  "fin-tracker",                                         // Issuer
		"exp":  time.Now().Unix() + int64(ExpirationTimeRT.Seconds()), // Expiration time
		"iat":  time.Now().Unix(),                                     // Issued at
	})

	refreshToken, err := rfsClaims.SignedString([]byte(a.rfsSecret))
	if err != nil {
		return Tokens{}, err
	}

	encryptedToken, err := encrypt.EncryptToken(refreshToken, []byte(a.encryptToken))
	if err != nil {
		return Tokens{}, err
	}

	err = a.session.Create(session.Session{
		Token:  encryptedToken,
		UserID: user.ID,
	})
	if err != nil {
		return Tokens{}, err
	}

	return Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *auth) Refresh(ctx context.Context, token string) (string, error) {

	encryptedToken, err := encrypt.EncryptToken(token, []byte(a.encryptToken))
	if err != nil {
		return "", err
	}

	slog.Info("encrypted Token", "token", encryptedToken)

	_, err = a.session.Get(ctx, encryptedToken)
	if err != nil {
		return "", err
	}

	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(a.atSecret), nil
	})

	// Check if token is valid
	if err != nil || !t.Valid {
		return "", ErrUnauthorized
	}

	cl := t.Claims.(jwt.MapClaims)

	atClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  cl["user_id"], // Subject (user identifier)
		"role": cl["role"],
		"iss":  "fin-tracker",                                         // Issuer
		"exp":  time.Now().Unix() + int64(ExpirationTimeAT.Seconds()), // Expiration time
		"iat":  time.Now().Unix(),                                     // Issued at
	})

	accessToken, err := atClaims.SignedString(a.atSecret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
