package main

import (
	"database/sql"
	"log"

	"github.com/TomyPY/FinTracker/cmd/api/routes"
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
		log.Fatal(err)
	}
}

func run() error {

	//Configurate depedencies
	// env := environment.GetFromString(os.Getenv("GO_ENVIRONMENT"))
	// depend, err := dependencies.BuildDependencies(env)
	// if err != nil {
	// 	return err
	// }

	//Init db
	// db, err := depend.Db.Init()
	// if err != nil {
	// 	return err
	// }

	var db *sql.DB

	//Configurate routes
	eng := gin.Default()
	router := routes.NewRouter(db, eng)
	router.MapRoutes()

	//Start server
	if err := eng.Run(); err != nil {
		return err
	}
	return nil
}
