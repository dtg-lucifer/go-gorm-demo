package repository

import (
	"github.com/gofiber/fiber/v3"
)

type Routes struct {}


func (routes *Routes) SetupRoutes(f *fiber.App, repository *Repository) {
  api_group := f.Group("/api/v1")

  api_group.Post("/create_book", repository.CreateBook)
  api_group.Delete("/delete_book/:id", repository.DeleteBook)
  api_group.Get("/book/:id", repository.GetBookByID)
  api_group.Get("/books", repository.GetAllBooks)
}

func NewRouter() *Routes {
  return &Routes{}
}
