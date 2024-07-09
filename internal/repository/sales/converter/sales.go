package converter

import (
	"Practice/internal/model"
	repoModel "Practice/internal/repository/sales/model"
)

func ToSalesInfoFromRepo(sale repoModel.SalesRepo) model.InfoSales {
	return model.InfoSales{
		NameProduct: sale.NameProduct,
		Company:     sale.Company,
		Price:       sale.Price,
		Count:       sale.Count,
		Total:       sale.Total,
		Date:        sale.CreatedAt,
	}

}
