package main

import (
	"os"

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

	cfg, err := NewConfig(os.Getenv("APP_ENV"))
	if err != nil {
		return err
	}

	deps, err := BuildDependencies(cfg)
	if err != nil {
		return err
	}

	router := gin.Default()
	router.GET("/ping", handler.PingHandler())

	api := router.Group("/api")

	userGroup := api.Group("/user")
	userGroup.POST("/user/register", handler.UserRegisterHandler(deps.userRepo, deps.encrypter))
	userGroup.POST("/user/auth", handler.LoginHandler(deps.userRepo, deps.auth, deps.encrypter))
	userGroup.GET("/me", handler.MeHandler(deps.auth, deps.userRepo))

	authGroup := userGroup.Group("/auth", middleware.AuthMiddleware(deps.auth))
	authGroup.POST("/logout", handler.LogoutHandler(deps.auth))

	if err := router.Run(":8080"); err != nil {
		return err
	}
	return nil
}
