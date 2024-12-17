package main

import (
	"log/slog"

	"github.com/TomyPY/FinTracker/cmd/server/config"
	"github.com/TomyPY/FinTracker/internal/fintracker/handler"
	"github.com/TomyPY/FinTracker/platform/log"
	"github.com/gin-gonic/gin"
)

//
// @title FinanceAPI
// @version 0.1
// @description Small project to manage your finances

//@host localhost:8080
//@BasePath /api/v1

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	log.SetLog()

	cfg := config.GetConfig()

	deps, err := BuildDependencies(cfg)
	if err != nil {
		slog.Error("config error", "error", err)
		return err
	}

	router := gin.Default()

	router.GET("/ping", handler.PingHandler())

	api := router.Group("/api")

	user := api.Group("/user")
	user.POST("/register", handler.UserRegisterHandler(deps.userRepo))
	user.POST("/auth", handler.LoginHandler(deps.userRepo, deps.auth))
	user.POST("/auth/refresh", handler.RefreshHandler(deps.auth))
	user.GET("/me", handler.MiddlewareHandler(deps.auth))

	if err := router.Run(":8080"); err != nil {
		return err
	}
	return nil
}
