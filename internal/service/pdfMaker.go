package service

import (
	"Practice/internal/model"
	"Practice/internal/repository/sales"
	"bytes"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"gorm.io/gorm"
	"strconv"
)

var widths = []float64{60, 60, 20, 20, 30}

func MakePdf(db *gorm.DB) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	createHeaderTable(pdf, "Company", "Name Product", "Price", "Count", "Total")
	createPdfTable(pdf, db)

	pdf.Cell(20, 10, "The biggest sales this month:")
	pdf.Ln(-1)
	createHeaderTable(pdf, "Company", "Name Product", "Price", "Count", "Total")
	createLargestTable(pdf, db)

	pdf.Cell(20, 10, "Lowest sales sales this month:")
	pdf.Ln(-1)
	createHeaderTable(pdf, "Company", "Name Product", "Price", "Count", "Total")
	createLowestTable(pdf, db)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func createPdfTable(pdf *gofpdf.Fpdf, db *gorm.DB) {
	var sum float64
	sum = createLineTable(pdf, "All", db)
	str := fmt.Sprintf("Total amount sales: %.2f", sum)
	pdf.Cell(20, 10, str)
	pdf.Ln(-1)
}

func createLargestTable(pdf *gofpdf.Fpdf, db *gorm.DB) {
	var sum float64
	sum = createLineTable(pdf, "Large", db)
	str := fmt.Sprintf("Total amount sales: %.2f", sum)
	pdf.Cell(20, 10, str)
	pdf.Ln(-1)
}

func createLowestTable(pdf *gofpdf.Fpdf, db *gorm.DB) {
	var sum float64
	sum = createLineTable(pdf, "Low", db)
	str := fmt.Sprintf("Total amount sales: %.2f", sum)
	pdf.Cell(20, 10, str)
	pdf.Ln(-1)
}

func createHeaderTable(pdf *gofpdf.Fpdf, header ...string) {
	for i, str := range header {
		pdf.CellFormat(widths[i], 10, str, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)
}

func createLineTable(pdf *gofpdf.Fpdf, param string, db *gorm.DB) (sum float64) {
	var info []model.InfoSales
	switch param {
	case "All":
		info = sales.GetAllSales(db)
	case "Large":
		info = sales.LargestSales("total", db)
	case "Low":
		info = sales.LowestSales("total", db)
	}
	data := [][]string{}
	for _, sale := range info {
		sum += (sale.Price * float64(sale.Count))
		res := []string{sale.Company, sale.NameProduct, strconv.FormatFloat(sale.Price, 'f', 2, 64),
			strconv.Itoa(sale.Count), strconv.FormatFloat(sale.Total, 'f', 2, 64)}
		data = append(data, res)
	}

	for _, row := range data {
		for i, str := range row {
			pdf.CellFormat(widths[i], 10, str, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}
	return sum
}
