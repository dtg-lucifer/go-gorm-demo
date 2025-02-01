package repository

import (
	"fmt"
	"net/http"

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

  err := ctx.Bind().Body(&book)
  if err != nil {
    ctx.Status(http.StatusUnprocessableEntity).JSON(
      &fiber.Map{
        "message": "Some of the fields are invalid",
        "error": fmt.Sprintf("Error: %v", err),
      },
    )
    return err
  }

  err = r.DB.Create(&book).Error
  if err != nil {
    ctx.Status(http.StatusBadRequest).JSON(
      &fiber.Map{
        "message": "Could not create the book",
      },
    )
  }

  ctx.Status(http.StatusOK).JSON(
    &fiber.Map{
      "message": "Succesfully created the book",
      "book": book,
    },
  )

  return nil
}

func (r *Repository) GetAllBooks(ctx fiber.Ctx) error {
  books := &[]models.Book{}

  err := r.DB.Find(books).Error
  if err != nil {
    ctx.Status(http.StatusBadRequest).JSON(
      &fiber.Map{
        "message": "Could not get books",
        "error": fmt.Sprintf("%v", err),
      },
    )

    return err
  }

  ctx.Status(http.StatusOK).JSON(
    &fiber.Map{
      "message": "Books fetched successfully",
      "books": books,
    },
  )
  return nil
}

func (r *Repository) DeleteBook(ctx fiber.Ctx) error {
  book := models.Book{}

  id := ctx.Params("id")
  if id == "" {
    ctx.Status(http.StatusBadRequest).JSON(
      &fiber.Map{
        "message": "Please provide the id to delete the book",
      },
    )
    return fmt.Errorf("Error: didn't recieve any id")
  }
  
  err := r.DB.Delete(book, id)
  if err.Error != nil {
    ctx.Status(http.StatusBadRequest).JSON(
      &fiber.Map{
        "message": "Error in deleting the entry from database",
      },
    )
    return err.Error
  }

  ctx.Status(http.StatusOK).JSON(
    &fiber.Map{
      "message": "book deleted succesfully",
    },
  )
  return nil
}

func (r *Repository) GetBookByID(ctx fiber.Ctx) error {
  book := models.Book{}

  id := ctx.Params("id")
  if id == "" {
    ctx.Status(http.StatusBadRequest).JSON(
      &fiber.Map{
        "message": "Please provide the id to delete the book",
      },
    )
    return fmt.Errorf("Error: didn't recieve any id")
  }

  err := r.DB.Where("id = ?", id).First(book).Error
  if err != nil {
    ctx.Status(http.StatusBadRequest).JSON(
      &fiber.Map{
        "message": fmt.Sprintf("Error finding the books with the following id - %s", id),
      },
    )
    return err
  }

  ctx.Status(http.StatusOK).JSON(
    &fiber.Map{
      "message": "Success!",
      "book": book,
    },
  )
  return nil
}
