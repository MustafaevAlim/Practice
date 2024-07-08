package controllers

import (
	"Practice/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

// GetPdfReport godoc
// @Summary      Получить PDF отчет
// @Description  Генерирует и возвращает PDF отчет
// @Tags         reports
// @Produce      application/pdf
// @Success      200  {file}  file
// @Failure      500  {object}  map[string]string  "Внутренняя ошибка сервера"
// @Router       /pdfReport [get]
func GetPdfReport(c echo.Context) error {
	pdfBytes, err := service.MakePdf()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate PDF"})
	}
	c.Response().Header().Set(echo.HeaderContentDisposition, "attachment; filename=report.pdf")
	return c.Blob(http.StatusOK, "application/pdf", pdfBytes)
}
