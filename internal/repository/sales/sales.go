package sales

import (
	"Practice/internal/model"
	"Practice/internal/repository"
	"gorm.io/gorm"
)

var db = repository.InitDB()

type salesRepo struct {
	gorm.Model
	NameProduct string `gorm:"size:255"`
	Company     string `gorm:"type:varchar(100);unique_index"`
	Price       float64
	Count       int
}

func AddSalesInRepo(sale model.Sale) {
	var err error
	err = db.AutoMigrate(&salesRepo{})
	if err != nil {
		panic(err)
	}

	s := salesRepo{NameProduct: sale.NameProduct, Company: sale.Company, Price: sale.Price, Count: sale.Count}
	result := db.Create(&s)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}

func GetAllSales() []salesRepo {
	var AllSales = make([]salesRepo, 0)
	result := db.Find(&AllSales)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return AllSales
}
