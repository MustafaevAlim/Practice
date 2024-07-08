package model

import "gorm.io/gorm"

type SalesRepo struct {
	gorm.Model
	NameProduct string `gorm:"size:255"`
	Company     string `gorm:"type:varchar(100);unique_index"`
	Price       float64
	Count       int
	Total       float64
}
