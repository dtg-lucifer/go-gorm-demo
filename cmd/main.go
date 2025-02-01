package main

import (
	"github.com/dtg-lucifer/go-gorm-postgres/repository"
	"github.com/dtg-lucifer/go-gorm-postgres/storage"
	"github.com/dtg-lucifer/go-gorm-postgres/utility"
	"github.com/gofiber/fiber/v3"
)

func main() {
  app := fiber.New()
  router := repository.NewRouter()

  db := utility.Must(storage.NewConnection(storage.StorageConfig{}))
  repo := repository.NewRepository(db)

  router.SetupRoutes(app, repo)

  app.Listen(":8080")
}
