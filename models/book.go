package models

import "gorm.io/gorm"

type Book struct {
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
	Price     float32 `json:"price"`
	Id        uint    `json:"id" gorm:"primary key;autoIncrement"`
}

func MigrateBooks(db *gorm.DB) error {
  err := db.AutoMigrate(&Book{})
  return err
}
