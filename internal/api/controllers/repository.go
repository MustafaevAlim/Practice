package controllers

import (
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(db *gorm.DB) *Database {
	DB := &Database{DB: db}
	return DB
}
