package storage

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Storage struct {}

type StorageConfig struct {
  Host string
  Port string
  Password string
  User string
  DBName string
  SSLMode string
}

func NewConnection(cfg *StorageConfig) (*gorm.DB, error) {
  dsn := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
    cfg.Host,
    cfg.Port,
    cfg.User,
    cfg.Password,
    cfg.DBName,
    cfg.SSLMode,
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return &gorm.DB{}, err
  }

  return db, nil
}
