package handler

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/TomyPY/FinTracker/internal/fintracker/auth"
	"github.com/TomyPY/FinTracker/internal/fintracker/user"
	"github.com/TomyPY/FinTracker/platform/encrypt"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func LoginHandler(repo user.Repository, a auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req LoginRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		slog.Info("prossesing login", "req", req)

		user, err := repo.GetByUsername(ctx, req.Username)
		if err != nil {
			slog.Error("error getting user by username", "error", err)
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			return
		}
		slog.Info("db user", "user", user)

		err = encrypt.VerifyPassword(req.Password, user.Password)
		if err != nil {
			slog.Error("error verifying password", "error", err)
			if errors.Is(err, encrypt.ErrInvalidPassword) {
				ctx.JSON(http.StatusBadRequest, err.Error())
			} else {
				ctx.JSON(http.StatusInternalServerError, "internal server error")
			}
			return
		}

		tokens, err := a.Create(&user)
		if err != nil {
			slog.Error("error creathing tokens", "error", err)
			ctx.JSON(http.StatusBadRequest, "bad_request")
			return
		}

		ctx.SetCookie("refresh_token", tokens.RefreshToken, int(auth.ExpirationTimeRT), "/", "localhost", false, true)

		ctx.JSON(http.StatusOK, LoginResponse{
			AccessToken: tokens.AccessToken,
		})
	}
}

func RefreshHandler(a auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("refresh_token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}
		slog.Info("Refreshing token", "token", token)

		accessToken, err := a.Refresh(ctx, token)
		if err != nil {
			slog.Error("error refreshing token", "error", err)
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		ctx.JSON(http.StatusAccepted, LoginResponse{
			AccessToken: accessToken,
		})
	}
}
