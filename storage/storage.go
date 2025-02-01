package storage

import "gorm.io/gorm"

type Storage struct {}

type StorageConfig struct {}

func NewConnection(cfg StorageConfig) (*gorm.DB, error) {
  return &gorm.DB{}, nil
}
