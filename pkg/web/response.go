package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	Data any `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Response(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data any) {
	Response(c, status, response{Data: data})
}

func Error(c *gin.Context, status int, format string, args ...any) {
	err := errorResponse{
		Code:    strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Message: fmt.Sprintf(format, args...),
		Status:  status,
	}

	Response(c, status, err)
}
