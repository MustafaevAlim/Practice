package sales

import (
	"Practice/internal/model"
	"Practice/internal/repository/sales/converter"
	repoModel "Practice/internal/repository/sales/model"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func InRepo(sale model.Sale, db *gorm.DB) {
	s := repoModel.SalesRepo{NameProduct: sale.NameProduct, Company: sale.Company,
		Price: sale.Price, Count: sale.Count, Total: (sale.Price * float64(sale.Count))}
	result := db.Create(&s)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}

func Del(id int, db *gorm.DB) {
	if err := db.Delete(&repoModel.SalesRepo{}, id).Error; err != nil {
		log.Fatalf("failed to delete record: %v", err)
	}
}

func GetAllSales(db *gorm.DB) []model.InfoSales {
	var AllSales = make([]repoModel.SalesRepo, 0)
	result := db.Find(&AllSales)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return Map(AllSales, converter.ToSalesInfoFromRepo)
}
func LargestSales(param string, db *gorm.DB) []model.InfoSales {
	var sales []repoModel.SalesRepo
	if err := db.Order(fmt.Sprintf("%s desc", param)).Limit(3).Find(&sales).Error; err != nil {
		log.Fatal("Failed to retrieve products:", err)
	}
	return Map(sales, converter.ToSalesInfoFromRepo)
}

func LowestSales(param string, db *gorm.DB) []model.InfoSales {
	var sales []repoModel.SalesRepo
	if err := db.Order(param).Limit(3).Find(&sales).Error; err != nil {
		log.Fatal("Failed to retrieve products:", err)
	}
	return Map(sales, converter.ToSalesInfoFromRepo)
}

func Map(vs []repoModel.SalesRepo, f func(repo repoModel.SalesRepo) model.InfoSales) []model.InfoSales {
	vsm := make([]model.InfoSales, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
