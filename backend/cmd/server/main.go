package main

import (
	"log/slog"

	"github.com/TomyPY/FinTracker/cmd/server/config"
	"github.com/TomyPY/FinTracker/internal/fintracker/handler"
	"github.com/TomyPY/FinTracker/internal/fintracker/middleware"
	"github.com/TomyPY/FinTracker/internal/platform/log"

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

	api.POST("/user/register", handler.UserRegisterHandler(deps.userRepo))
	api.POST("/user/auth", handler.LoginHandler(deps.userRepo, deps.auth))

	userGroup := api.Group("/user")
	userGroup.GET("/me", handler.MeHandler(deps.auth, deps.userRepo))

	authGroup := userGroup.Group("/auth", middleware.AuthMiddleware(deps.auth))
	authGroup.POST("/logout", handler.LogoutHandler(deps.auth))

	api.GET("/protected", middleware.AuthMiddleware(deps.auth), func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})

	if err := router.Run(":8080"); err != nil {
		return err
	}
	return nil
}
