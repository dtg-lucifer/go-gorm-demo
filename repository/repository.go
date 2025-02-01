package repository

import (
	"github.com/dtg-lucifer/go-gorm-postgres/models"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type Repository struct {
  DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
  return &Repository{
    DB: db,
  }
}

func (r *Repository) CreateBook(ctx fiber.Ctx) error {
  book := models.Book{}

  ctx.Bind().Body(&book)

  return nil
}
