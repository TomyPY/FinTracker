package auth

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/TomyPY/FinTracker/internal/fintracker/encrypt"
	"github.com/TomyPY/FinTracker/internal/fintracker/session"
	"github.com/TomyPY/FinTracker/internal/fintracker/user"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ExpirationTimeAT = time.Minute * 15
	ExpirationTimeRT = time.Hour * 1
)

var (
	ErrUnauthorized = errors.New("unauthorized")
)

var (
	subClaim        = "sub"
	RoleClaim       = "role"
	FinTrackerClaim = "fin-tracker"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Authenticator interface {
	Create(ctx context.Context, user *user.User) (Tokens, error)
	Auth(t string) (*jwt.Token, error)
	Refresh(ctx context.Context, token string) (string, error)
	Invalidate(ctx context.Context, userID uint64) error
}

type auth struct {
	session   session.Repository
	atSecret  string
	rfsSecret string
	encrypter encrypt.Encrypter
}

func NewAuthenticator(atSecret, rfsSecret string, enc encrypt.Encrypter, repo session.Repository) Authenticator {
	return &auth{
		session:   repo,
		atSecret:  atSecret,
		rfsSecret: rfsSecret,
		encrypter: enc,
	}
}

func (a *auth) Auth(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(a.atSecret), nil
	})

	// Check if token is valid
	if err != nil || !token.Valid {
		return nil, ErrUnauthorized
	}

	return token, nil
}

func (a *auth) Create(ctx context.Context, user *user.User) (Tokens, error) {

	atClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID, // Subject (user identifier)
		"role": user.Role,
		"iss":  FinTrackerClaim,                                       // Issuer
		"exp":  time.Now().Unix() + int64(ExpirationTimeAT.Seconds()), // Expiration time
		"iat":  time.Now().Unix(),                                     // Issued at
	})

	accessToken, err := atClaims.SignedString([]byte(a.atSecret))
	if err != nil {
		return Tokens{}, err
	}

	rfsClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID, // Subject (user identifier)
		"role": user.Role,
		"iss":  FinTrackerClaim,                                       // Issuer
		"exp":  time.Now().Unix() + int64(ExpirationTimeRT.Seconds()), // Expiration time
		"iat":  time.Now().Unix(),                                     // Issued at
	})

	refreshToken, err := rfsClaims.SignedString([]byte(a.rfsSecret))
	if err != nil {
		return Tokens{}, err
	}

	rtEncrypted, err := a.encrypter.EncryptToken(refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	err = a.session.Create(ctx, session.Session{
		Token:  rtEncrypted,
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

func (a *auth) Refresh(ctx context.Context, userToken string) (string, error) {

	session, err := a.session.Get(ctx, userToken)
	if err != nil {
		return "", err
	}

	sessTkn, err := a.encrypter.DecryptToken(session.Token)
	if err != nil {
		return "", err
	}

	if sessTkn != userToken || !session.IsValid {
		slog.Error("invalid token", "error", err)
		return "", ErrUnauthorized
	}

	t, err := jwt.Parse(userToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return []byte(a.rfsSecret), nil
	})

	// Check if token is valid
	if err != nil || !t.Valid {
		slog.Error("error token invalid", "error", err)
		return "", ErrUnauthorized
	}

	cl := t.Claims.(jwt.MapClaims)

	atClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  cl[subClaim], // Subject (user identifier)
		"role": cl[RoleClaim],
		"iss":  FinTrackerClaim,                                       // Issuer
		"exp":  time.Now().Unix() + int64(ExpirationTimeAT.Seconds()), // Expiration time
		"iat":  time.Now().Unix(),                                     // Issued at
	})

	accessToken, err := atClaims.SignedString([]byte(a.atSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (a *auth) Invalidate(ctx context.Context, userID uint64) error {
	return a.session.Invalidate(ctx, userID)
}
