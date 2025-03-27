package main

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"

	"github.com/dtg-lucifer/go-gorm-postgres/models"
	"github.com/dtg-lucifer/go-gorm-postgres/repository"
	"github.com/dtg-lucifer/go-gorm-postgres/storage"
	"github.com/dtg-lucifer/go-gorm-postgres/utility"
)

func main() {
	utility.Must[any](nil, godotenv.Load(".env"))

	app := fiber.New()
	router := repository.NewRouter()

	db := utility.Must(storage.NewConnection(&storage.StorageConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}))
	utility.Must[any](nil, models.MigrateBooks(db))
	repo := repository.NewRepository(db)

	router.SetupRoutes(app, repo)

	app.Listen(":8080")
}
