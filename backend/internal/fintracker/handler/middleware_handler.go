package handler

import (
	"log/slog"
	"net/http"

	"github.com/TomyPY/FinTracker/internal/fintracker/auth"
	"github.com/gin-gonic/gin"
)

func MiddlewareHandler(a auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := a.Auth(ctx.GetHeader("Authorization"))
		if err != nil {
			slog.Info("error authenticating", "err", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}
		ctx.Next()
	}
}
