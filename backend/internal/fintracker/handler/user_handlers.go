package handler

import (
	"log/slog"
	"net/http"

	"github.com/TomyPY/FinTracker/internal/fintracker/user"
	"github.com/TomyPY/FinTracker/internal/platform/encrypt"
	"github.com/gin-gonic/gin"
)

type UserRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserRegisterHandler(repo user.Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req UserRegisterRequest
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		slog.Info("creating user", "req", req)

		hashedPassword, err := encrypt.HashPassword(req.Password)
		if err != nil {
			slog.Error("error encrypting password", "error", err)
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		err = repo.Create(ctx, user.User{
			Username: req.Username,
			Password: hashedPassword,
		})
		if err != nil {
			slog.Error("error creating user", "error", err)
			ctx.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

	}
}
