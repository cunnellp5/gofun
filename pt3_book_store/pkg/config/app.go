package config

import (
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	connectionString := "postgres://postgres:postgres@localhost:5432/books?sslmode=disable"
	d, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
