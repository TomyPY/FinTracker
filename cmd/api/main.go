package main

import (
	"log"

	"github.com/TomyPY/FinTracker/cmd/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	eng := gin.Default()
	router := routes.NewRouter(eng)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		return err
	}
	return nil
}
