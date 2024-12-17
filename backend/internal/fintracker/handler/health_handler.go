package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Pong")
	}
}
