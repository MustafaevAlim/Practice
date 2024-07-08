package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() (db *gorm.DB) {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Ошибка получения DB из GORM: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Ошибка проверки подключения к базе данных: %v", err)
	}
	return db
}
