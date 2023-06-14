package dependencies

import (
	"fmt"
	"os"

	"github.com/TomyPY/FinTracker/internal/platform/database"
	"github.com/TomyPY/FinTracker/internal/platform/environment"
	"github.com/joho/godotenv"
)

type Dependencies struct {
	Db *database.Database
}

func BuildDependencies(env int) (*Dependencies, error) {
	var db *database.Database

	switch env {
	case environment.Production:
		db = database.NewDb(
			database.WithUsername(os.Getenv("DB_USERNAME")),
			database.WithPassword(os.Getenv("DB_PASSWORD")),
			database.WithHost(os.Getenv("DB_HOST")),
			database.WithName("DB_NAME"),
		)
	default:
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}

		db = database.NewDb(
			database.WithUsername(os.Getenv("DB_USERNAME")),
			database.WithPassword(os.Getenv("DB_PASSWORD")),
			database.WithHost(os.Getenv("DB_HOST")),
			database.WithName(os.Getenv("DB_NAME")),
		)

	}

	fmt.Printf("%v", db)

	return &Dependencies{
		Db: db,
	}, nil
}
