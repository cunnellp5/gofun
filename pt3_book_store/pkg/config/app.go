package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// "postgres://postgres:postgres@localhost:5432/books?sslmode=disable"
func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=books port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
