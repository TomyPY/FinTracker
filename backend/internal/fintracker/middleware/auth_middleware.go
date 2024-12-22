package middleware

import (
	"log/slog"
	"net/http"

	"github.com/TomyPY/FinTracker/internal/fintracker/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(a auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken := ctx.GetHeader("Authorization")
		if accessToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		token, err := a.Auth(accessToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		userID := uint64(token.Claims.(jwt.MapClaims)["sub"].(float64))
		if userID == 0 {
			slog.Error("error getting user_id from token")
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", userID)
	}
}
