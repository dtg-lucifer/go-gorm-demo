package models

import (
  "github.com/google/uuid"
)

type Book struct {
	Author    string  `json:"author"`
	Title     string  `json:"title"`
	Publisher string  `json:"publisher"`
	Price     float32 `json:"price"`
  Id        string  `json:"id"`
}

func (b *Book) CreateBook(book Book) (Book, error) {
  return Book{}, nil
}

func (b *Book) DeleteBook(id uuid.UUID) error {
  return nil
}

func (b *Book) GetBookById(id uuid.UUID) Book {
  return Book{}
}

func (b *Book) GetAllBooks() []Book {
  return []Book{}
}
