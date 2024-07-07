package service

import (
	"Practice/internal/repository/sales"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"strconv"
)

func MakePdf() {
	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()

	pdf.SetFont("Arial", "B", 12)

	header := []string{"Company", "Name Product", "Price", "Count", "Total"}

	var sum float64
	info := sales.GetAllSales()
	data := [][]string{}
	for _, sale := range info {
		sum += (sale.Price * float64(sale.Count))
		res := []string{sale.Company, sale.NameProduct, strconv.FormatFloat(sale.Price, 'f', 2, 64),
			strconv.Itoa(sale.Count), strconv.FormatFloat((sale.Price * float64(sale.Count)), 'f', 2, 64)}
		data = append(data, res)
	}
	widths := []float64{60, 60, 20, 20, 30}

	for i, str := range header {
		pdf.CellFormat(widths[i], 10, str, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	for _, row := range data {
		for i, str := range row {
			pdf.CellFormat(widths[i], 10, str, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}
	str := fmt.Sprintf("Total amount sales: %.2f", sum)
	pdf.Cell(20, 10, str)
	err := pdf.OutputFileAndClose("/home/traktor/GolandProjects/Practice/internal/service/table.pdf")
	if err != nil {

		panic(err)
	}
}
